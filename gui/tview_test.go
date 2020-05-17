package gui

import (
	"os"
	"testing"

	"github.com/rivo/tview"
	_ "github.com/rivo/tview"
)

func TestSimpleGUI_1(t *testing.T) {
	t.Skip()
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world")
	tview.NewApplication().SetRoot(box, true).Run() // blocking
}

func TestSimpleGUI_ListView(t *testing.T) {
	app := tview.NewApplication()
	listControl := tview.NewList().
		AddItem("a", "desc", '1', nil).
		AddItem("b", "desc", '2', nil).
		AddItem("c", "desc", '3', nil).
		AddItem("q", "desc", '4', func() {
			os.Exit(0)
		})

	app.SetRoot(listControl, false).SetFocus(listControl).Run()
}
