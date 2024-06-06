package pagevalue

import (
	"context"
	"log"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func New(ctx context.Context, controller ui.Controller, dataSource *data.EtcdDataSource, enitity *domain.Entity) tview.Primitive {
	helpView := tview.NewTextView().
		SetText(" Press Enter to save the value\n Press Esc to go back")

	textAreaView := tview.NewTextArea()
	textAreaView.SetBorder(false).SetTitle(enitity.Key).SetBorderPadding(1, 1, 1, 1)
	textAreaView.SetText(enitity.Value, true)
	textAreaView.SetDisabled(false)

	gridView := tview.NewGrid().
		SetRows(0, 3).
		AddItem(textAreaView, 0, 0, 1, 1, 0, 0, true).
		AddItem(helpView, 1, 0, 1, 1, 2, 0, false)

	isTextSent := false
	gridView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			if !isTextSent {
				isTextSent = true
				textAreaView.SetDisabled(false)
				controller.Focus(textAreaView)
				go saveKeyValue(ctx, controller, dataSource, enitity.Key, textAreaView.GetText())
				return nil
			}
		} else if event.Key() == tcell.KeyEsc {
			if !isTextSent {
				controller.CloseValuePage()
			}
			return nil
		}
		return event
	})

	return gridView
}

func saveKeyValue(ctx context.Context, controller ui.Controller, dataSource *data.EtcdDataSource, key string, value string) {
	if err := dataSource.SaveKeyValue(ctx, key, value); err != nil {
		log.Fatalf("failed to save value: %v", err)
	}
	list, err := dataSource.GetEntityList(ctx)
	if err != nil {
		log.Fatalf("failed to get keys: %v", err)
	}
	controller.Enque(func() { controller.ShowItems(list) })
}
