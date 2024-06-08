package pagedelete

import (
	"context"
	"fmt"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type view struct {
	modalView *tview.Modal
}

type viewmodel struct {
	deleting bool
}

func New(ctx context.Context, controller ui.Controller, dataSource *data.EtcdDataSource, enitity *domain.Entity) tview.Primitive {
	key := enitity.Key
	if len(key) > 100 {
		key = key[:97] + "..."
	}
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Do you want to delete the key?\n '%s'", key)).
		AddButtons([]string{"Delete", "Cancel"})

	v := &view{modalView: modal}
	vm := &viewmodel{deleting: false}

	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if vm.deleting {
			return
		}
		controller.Unfocus()
		if buttonIndex == 0 {
			vm.deleting = true
			go deleteKey(ctx, controller, dataSource, enitity.Key, v, vm)
		} else if buttonIndex == 1 {
			controller.CloseDeletePage("")
		}
	})

	modal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if vm.deleting {
			return event
		}
		if event.Key() == tcell.KeyEsc {
			controller.CloseDeletePage("")
			return nil
		}
		return event
	})
	return modal
}

func deleteKey(ctx context.Context, controller ui.Controller, dataSource *data.EtcdDataSource, key string, v *view, vm *viewmodel) {
	if err := dataSource.DeleteKey(ctx, key); err != nil {
		controller.Enque(func() {
			controller.CloseDeletePage("[red]Failed to delete the key")
		})
		return
	}

	list, err := dataSource.GetEntityList(ctx)
	if err != nil {
		controller.Enque(func() { controller.ShowItems(nil) })
		return
	}
	controller.Enque(func() { controller.ShowItems(list) })
}
