package main

import (
	"encoding/json"
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

	application := app.New()
	window := application.NewWindow("Hello Future!")

	tabMain := createMainTab()
	tabSubmit, tabQuery := createSubmitTabs(window)

	// Create tabs
	tabs := container.NewAppTabs(
		tabMain,
		tabSubmit,
		tabQuery,
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	// Set content
	window.SetContent(tabs)
	window.Resize(fyne.NewSize(600, 600*2.2222))
	window.ShowAndRun()
}

func createMainTab() *container.TabItem {
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

func createSubmitTabs(window fyne.Window) (*container.TabItem, *container.TabItem) {
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
		OnSubmit: func() {
			onSubmitForm(widgetPathToImage.Text, widgetTopicID.Text, widgetOperatorID.Text, widgetOperatorKey.Text, window)
		},
	}

	// Query form
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
			onSubmitQueryForm(widgetTopicID.Text, widgetPathToImageQuery.Text, widgetOperatorID.Text, widgetOperatorKey.Text, window)
		},
	}
	tabSubmit := container.NewTabItem("Submit a new meme", submitForm)
	tabQuery := container.NewTabItem("Query meme", queryForm)
	return tabSubmit, tabQuery
}

func onSubmitForm(pathToImage string, topicID string, operatorID string, operatorKey string, window fyne.Window) {
	log.Println("Form submitted:", pathToImage)
	imageSha256, err := image_processor.HashImageSha256(pathToImage)
	if err != nil {
		dialog.ShowInformation("Error", fmt.Sprintf("Can't read image %s", err), window)
	} else {
		textContent, _ := image_processor.ReadTextFromImage(pathToImage)
		hashMemeMessage := consensus.NewMessage(operatorID, textContent, imageSha256)

		operatorID = operatorID
		operatorKey = operatorKey

		// Send to hgraph
		client, readOperatorIDErr, readOperatorKeyErr := consensus.CreateClient(operatorID, operatorKey)
		if readOperatorIDErr != nil || readOperatorKeyErr != nil {
			errorMessage := fmt.Sprintf("%s\n%s", readOperatorIDErr, readOperatorKeyErr)
			dialog.ShowInformation("Error", errorMessage, window)
		} else {
			txResponse := consensus.SendMessage(client, imageSha256, topicID, hashMemeMessage)
			// Display txResponse
			fmt.Println(txResponse)
			displayMessage := fmt.Sprintf("Meme submitted!\nImageSha256: %s\nAuthor: %s\nTextContent: %s", imageSha256, operatorID, textContent)
			dialog.ShowInformation("Result", displayMessage, window)
		}
	}
}

func onSubmitQueryForm(topicIDText string, pathToImage string, operatorID string, operatorKey string, window fyne.Window) {
	client, _, _ := consensus.CreateClient(operatorID, operatorKey)
	topicID, topicIDParseErr := hedera.TopicIDFromString(topicIDText)
	if topicIDParseErr != nil {
		dialog.ShowInformation("Result", "Couldn't find meme :(", window)
	} else {
		imageSha256, err := image_processor.HashImageSha256(pathToImage)

		if err != nil {
			dialog.ShowInformation("Error", fmt.Sprintf("Can't read image %s", err), window)
		}

		memeFound := false
		attemptsDone := false
		fmt.Printf("Looking for: %s\n", imageSha256)
		_, _ = hedera.NewTopicMessageQuery().
			SetTopicID(topicID).
			SetLimit(NUM_QUERY_LIMIT).
			SetMaxAttempts(MAX_QUERY_ATTEMPTS).
			Subscribe(client, func(message hedera.TopicMessage) {
				for !memeFound {
					var receivedMessage consensus.HashMemeMessage
					errMarshal := json.Unmarshal(message.Contents, &receivedMessage)
					if errMarshal != nil {
						log.Println(errMarshal)
					}

					if receivedMessage.ImageSha256 == imageSha256 {
						consensusTimestamp := message.ConsensusTimestamp
						displayMessage := fmt.Sprintf("Found meme with hash: %s\nAuthor: %s\nTextContent: %s\nConsensusTimestamp: %s", imageSha256, receivedMessage.Author, receivedMessage.TextContent, consensusTimestamp)
						log.Println(displayMessage)
						dialog.ShowInformation("Result", displayMessage, window)
						memeFound = true
						attemptsDone = true
					} else {
						fmt.Println("Could not find message. Waiting...")
						time.Sleep(2 * time.Second)
					}
				}

			})
		if !memeFound && attemptsDone {
			dialog.ShowInformation("Result", "Couldn't find meme :(", window)
		}
	}

}
