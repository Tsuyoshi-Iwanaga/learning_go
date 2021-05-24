package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")

	//Toolbar
	tb := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			l.SetText("Select Home Icon!")
		}),
		widget.NewToolbarAction(theme.InfoIcon(), func() {
			l.SetText("Select infomation Icon!")
		}),
	)

	//Menu
	mm := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				l.SetText("select 'New' menu item.")
			}),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
	)

	w.SetMainMenu(mm)

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				nil, tb, nil, nil,
			),
			l,
			tb,
		),
	)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
