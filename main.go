package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

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

	tabMain := createMainTab()
	tabSubmit, tabQuery := createSubmitTabs(w)

	// Create tabs
	tabs := container.NewAppTabs(
		tabMain,
		tabSubmit,
		tabQuery,
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	// Set content
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(600, 600*2.2222))
	w.ShowAndRun()
}

func createMainTab() *container.TabItem {
	var FONT_TITLE_SIZE float32 = 36
	textHashMeme := canvas.NewText("HashMeme", color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	textHashMeme.TextSize = FONT_TITLE_SIZE
	title := container.New(layout.NewCenterLayout(), textHashMeme)
	imageResource, imageResourceLoadErr := fyne.LoadResourceFromPath("./hashmeme.png")
	if imageResourceLoadErr != nil {
		panic(imageResourceLoadErr)
	}
	imageContainer := canvas.NewImageFromResource(imageResource)
	imageContainer.FillMode = canvas.ImageFillOriginal
	mainContainer := container.New(layout.NewVBoxLayout(), title, imageContainer)
	return container.NewTabItem("Main menu", mainContainer)
}

func createSubmitTabs(w fyne.Window) (*container.TabItem, *container.TabItem) {
	var FONT_SUBHEADING_SIZE float32 = 24
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

	// Submit form
	textSubmit := canvas.NewText("Submit a new meme here", color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	textSubmit.TextSize = FONT_SUBHEADING_SIZE
	submitForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "", Widget: textSubmit},
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

	// Query form
	var NUM_QUERY_LIMIT uint64 = 1
	var MAX_QUERY_ATTEMPTS uint64 = 3
	widgetPathToImageQuery := widget.NewMultiLineEntry()
	widgetPathToImageQuery.SetText(fmt.Sprintf("%s/hashmeme.png", cwd))

	textQuery := canvas.NewText("Search the consensus records for a meme", color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	textQuery.TextSize = FONT_SUBHEADING_SIZE
	queryForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "", Widget: textQuery},
			{Text: "Path to image:", Widget: widgetPathToImageQuery},
		},
		OnSubmit: func() {
			client := consensus.CreateClient(operatorID, operatorKey)
			topicID, topicIDParseErr := hedera.TopicIDFromString(widgetTopicID.Text)
			if topicIDParseErr != nil {
				log.Fatalf("TopicID couldn't be parsed, check input %s", topicIDParseErr)
			}
			textContent := image_processor.ReadTextFromImage(widgetPathToImageQuery.Text)
			imageSha256 := image_processor.HashImageSha256(widgetPathToImageQuery.Text)
			hashMemeMessage := consensus.NewMessage(widgetOperatorID.Text, textContent, imageSha256)
			memeFound := false
			attemptsDone := false
			fmt.Printf("Looking for: %s\n", hashMemeMessage)
			_, err = hedera.NewTopicMessageQuery().
				SetTopicID(topicID).
				SetLimit(NUM_QUERY_LIMIT).
				SetMaxAttempts(MAX_QUERY_ATTEMPTS).
				Subscribe(client, func(message hedera.TopicMessage) {
					for !memeFound {
						// TODO: Change check only with imageSha256Hash
						if string(message.Contents) == hashMemeMessage {
							// byteArray := message.Contents
							consensusTimestamp := message.ConsensusTimestamp
							// contents := bytes.NewBuffer(byteArray).String()
							displayMessage := fmt.Sprintf("Found meme with hash: %s\nConsensusTimestamp: %s\n", imageSha256, consensusTimestamp)
							log.Println(displayMessage)
							dialog.ShowInformation("Result", displayMessage, w)
							memeFound = true
							attemptsDone = true
						} else {
							fmt.Println("Could not find message. Waiting...")
							time.Sleep(2 * time.Second)
						}
					}

				})
			if !memeFound && attemptsDone {
				dialog.ShowInformation("Result", "Couldn't find meme :(", w)
			}
		},
	}
	tabSubmit := container.NewTabItem("Submit a new meme", submitForm)
	tabQuery := container.NewTabItem("Query meme", queryForm)
	return tabSubmit, tabQuery
}
