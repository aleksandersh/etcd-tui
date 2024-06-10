package pageentitylist

import (
	"context"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Page struct {
	v         *view
	vm        *viewmodel
	Primitive *tview.Grid
}

type view struct {
	statusView *tview.TextView
	itemsView  *tview.List
}

type viewmodel struct {
	refreshing bool
}

func New(ctx context.Context, config *domain.Config, controller ui.Controller, dataSource *data.EtcdDataSource, list *domain.EntityList) *Page {
	itemsView := createEntityListView(config, controller, list)
	statusView := ui.CreateStatusTextView(" Press h to show the help")

	v := &view{statusView: statusView, itemsView: itemsView}
	vm := &viewmodel{refreshing: false}

	if list == nil {
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
				v.showStatusText("Refreshing...")
				go refresh(ctx, controller, dataSource, v, vm)
			case ui.RuneH:
				controller.ShowHelpPage()
			}
		}
		return event
	})

	p := ui.CreateContainerGrid(itemsView, statusView)
	return &Page{v: v, vm: vm, Primitive: p}
}

func (p *Page) ShowStatusText(text string) {
	p.v.showStatusText(text)
}

func createEntityListView(config *domain.Config, controller ui.Controller, list *domain.EntityList) *tview.List {
	itemsView := tview.NewList()
	itemsView.SetHighlightFullLine(true).
		ShowSecondaryText(true).
		SetWrapAround(false).
		SetBorder(true)

	if len(config.Title) > 0 {
		itemsView.SetTitle(" " + config.Title + " ")
	}

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
		controller.ShowItems(list)
	})
}

func (v *view) showRefreshingError() {
	v.showStatusText("[red]Failed to load entities[white], press r to refresh")
}

func (v *view) showStatusText(text string) {
	v.statusView.SetText(" " + text)
}
