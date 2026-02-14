package ui

import (
	"github.com/rivo/tview"
)

type App struct {
	app  *tview.Application
	root tview.Primitive
}

func NewApp() App {
	app := tview.NewApplication()
	mainMenu := NewMenu()
	searchPage := NewSearchPage()
	pages := tview.NewPages()
	pages.SetBorder(true)
	pages.AddPage("Search Page", searchPage, true, true)
	flex := tview.NewFlex()
	flex.AddItem(mainMenu, 30, 1, true)
	flex.AddItem(pages, 0, 1, false)
	return App{
		app:  app,
		root: flex,
	}
}

func (a App) Run() error {
	return a.app.SetRoot(a.root, true).Run()
}
