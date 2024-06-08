package pageentitylist

import (
	"context"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type view struct {
	statusView *tview.TextView
	itemsView  *tview.List
}

type viewmodel struct {
	refreshing bool
}

func New(ctx context.Context, config *domain.Config, controller ui.Controller, dataSource *data.EtcdDataSource, list *domain.EntityList, failedToLoad bool) tview.Primitive {
	itemsView := createEntityListView(ctx, config, controller, list)
	statusView := ui.CreateStatusTextView(" Press h to show the help")

	v := &view{statusView: statusView, itemsView: itemsView}
	vm := &viewmodel{refreshing: false}

	if failedToLoad {
		v.showRefreshingError()
	}

	itemsView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if vm.refreshing {
			return nil
		}
		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
			case ui.RuneA:
				controller.ShowKeyPage()
				return nil
			case ui.RuneD:
				if itemsView.GetItemCount() <= 0 {
					return nil
				}
				idx := itemsView.GetCurrentItem()
				if idx >= 0 {
					controller.ShowDeletePage(&list.Entities[idx])
				}
				return nil
			case ui.RuneR:
				vm.refreshing = true
				statusView.SetText(" Refreshing...")
				go refresh(ctx, controller, dataSource, v, vm)
			case ui.RuneH:
				controller.ShowHelpPage()
			}
		}
		return event
	})

	return ui.CreateContainerGrid(itemsView, statusView)
}

func createEntityListView(ctx context.Context, config *domain.Config, controller ui.Controller, list *domain.EntityList) *tview.List {
	itemsView := tview.NewList()
	itemsView.SetHighlightFullLine(true).
		ShowSecondaryText(true).
		SetWrapAround(false).
		SetTitle(" " + config.Title + " ").
		SetBorder(true)

	if list == nil {
		return itemsView
	}

	for _, entity := range list.Entities {
		e := entity
		itemsView.AddItem(e.Key, e.Value, 0, func() {
			controller.ShowValuePage(&e)
		})
	}

	return itemsView
}

func refresh(ctx context.Context, controller ui.Controller, dataSource *data.EtcdDataSource, v *view, vm *viewmodel) {
	list, err := dataSource.GetEntityList(ctx)
	if err != nil {
		controller.Enque(func() {
			v.showRefreshingError()
			vm.refreshing = false
		})
		return
	}
	controller.Enque(func() {
		controller.ShowItems(list, false)
	})
}

func (v *view) showRefreshingError() {
	v.statusView.SetText(" [red]Failed to load entities[white], press r to refresh")
}
