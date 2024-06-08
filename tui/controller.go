package tui

import (
	"context"

	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/aleksandersh/etcd-tui/tui/page/pagedelete"
	"github.com/aleksandersh/etcd-tui/tui/page/pageentitylist"
	"github.com/aleksandersh/etcd-tui/tui/page/pagehelp"
	"github.com/aleksandersh/etcd-tui/tui/page/pagekey"
	"github.com/aleksandersh/etcd-tui/tui/page/pagevalue"
	"github.com/aleksandersh/etcd-tui/tui/ui"
	"github.com/rivo/tview"
)

const (
	pageNameEntityList = "entity-list"
	pageNameKey        = "key"
	pageNameValue      = "value"
	pageNameDelete     = "delete"
	pageNameHelp       = "help"
)

type controller struct {
	ctx        context.Context
	config     *domain.Config
	app        *tview.Application
	dataSource *data.EtcdDataSource
	pagesView  *tview.Pages
}

func NewController(ctx context.Context, config *domain.Config, app *tview.Application, dataSource *data.EtcdDataSource, pagesView *tview.Pages) ui.Controller {
	return &controller{ctx: ctx, config: config, app: app, dataSource: dataSource, pagesView: pagesView}
}

func (c *controller) ShowItems(enitityList *domain.EntityList, failedToLoad bool) {
	page := pageentitylist.New(c.ctx, c.config, c, c.dataSource, enitityList, failedToLoad)
	for _, page := range c.pagesView.GetPageNames(false) {
		c.pagesView.RemovePage(page)
	}
	c.pagesView.AddAndSwitchToPage(pageNameEntityList, page, true)
}

func (c *controller) ShowValuePage(enitity *domain.Entity) {
	page := pagevalue.New(c.ctx, c, c.dataSource, enitity)
	c.pagesView.AddAndSwitchToPage(pageNameValue, page, true)
}

func (c *controller) ShowKeyPage() {
	page := pagekey.New(c.ctx, c)
	c.pagesView.AddAndSwitchToPage(pageNameKey, page, true)
}

func (c *controller) ShowDeletePage(enitity *domain.Entity) {
	page := pagedelete.New(c.ctx, c, c.dataSource, enitity)
	c.pagesView.AddAndSwitchToPage(pageNameDelete, page, true)
}

func (c *controller) ShowHelpPage() {
	page := pagehelp.New(c.ctx, c)
	c.pagesView.AddAndSwitchToPage(pageNameHelp, page, true)
}

func (c *controller) CloseKeyPage() {
	c.pagesView.RemovePage(pageNameKey)
}

func (c *controller) CloseValuePage() {
	c.pagesView.RemovePage(pageNameValue)
}

func (c *controller) CloseDeletePage() {
	c.pagesView.RemovePage(pageNameDelete)
}

func (c *controller) CloseHelpPage() {
	c.pagesView.RemovePage(pageNameHelp)
}

func (c *controller) Focus(view tview.Primitive) {
	c.app.SetFocus(view)
}

func (c *controller) Enque(f func()) {
	c.app.QueueUpdateDraw(f)
}
