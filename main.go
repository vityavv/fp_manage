package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
	"log"
	"fmt"
	"os/exec"
)

func main() {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	table := tview.NewTable().SetBorders(true)
	table_frame := tview.NewFrame(table).
		SetBorders(0, 0, 1, 0, 0, 0).
		AddText("Games", true, tview.AlignCenter, tcell.ColorGreen)
	form := tview.NewForm()
	form_frame := tview.NewFrame(form).
		SetBorders(0, 0, 1, 0, 0, 0).
		AddText("Press a to add a game", true, tview.AlignCenter, tcell.ColorGreen)
	flex.
		SetBorder(true).
		SetTitle("Flashplayer Manager").
		SetBorderPadding(1, 1, 1, 1)
	flex.AddItem(table_frame, 0, 1, true)
	flex.AddItem(form_frame, 0, 1, false)
	app := tview.NewApplication()

	//set headers
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
	}

	//Handle events
	table.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
	}).SetSelectedFunc(func(row int, col int) {
		path := table.GetCell(row, 1).Text
		cmd := exec.Command("sh", "-c", fmt.Sprintf("flashplayer %s", path))
		output, err := cmd.CombinedOutput()
		if err != nil {
			app.Stop()
			log.Print(string(output))
			log.Fatal(err)
		}
	}).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'a' {
			app.SetFocus(form)
			return nil
		}
		return event
	})

	//set form fields
	nameInput := tview.NewInputField().
		SetLabel("Name").
		SetFormAttributes(0, tcell.ColorWhite, tcell.ColorBlack, tcell.ColorWhite, tcell.ColorGreen)
	pathInput := tview.NewInputField().
		SetLabel("Path").
		SetFormAttributes(0, tcell.ColorWhite, tcell.ColorBlack, tcell.ColorWhite, tcell.ColorGreen)
	form.
		AddFormItem(nameInput).
		AddFormItem(pathInput).
		AddButton("SSSSSS", func() {
			app.Stop()
		})

	//Start application
	table.Select(1, 2).SetSelectable(true, true)
	if err := app.SetRoot(flex, true).Run(); err != nil {
		log.Fatal(err)
	}
}