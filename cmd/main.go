package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack
	app := tview.NewApplication()
	app.SetTitle("Music-Library")
	menu := tview.NewList().AddItem("Search", "Search Songs", 's', nil).AddItem("Library", "Your music library", 'l', nil).AddItem("Playlists", "Your playlists", 'p', nil)
	menu.SetBorder(true).SetTitle("Main Menu")

	pages := tview.NewPages()
	searchPage := tview.NewTextView().SetText("Search Page Content")
	searchPage.SetBorder(true).SetTitle("Search")

	libraryPage := tview.NewTextView().SetText("Library Page Content")
	libraryPage.SetBorder(true).SetTitle("Library")

	playlistsPage := tview.NewTextView().SetText("Playlist Page Content")
	playlistsPage.SetBorder(true).SetTitle("Playlists")

	pages.AddPage("search", searchPage, true, true).AddPage("library", libraryPage, true, true).AddPage("playlists", playlistsPage, true, true)
	app.SetRoot(pages, true).Run()
}
