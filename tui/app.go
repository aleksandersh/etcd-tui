package tui

import (
	"context"
	"fmt"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/rivo/tview"
)

func RunApp(ctx context.Context, config *domain.Config, dataSource *data.EtcdDataSource, list *domain.EntityList) error {
	app := tview.NewApplication()

	pagesView := tview.NewPages()

	controller := NewController(ctx, config, app, dataSource, pagesView)
	controller.ShowItems(list)

	app.SetRoot(pagesView, true)
	if err := app.Run(); err != nil {
		return fmt.Errorf("error in app.Run: %w", err)
	}
	app.Stop()
	return nil
}
