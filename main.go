package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
	"log"
)

func main() {
	table := tview.NewTable().SetBorders(true)
	app := tview.NewApplication()

	//set headers
	table.SetCell(0, 0, tview.NewTableCell("Name").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter))
	table.SetCell(0, 1, tview.NewTableCell("Path").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter))
	table.SetCell(0, 2, tview.NewTableCell("Run").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter))

	//Print games
	games, err := GetGames()
	if err != nil {
		log.Fatal(err)
	}
	for i, game := range games {
		table.SetCell(i + 1, 0, tview.NewTableCell(game.Name))
		table.SetCell(i + 1, 1, tview.NewTableCell(game.Path))
		table.SetCell(i + 1, 2, tview.NewTableCell("Not yet"))
	}

	if err := app.SetRoot(table, true).Run(); err != nil {
		log.Fatal(err)
	}
}