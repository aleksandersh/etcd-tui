package pagehelp

import (
	"context"

	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func New(ctx context.Context, controller ui.Controller) tview.Primitive {
	var textView = tview.NewTextView().
		SetText(` Press a to add a new entity
 Press d to delete an entity
 Press r to refresh the entities

 Press Enter to choose an entity

 Press Ctrl+C to exit`)
	textView.SetBorder(true).SetTitle(" Help ")

	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			controller.CloseHelpPage()
			return nil
		}

		return event
	})

	statusView := ui.CreateStatusTextView(" Press Esc to go back")

	return ui.CreateContainerGrid(textView, statusView)
}
