package main

import (
    "log"
    "image/color"

    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello Future!")

	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Author", Widget: entry},
			{Text: "", Widget: textArea}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			log.Println("multiline:", textArea.Text)
		},
	}

    text1 := canvas.NewText("Hashmeme", color.NRGBA{R: 255, G: 255, B: 255, A:255 })
	content := container.New(layout.NewGridLayout(2), text1, form)

	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 700))
	w.ShowAndRun()
}
