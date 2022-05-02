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

func CreateClient(_operatorID string, _operatorKey string) (*hedera.Client, error, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load environment variables from .env file. Error:\n%s\n", err)
	}

	operatorAccountID, readOperatorIDErr := hedera.AccountIDFromString(_operatorID)
	if readOperatorIDErr != nil {
		panic(readOperatorIDErr)
	}

	operatorKey, readOperatorKeyErr := hedera.PrivateKeyFromString(_operatorKey)
	if readOperatorKeyErr != nil {
		panic(readOperatorKeyErr)
	}

	client := hedera.ClientForTestnet()
	client.SetOperator(operatorAccountID, operatorKey)

	return client, readOperatorIDErr, readOperatorKeyErr
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
