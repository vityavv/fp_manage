package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"log"
	"os/exec"
)

func main() {
	table := tview.NewTable().SetBorders(true)
	table_frame := tview.NewFrame(table).
		SetBorders(0, 0, 1, 0, 0, 0).
		AddText("Games", true, tview.AlignCenter, tcell.ColorGreen)
	form := tview.NewForm()
	form_frame := tview.NewFrame(form).
		SetBorders(0, 0, 1, 0, 0, 0).
		AddText("Press a to add a game", true, tview.AlignCenter, tcell.ColorGreen)
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.
		SetBorder(true).
		SetTitle("Flashplayer Manager").
		SetBorderPadding(1, 1, 1, 1)
	flex.AddItem(table_frame, 0, 1, true)
	flex.AddItem(form_frame, 0, 1, false)
	app := tview.NewApplication()

	//set headers & add games
	table.SetCell(0, 0, tview.NewTableCell("Name").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter).
		SetSelectable(false).
		SetExpansion(1))
	table.SetCell(0, 1, tview.NewTableCell("Path").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter).
		SetSelectable(false).
		SetExpansion(1))
	table.SetCell(0, 2, tview.NewTableCell("Run").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter).
		SetSelectable(false).
		SetExpansion(0))
	table.SetCell(0, 3, tview.NewTableCell("Delete").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter).
		SetSelectable(false).
		SetExpansion(0))
	addGamesToTable(table)

	//Handle events
	table.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape || key == 'q' {
			app.Stop()
		}
	}).SetSelectedFunc(func(row, col int) {
		switch col {
			case 2:
				path := table.GetCell(row, 1).Text
				cmd := exec.Command("sh", "-c", fmt.Sprintf("flashplayer %s", path))
				output, err := cmd.CombinedOutput()
				if err != nil {
					app.Stop()
					log.Print(string(output))
					log.Fatal(err)
				}
			case 3:
				err := DeleteGame(row - 1)
				if err != nil {
					app.Stop()
					log.Fatal(err)
				}
				table.RemoveRow(row)
		}
	}).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'a' {
			table.SetSelectable(false, false)
			app.SetFocus(form)
			return nil
		}
		if event.Rune() == 'q' {
			app.Stop()
			return nil
		}
		return event
	})

	//set form fields
	nameInput := tview.NewInputField().SetLabel("Name")
	nameInput.SetFormAttributes(0, tcell.ColorWhite, tcell.ColorBlack, tcell.ColorWhite, tcell.ColorGreen)
	pathInput := tview.NewInputField().SetLabel("Path")
	pathInput.SetFormAttributes(0, tcell.ColorWhite, tcell.ColorBlack, tcell.ColorWhite, tcell.ColorGreen)
	form.
		AddFormItem(nameInput).
		AddFormItem(pathInput).
		AddButton("Add", func() {
			AddGame(nameInput.GetText(), pathInput.GetText())
			nameInput.SetText("")
			pathInput.SetText("")
			addGamesToTable(table)
			table.SetSelectable(true, true)
			app.SetFocus(table)
		}).AddButton("Cancel", func() {
			nameInput.SetText("")
			pathInput.SetText("")
			table.SetSelectable(true, true)
			app.SetFocus(table)
		}).AddButton("Quit", func() {
			app.Stop()
		})

	//Start application
	table.Select(1, 2).SetSelectable(true, true)
	if err := app.SetRoot(flex, true).Run(); err != nil {
		log.Fatal(err)
	}
}

func addGamesToTable(table *tview.Table) {
	//Print games
	games, err := GetGames()
	if err != nil {
		log.Fatal(err)
	}
	for i, game := range games {
		table.SetCell(i + 1, 0, tview.NewTableCell(game.Name).
			SetSelectable(false))
		table.SetCell(i + 1, 1, tview.NewTableCell(game.Path).
			SetSelectable(false))
		table.SetCell(i + 1, 2, tview.NewTableCell("Run " + game.Name).
			SetSelectable(true))
		table.SetCell(i + 1, 3, tview.NewTableCell("Delete " + game.Name).
			SetSelectable(true))
	}

}
