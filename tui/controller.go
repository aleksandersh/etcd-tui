package tui

import (
	"context"

	"aleksandersh.dev/etcd-tui/data"
	"aleksandersh.dev/etcd-tui/domain"
	"github.com/rivo/tview"
)

const (
	pageNameEntityList = "entity-list"
	pageNameKey        = "key"
	pageNameValue      = "value"
	pageNameDelete     = "delete"
)

type Controller struct {
	ctx        context.Context
	app        *tview.Application
	dataSource *data.EtcdDataSource
	pagesView  *tview.Pages
}

func NewController(ctx context.Context, app *tview.Application, dataSource *data.EtcdDataSource, pagesView *tview.Pages) *Controller {
	return &Controller{ctx: ctx, app: app, dataSource: dataSource, pagesView: pagesView}
}

func (c *Controller) ShowItems(enitityList *domain.EntityList) {
	page := NewEntityListPage(c.ctx, c, c.dataSource, enitityList)
	for _, page := range c.pagesView.GetPageNames(false) {
		c.pagesView.RemovePage(page)
	}
	c.pagesView.AddAndSwitchToPage(pageNameEntityList, page, true)
}

func (c *Controller) ShowValuePage(enitity *domain.Entity) {
	page := NewValuePage(c.ctx, c, c.dataSource, enitity)
	c.pagesView.AddAndSwitchToPage(pageNameValue, page, true)
}

func (c *Controller) ShowKeyPage() {
	page := NewKeyPage(c.ctx, c)
	c.pagesView.AddAndSwitchToPage(pageNameValue, page, true)
}

func (c *Controller) ShowDeletePage(enitity *domain.Entity) {
	page := NewDeletePage(c.ctx, c, c.dataSource, enitity)
	c.pagesView.AddAndSwitchToPage(pageNameDelete, page, true)
}

func (c *Controller) CloseKeyPage() {
	c.pagesView.RemovePage(pageNameKey)
}

func (c *Controller) CloseValuePage() {
	c.pagesView.RemovePage(pageNameValue)
}

func (c *Controller) CloseDeletePage() {
	c.pagesView.RemovePage(pageNameDelete)
}

func (c *Controller) Focus(view tview.Primitive) {
	c.app.SetFocus(view)
}

func (c *Controller) Enque(f func()) {
	c.app.QueueUpdateDraw(f)
}
