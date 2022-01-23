package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gopkg.in/yaml.v3"
)

func main() {
	// create the main app
	app := tview.NewApplication()

	// create a flex layout so we can put more stuff more neatly
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
			return nil
		}
		return event
	})

	// add a box for the content

	list := tview.NewList()

	list.SetBorder(true).SetTitle(" Todo Items ")
	flex.AddItem(list, 0, 8, true)

	// add a box for the footer/help
	footerBox := tview.NewTextView()

	fmt.Fprint(footerBox, "To quit press (q), to add new item press (n) and to mark/unmark a task press (space bar)")
	footerBox.SetBorder(true).SetTitle(" Help ?")
	flex.AddItem(footerBox, 0, 1, false)

	// create the list

	// read or create the yaml file.
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.ReadFile(fmt.Sprintf("%s/%s", wd, "todo.yaml"))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &list)
	if err != nil {
		panic(err)
	}

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

}
