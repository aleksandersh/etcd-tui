package tui

import (
	"context"
	"fmt"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RunApp(ctx context.Context, config *domain.Config, dataSource *data.EtcdDataSource, list *domain.EntityList) error {
	tview.Styles.TertiaryTextColor = tcell.ColorDarkOrange
	tview.Styles.ContrastBackgroundColor = tcell.ColorDarkSlateGray

	app := tview.NewApplication()

	pagesView := tview.NewPages()

	controller := NewController(ctx, config, app, dataSource, pagesView)
	controller.ShowItems(list, false)

	app.SetRoot(pagesView, true)
	if err := app.Run(); err != nil {
		return fmt.Errorf("error in app.Run: %w", err)
	}
	app.Stop()
	return nil
}
