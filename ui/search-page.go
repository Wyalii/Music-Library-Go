package ui

import "github.com/rivo/tview"

func NewSearchPage() *tview.Flex {
	searchFlex := tview.NewFlex()
	searchInput := tview.NewInputField()
	searchInput.SetLabel("Search Songs:")
	searchInput.SetFieldWidth(0)
	searchFlex.AddItem(searchInput, 0, 1, true)
	return searchFlex

}
