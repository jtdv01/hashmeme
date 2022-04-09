package main

import (
    "image/color"

    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello Future!")

    text1 := canvas.NewText("Hello", color.NRGBA{R: 255, G: 255, B: 255, A:255 })
	content := container.NewWithoutLayout(text1)

	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 700))
	w.ShowAndRun()
}
