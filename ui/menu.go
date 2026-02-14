package ui

import "github.com/rivo/tview"

func NewMenu() *tview.List {
	menu := tview.NewList()
	menu.AddItem("Search Songs", "", '1', nil)
	menu.AddItem("Playlists", "", '2', nil)
	menu.SetBorder(true)
	return menu
}
