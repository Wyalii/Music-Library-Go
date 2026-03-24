package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	pages := tview.NewPages()
	menu := tview.NewList()
	mainFlex := tview.NewFlex()

	songsList := tview.NewList()
	songSearchInput := tview.NewInputField()

	songsPage := tview.NewFlex()
	songsPage.SetDirection(tview.FlexRow)
	songsPage.AddItem(songSearchInput, 3, 1, false)
	songsPage.AddItem(songsList, 0, 1, false)

	songSearchInput.SetLabel("Search: ")
	songSearchInput.SetPlaceholder("Type a song name...")
	songSearchInput.SetBorder(true)

	playlistsView := tview.NewBox().SetBorder(true).SetTitle(" Your Playlists ")

	// loadSongs := func(query string) {
	// 	songs, err := api.SearchSongs(query)
	// 	if err != nil {
	// 		log.Println("Error fetching songs:", err)
	// 		return
	// 	}
	// 	songsList.Clear()

	// 	for _, song := range songs {
	// 		s := song
	// 		songsList.AddItem(
	// 			s.TrackName,
	// 			s.ArtistName,
	// 			0,
	// 			func() {

	// 			})
	// 	}
	// }

	menu.SetBorder(true)
	menu.SetTitle("Left Panel.")
	menu.AddItem("Songs", "Browse songs", 's', func() {
		pages.SwitchToPage("songs page")
		app.SetFocus(songSearchInput)

	})
	menu.AddItem("Playlists", "Your PLaylists", 'p', func() {
		pages.SwitchToPage("playlist page")

	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			app.SetFocus(menu)
		}
		switch event.Rune() {
		case 'q':
			app.Stop()
		}
		return event
	})

	pages.AddPage("songs page", songsPage, true, true)
	pages.AddPage("playlist page", playlistsView, true, false)

	mainFlex.AddItem(menu, 0, 1, true)
	mainFlex.AddItem(pages, 0, 2, false)

	app.SetRoot(mainFlex, true).Run()
}
