package tui

import (
	"context"

	"aleksandersh.dev/etcd-tui/domain"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewKeyPage(ctx context.Context, controller *Controller) tview.Primitive {
	helpView := tview.NewTextView().
		SetText(" Press Enter to enter a value\n Press Esc to go back")

	textAreaView := tview.NewTextArea()
	textAreaView.SetBorder(false).SetTitle("Enter a key").SetBorderPadding(1, 1, 1, 1)

	gridView := tview.NewGrid().
		SetRows(0, 3).
		AddItem(textAreaView, 0, 0, 1, 1, 0, 0, true).
		AddItem(helpView, 1, 0, 1, 1, 2, 0, false)

	gridView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			textAreaView.SetDisabled(false)
			controller.Focus(textAreaView)
			entity := domain.NewEntity(textAreaView.GetText(), "")
			controller.ShowValuePage(entity)
			return nil
		} else if event.Key() == tcell.KeyEsc {
			controller.CloseValuePage()
			return nil
		}
		return event
	})

	return gridView
}
