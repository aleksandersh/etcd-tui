package tui

import (
	"context"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	runeA = 97
	runeD = 100
)

func NewEntityListPage(ctx context.Context, config *domain.Config, controller *Controller, dataSource *data.EtcdDataSource, list *domain.EntityList) tview.Primitive {
	helpView := createHelpView()
	itemsView := createEntityListView(ctx, config, controller, list)
	containerView := createContainerView(itemsView, helpView)

	containerView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune {
			if event.Rune() == runeA {
				controller.ShowKeyPage()
				return nil
			} else if event.Rune() == runeD {
				idx := itemsView.GetCurrentItem()
				if idx >= 0 {
					controller.ShowDeletePage(&list.Entities[idx])
				}
				return nil
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
