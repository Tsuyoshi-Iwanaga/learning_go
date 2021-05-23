package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	dt := widget.NewButton("Top", nil)
	db := widget.NewButton("Bottom", nil)
	dl := widget.NewButton("Left", nil)
	dr := widget.NewButton("Right", nil)
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				dt, db, dl, dr,
			),
			dt, db, dl, dr, widget.NewLabel("Center."),
		),
	)
	w.ShowAndRun()
}
