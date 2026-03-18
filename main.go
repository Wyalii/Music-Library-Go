package main

import (
	"GoMusicLibrary/api"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	flex := tview.NewFlex()
	menu := tview.NewList()
	songsView := tview.NewList()
	songsView.SetBorder(true)
	songsView.SetTitle(" Songs ")
	playlistsView := tview.NewBox().SetBorder(true).SetTitle(" Your Playlists ")
	var currentRight tview.Primitive = songsView
	loadSongs := func(query string) {
		songs, err := api.SearchSongs(query)
		if err != nil {
			log.Println("Error fetching songs:", err)
			return
		}
		songsView.Clear()

		for _, song := range songs {
			s := song
			songsView.AddItem(
				s.TrackName,
				s.ArtistName,
				0,
				func() {

				})
		}
	}

	showInRight := func(view tview.Primitive) {
		flex.RemoveItem(currentRight)
		flex.AddItem(view, 0, 1, true)
		currentRight = view
		app.SetFocus(view)
	}
	menu.SetBorder(true)
	menu.SetTitle("Left Panel.")
	menu.AddItem("Songs", "Browse songs", 's', func() {
		showInRight(songsView)
		loadSongs("queens")
	})
	menu.AddItem("Playlists", "Your PLaylists", 'p', func() {
		showInRight(playlistsView)
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

	flex.AddItem(menu, 30, 1, true)
	flex.AddItem(songsView, 0, 1, false)

	app.SetRoot(flex, true).Run()
}
