package tui

import (
	"context"
	"fmt"

	"aleksandersh.dev/etcd-tui/data"
	"aleksandersh.dev/etcd-tui/domain"
	"github.com/rivo/tview"
)

func RunApp(ctx context.Context, dataSource *data.EtcdDataSource, list *domain.EntityList) error {
	app := tview.NewApplication()

	pagesView := tview.NewPages()

	controller := NewController(ctx, app, dataSource, pagesView)
	controller.ShowItems(list)

	app.SetRoot(pagesView, true)
	if err := app.Run(); err != nil {
		return fmt.Errorf("error in app.Run: %w", err)
	}
	app.Stop()
	return nil
}
