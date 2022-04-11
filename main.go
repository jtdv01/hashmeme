package main

import (
	"fmt"
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
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
	"github.com/jtdv01/hashmeme/consensus"
	"github.com/jtdv01/hashmeme/image_processor"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello Future!")

	widgetOperatorID := widget.NewEntry()
	widgetTopicID := widget.NewEntry()
	widgetPathToImage := widget.NewMultiLineEntry()
	widgetOperatorKey := widget.NewPasswordEntry()

	// Add some defaults
	cwd, _ := os.Getwd()
	widgetPathToImage.SetText(fmt.Sprintf("%s/hashmeme.png", cwd))
	var operatorID string
	var operatorKey string

	// Read .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not read .env file, you'll have to enter details manually")
	} else {
		operatorID = os.Getenv("OPERATOR_ID")
		operatorKey = os.Getenv("OPERATOR_KEY")
		widgetOperatorID.SetText(operatorID)
		widgetOperatorKey.SetText(operatorKey)
		widgetTopicID.SetText(os.Getenv("TOPIC_ID"))
	}

	submitForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "OperatorID:", Widget: widgetOperatorID},
			{Text: "TopicID:", Widget: widgetTopicID},
			{Text: "Path to image:", Widget: widgetPathToImage},
			{Text: "OperatorKey:", Widget: widgetOperatorKey},
		},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", widgetPathToImage.Text)
			textContent := image_processor.ReadTextFromImage(widgetPathToImage.Text)
			imageSha256 := image_processor.HashImageSha256(widgetPathToImage.Text)
			topicID := widgetTopicID.Text
			hashMemeMessage := consensus.NewMessage(widgetOperatorID.Text, textContent, imageSha256)

			operatorID = widgetOperatorID.Text
			operatorKey = widgetOperatorKey.Text

			// Send to hgraph
			client := consensus.CreateClient(operatorID, operatorKey)
			txResponse := consensus.SendMessage(client, imageSha256, topicID, hashMemeMessage)

			// Display txResponse
			fmt.Println(txResponse)
			displayMessage := fmt.Sprintf("TextContent: %s\nHash:%s\n", textContent, imageSha256)
			dialog.ShowInformation("Result", displayMessage, w)
		},
	}

	queryForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Path to image:", Widget: widgetPathToImage},
		},
		OnSubmit: func() {
			topicID := widgetTopicID.Text
			textContent := image_processor.ReadTextFromImage(widgetPathToImage.Text)
			imageSha256 := image_processor.HashImageSha256(widgetPathToImage.Text)
			hashMemeMessage := consensus.NewMessage(widgetOperatorID.Text, textContent, imageSha256)
			wait := false
			_, err = hedera.NewTopicMessageQuery().
				SetTopicID(topicID).
				SetLimit(1).
				Subscribe(client, func(message hedera.TopicMessage) {
					for wait {
						time.Sleep(4 * time.Second)
						if string(message.Contents) == content {
							byteArray := message.Contents
							consensusTimestamp := message.ConsensusTimestamp
							contents := bytes.NewBuffer(byteArray).String()
							fmt.Printf("Found message: %s ConsensusTimestamp: %s\n", contents, consensusTimestamp)
							wait = false
						}
					}
				})
		},
	}

	text1 := canvas.NewText("Hashmeme", color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	content := container.New(layout.NewGridLayout(2), text1, submitForm, queryForm)

	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 700))
	w.ShowAndRun()
}
