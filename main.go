package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/dialog"

	"github.com/jtdv01/hashmeme/image_processor"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello Future!")

	author := widget.NewEntry()
	pathToImage := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Author:", Widget: author},
			{Text: "Path to image:", Widget: pathToImage}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", pathToImage.Text)
			imageText := image_processor.ReadTextFromImage(pathToImage.Text)
			dialog.ShowInformation("Result", imageText, w)

		},
	}

	text1 := canvas.NewText("Hashmeme", color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	content := container.New(layout.NewGridLayout(2), text1, form)

	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 700))
	w.ShowAndRun()
}
