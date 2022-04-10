package main

import (
	"image/color"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/jtdv01/hashmeme/consensus"
	"github.com/jtdv01/hashmeme/image_processor"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello Future!")

	author := widget.NewEntry()
	pathToImage := widget.NewMultiLineEntry()
	operatorKey := widget.NewPasswordEntry()

	// Add some defaults
	pwd, _ := os.Getwd()
	pathToImage.Text = pwd

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Author/OperatorID:", Widget: author},
			{Text: "Path to image:", Widget: pathToImage},
			{Text: "Operator Key:", Widget: operatorKey},
		},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", pathToImage.Text)
			textContent := image_processor.ReadTextFromImage(pathToImage.Text)
			imageSha256 := image_processor.HashImageSha256(pathToImage.Text)
			// client := consensus.CreateClient()
			hashMemeMessage := consensus.NewMessage(author.Text, textContent, imageSha256)
			dialog.ShowInformation("Result", hashMemeMessage, w)

		},
	}

	text1 := canvas.NewText("Hashmeme", color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	content := container.New(layout.NewGridLayout(2), text1, form)

	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 700))
	w.ShowAndRun()
}
