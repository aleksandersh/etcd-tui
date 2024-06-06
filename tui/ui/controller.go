package ui

import (
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/rivo/tview"
)

type Controller interface {
	ShowItems(enitityList *domain.EntityList)
	ShowValuePage(enitity *domain.Entity)
	ShowKeyPage()
	ShowDeletePage(enitity *domain.Entity)
	ShowHelpPage()
	CloseKeyPage()
	CloseValuePage()
	CloseDeletePage()
	CloseHelpPage()
	Focus(view tview.Primitive)
	Enque(f func())
}
