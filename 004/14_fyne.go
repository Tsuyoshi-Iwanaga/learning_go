package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	b := widget.NewButton("Click", func() {
		// dialog.ShowInformation("Alert", "This is sample alert!", w)
		dialog.ShowConfirm("Alert", "Plese check YES!", func(f bool) {
			if f {
				l.SetText("OK, thank you!")
			} else {
				l.SetText("Oh...")
			}
		}, w)
	})
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				nil, b, nil, nil,
			),
			l, b,
		),
	)
	w.Resize(fyne.NewSize(350, 250))
	w.ShowAndRun()
}
