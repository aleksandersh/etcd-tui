package pagekey

import (
	"context"

	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func New(ctx context.Context, controller ui.Controller) tview.Primitive {
	textAreaView := tview.NewTextArea()
	textAreaView.SetBorder(true).SetTitle(" Enter the new key ")

	statusView := ui.CreateStatusTextView(" Press Enter to enter a value")

	textAreaView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			text := textAreaView.GetText()
			if len(text) == 0 {
				statusView.SetText(" [yellow]The key cannot be empty")
				return nil
			}

			textAreaView.SetDisabled(false)
			controller.Focus(textAreaView)
			entity := domain.NewEntity(text, "")
			controller.ShowValuePage(entity)
			return nil
		} else if event.Key() == tcell.KeyEsc {
			controller.CloseKeyPage()
			return nil
		}
		return event
	})

	return ui.CreateContainerGrid(textAreaView, statusView)
}
