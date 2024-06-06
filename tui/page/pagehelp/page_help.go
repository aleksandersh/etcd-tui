package pagehelp

import (
	"context"

	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func New(ctx context.Context, controller ui.Controller) tview.Primitive {
	var textView = tview.NewTextView().
		SetText(`
 Press a to add a new entity
 Press d to delete an entity
 Press r to refresh the entities

 Press Enter to choose an entity
 Press Esc to go back

 Press Ctrl+C to exit`)
	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			controller.CloseHelpPage()
			return nil
		}

		return event
	})
	return textView
}
