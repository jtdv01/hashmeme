package main

import (
	"os"
	"fmt"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/jtdv01/hashmeme/consensus"
	"github.com/joho/godotenv"
)

func main() {

	godotErr := godotenv.Load()
	if godotErr != nil {
		panic(godotErr)
	}

	operatorID := os.Getenv("OPERATOR_ID")
	operatorKey := os.Getenv("OPERATOR_KEY")

	client := consensus.CreateClient(operatorID, operatorKey)

	//Create the transaction
	transaction := hedera.NewTopicCreateTransaction().
		SetTopicMemo("A topic for hashmemes")

	//Sign with the client operator private key and submit the transaction to a Hedera network
	txResponse, err := transaction.Execute(client)

	if err != nil {
		panic(err)
	}

	//Request the receipt of the transaction
	transactionReceipt, err := txResponse.GetReceipt(client)

	if err != nil {
		panic(err)
	}

	//Get the topic ID
	newTopicID := *transactionReceipt.TopicID

	fmt.Println(newTopicID)
}
