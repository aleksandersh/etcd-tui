package pagedelete

import (
	"context"
	"fmt"
	"log"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func New(ctx context.Context, controller ui.Controller, dataSource *data.EtcdDataSource, enitity *domain.Entity) tview.Primitive {
	isKeyDeleted := false
	key := enitity.Key
	if len(key) > 100 {
		key = key[:97] + "..."
	}
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Do you want to delete the key?\n '%s'", key)).
		AddButtons([]string{"Delete", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if isKeyDeleted {
				return
			}
			if buttonIndex == 0 {
				isKeyDeleted = true
				go deleteKey(ctx, controller, dataSource, enitity.Key)
			} else if buttonIndex == 1 {
				controller.CloseDeletePage()
			}
		})
	modal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if isKeyDeleted {
			return event
		}
		if event.Key() == tcell.KeyEsc {
			controller.CloseDeletePage()
			return nil
		}
		return event
	})
	return modal
}

func deleteKey(ctx context.Context, controller ui.Controller, dataSource *data.EtcdDataSource, key string) {
	if err := dataSource.DeleteKey(ctx, key); err != nil {
		log.Fatalf("failed to delete key: %v", err) // todo
	}
	list, err := dataSource.GetEntityList(ctx)
	if err != nil {
		log.Fatalf("failed to get keys: %v", err) // todo
	}
	controller.Enque(func() { controller.ShowItems(list, false) })
}
