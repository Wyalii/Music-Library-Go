package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	app.SetTitle("Music-Library")
	background := tview.NewBox().SetBackgroundColor(tcell.ColorDarkSlateGray)
	trackBox := tview.NewBox().SetBorder(true).SetTitle("Tracks:")
	musicList := tview.NewList().
		AddItem("Song 1", "Artist A", '1', nil).
		AddItem("Song 2", "Artist B", '2', nil).
		AddItem("Song 3", "Artist C", '3', nil)
	musicList.SetBorder(true).SetTitle("Music List:")
	mainFlex := tview.NewFlex().AddItem(musicList, 30, 1, true).AddItem(trackBox, 0, 2, false)
	pages := tview.NewPages().AddPage("background", background, true, true).AddPage("main", mainFlex, true, true)
	app.SetRoot(pages, true).Run()
}
