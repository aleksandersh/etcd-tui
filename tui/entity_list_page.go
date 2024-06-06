package tui

import (
	"context"
	"log"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	runeA = 97
	runeD = 100
	runeH = 104
	runeR = 114
)

func NewEntityListPage(ctx context.Context, config *domain.Config, controller *Controller, dataSource *data.EtcdDataSource, list *domain.EntityList) tview.Primitive {
	helpView := createHelpView()
	itemsView := createEntityListView(ctx, config, controller, list)
	containerView := createContainerView(itemsView, helpView)

	refreshing := false
	containerView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if refreshing {
			return nil
		}
		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
			case runeA:
				controller.ShowKeyPage()
				return nil
			case runeD:
				idx := itemsView.GetCurrentItem()
				if idx >= 0 {
					controller.ShowDeletePage(&list.Entities[idx])
				}
				return nil
			case runeR:
				refreshing = true
				go refresh(ctx, controller, dataSource)
			}
		}
		return event
	})

	return containerView
}

func createHelpView() *tview.TextView {
	helpView := tview.NewTextView().
		SetText(" Press a to add a new key/value\n Press d to delete the key/value")
	return helpView
}

func createContainerView(itemsView tview.Primitive, helpView tview.Primitive) *tview.Grid {
	gridView := tview.NewGrid().
		SetRows(0, 3).
		AddItem(itemsView, 0, 0, 1, 1, 0, 0, true).
		AddItem(helpView, 1, 0, 1, 1, 2, 0, false)
	return gridView
}

func createEntityListView(ctx context.Context, config *domain.Config, controller *Controller, list *domain.EntityList) *tview.List {
	itemsView := tview.NewList()
	itemsView.SetHighlightFullLine(true).
		ShowSecondaryText(true).
		SetWrapAround(false).
		SetTitle(config.Title).
		SetBorder(true)

	for _, entity := range list.Entities {
		e := entity
		itemsView.AddItem(e.Key, e.Value, 0, func() {
			controller.ShowValuePage(&e)
		})
	}

	return itemsView
}

func refresh(ctx context.Context, controller *Controller, dataSource *data.EtcdDataSource) {
	list, err := dataSource.GetEntityList(ctx)
	if err != nil {
		log.Fatalf("failed to get keys: %v", err)
	}
	controller.Enque(func() { controller.ShowItems(list) })
}
