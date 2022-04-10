package main

import (
	"image/color"
	"log"
	"os"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
	"github.com/jtdv01/hashmeme/consensus"
	"github.com/jtdv01/hashmeme/image_processor"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello Future!")

	widgetOperatorID := widget.NewEntry()
	widgetPathToImage := widget.NewMultiLineEntry()
	widgetOperatorKey := widget.NewPasswordEntry()

	// Add some defaults
	cwd, _ := os.Getwd()
	widgetPathToImage.Text = fmt.Sprintf("%s/hashmeme.png", cwd)
	var operatorID string
	var operatorKey string

	// Read .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not read .env file, you'll have to enter details manually")
	} else {
		operatorID = os.Getenv("OPERATOR_ID")
		operatorKey = os.Getenv("OPERATOR_KEY")
		widgetOperatorID.Text = operatorID
		widgetOperatorKey.Text = operatorKey
		log.Println(operatorID)
		log.Println(operatorKey)
	}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "OperatorID:", Widget: widgetOperatorID},
			{Text: "Path to image:", Widget: widgetPathToImage},
			{Text: "OperatorKey:", Widget: widgetOperatorKey},
		},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", widgetPathToImage.Text)
			textContent := image_processor.ReadTextFromImage(widgetPathToImage.Text)
			imageSha256 := image_processor.HashImageSha256(widgetPathToImage.Text)
			hashMemeMessage := consensus.NewMessage(widgetOperatorID.Text, textContent, imageSha256)

			operatorID = widgetOperatorID.Text
			operatorKey = widgetOperatorKey.Text

			dialog.ShowInformation("Result", hashMemeMessage, w)

			// Send to hgraph
			// client := consensus.CreateClient()
			// client.SendMessage(hashMemeMessage)

		},
	}

	text1 := canvas.NewText("Hashmeme", color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	content := container.New(layout.NewGridLayout(2), text1, form)

	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 700))
	w.ShowAndRun()
}
