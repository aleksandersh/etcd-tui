package pagevalue

import (
	"context"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type view struct {
	statusView   *tview.TextView
	textAreaView *tview.TextArea
}

type viewmodel struct {
	isValueSaving bool
}

func New(ctx context.Context, controller ui.Controller, dataSource *data.EtcdDataSource, enitity *domain.Entity) tview.Primitive {
	textAreaView := tview.NewTextArea()
	textAreaView.SetBorder(true).SetTitle(" Enter the value ")
	textAreaView.SetText(enitity.Value, true)

	statusView := ui.CreateStatusTextView(" Press Enter to save the entry")

	v := &view{statusView: statusView, textAreaView: textAreaView}
	vm := &viewmodel{isValueSaving: false}

	textAreaView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			if !vm.isValueSaving {
				vm.isValueSaving = true
				controller.Unfocus()
				statusView.SetText(" Saving...")
				go saveKeyValue(ctx, controller, dataSource, enitity.Key, textAreaView.GetText(), v, vm)
				return nil
			}
		}
		return event
	})

	grid := ui.CreateContainerGrid(textAreaView, statusView)

	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			if !vm.isValueSaving {
				controller.CloseValuePage()
			}
			return nil
		}
		return event
	})

	return grid
}

func saveKeyValue(ctx context.Context, controller ui.Controller, dataSource *data.EtcdDataSource, key string, value string, v *view, vm *viewmodel) {
	if err := dataSource.SaveKeyValue(ctx, key, value); err != nil {
		controller.Enque(func() {
			vm.isValueSaving = false
			controller.Focus(v.textAreaView)
			v.statusView.SetText(" [red]Failed to save the entry[white], press Enter to retry")
		})
		return
	}
	list, err := dataSource.GetEntityList(ctx)
	if err != nil {
		controller.Enque(func() {
			controller.ShowItems(nil, true)
		})
	}
	controller.Enque(func() {
		controller.ShowItems(list, false)
	})
}
