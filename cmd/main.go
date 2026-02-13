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

	searchBar := tview.NewInputField().SetLabel("Search Songs:").SetFieldWidth(0)

	pages := tview.NewPages()
	searchPage := tview.NewFlex()
	searchPage.SetBorder(true).SetTitle("Search")
	searchPage.AddItem(searchBar, 0, 1, true)

	libraryPage := tview.NewFlex()
	libraryPage.SetBorder(true).SetTitle("Library")

	playlistsPage := tview.NewFlex()
	playlistsPage.SetBorder(true).SetTitle("Playlists")
	content := searchPage
	pages.AddPage("Search", searchPage, true, true).AddPage("Library", libraryPage, true, true).AddPage("Playlists", playlistsPage, true, true)

	menu.SetChangedFunc(func(index int, mainText string, secondary string, shortcut rune) {
		pages.SwitchToPage(mainText)
		switch mainText {
		case "Search":
			content = searchPage
		case "Library":
			content = libraryPage
		case "Playlists":
			content = playlistsPage
		}
	})

	layout := tview.NewFlex().AddItem(menu, 25, 0, true).AddItem(pages, 0, 1, false)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			if app.GetFocus() == menu {
				app.SetFocus(content)
			} else {
				app.SetFocus(menu)
			}
			return nil
		}
		return event
	})
	app.SetRoot(layout, true).SetFocus(menu).Run()
}
