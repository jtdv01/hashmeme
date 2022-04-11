package consensus

import (
	"encoding/json"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
	"log"
)

type HashMemeMessage struct {
	Author      string
	TextContent string
	ImageSha256 string
}

func CreateClient(_operatorID string, _operatorKey string) *hedera.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load environment variables from .env file. Error:\n%s\n", err)
	}

	operatorAccountID, err := hedera.AccountIDFromString(_operatorID)
	if err != nil {
		panic(err)
	}

	operatorKey, err := hedera.PrivateKeyFromString(_operatorKey)
	if err != nil {
		panic(err)
	}

	client := hedera.ClientForTestnet()
	client.SetOperator(operatorAccountID, operatorKey)

	return client
}

func NewMessage(author string, textContent string, imageSha256 string) string {
	var hashMemeMessage HashMemeMessage
	hashMemeMessage = HashMemeMessage{Author: author, TextContent: textContent, ImageSha256: imageSha256}
	bytes, err := json.Marshal(hashMemeMessage)
	if err != nil {
		log.Fatalf("Could not marshall message %v", err)
	}
	return string(bytes)
}

func SendMessage(client *hedera.Client, memo string, topicIDString string, content string) hedera.TransactionResponse {
	topicID, errTopicId := hedera.TopicIDFromString(topicIDString)
	if errTopicId != nil {
		panic(errTopicId)
	}

	//Create the transaction
	transaction := hedera.NewTopicMessageSubmitTransaction().
		SetTransactionMemo(memo).
		SetTopicID(topicID).
		SetMessage([]byte(content))

	//Sign with the client operator private key and submit the transaction to a Hedera network
	txResponse, err := transaction.Execute(client)
	if err != nil {
		panic(err)
	}
	return txResponse
}
