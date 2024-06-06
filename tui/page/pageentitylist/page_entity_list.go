package pageentitylist

import (
	"context"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func New(ctx context.Context, config *domain.Config, controller ui.Controller, dataSource *data.EtcdDataSource, list *domain.EntityList) tview.Primitive {
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
			case ui.RuneA:
				controller.ShowKeyPage()
				return nil
			case ui.RuneD:
				idx := itemsView.GetCurrentItem()
				if idx >= 0 {
					controller.ShowDeletePage(&list.Entities[idx])
				}
				return nil
			case ui.RuneR:
				refreshing = true
				helpView.SetText(" Refreshing...")
				go refresh(ctx, controller, helpView, dataSource, &refreshing)
			case ui.RuneH:
				controller.ShowHelpPage()
			}
		}
		return event
	})

	return containerView
}

func createHelpView() *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetMaxLines(1).
		SetText(" Press h to show the help")
}

func createContainerView(itemsView tview.Primitive, helpView tview.Primitive) *tview.Grid {
	gridView := tview.NewGrid().
		SetRows(0, 2).
		AddItem(itemsView, 0, 0, 1, 1, 0, 0, true).
		AddItem(helpView, 1, 0, 1, 1, 2, 0, false)
	return gridView
}

func createEntityListView(ctx context.Context, config *domain.Config, controller ui.Controller, list *domain.EntityList) *tview.List {
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

func refresh(ctx context.Context, controller ui.Controller, statusView *tview.TextView, dataSource *data.EtcdDataSource, refreshing *bool) {
	list, err := dataSource.GetEntityList(ctx)
	if err != nil {
		controller.Enque(func() {
			statusView.SetText(" [red]Failed to load entities[white], press r to refresh")
			*refreshing = false
		})
		return
	}
	controller.Enque(func() { controller.ShowItems(list) })
}
