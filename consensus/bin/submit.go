package main

import (
	"flag"
	"fmt"
	hedera "github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/jtdv01/hashmeme/consensus"
)

func main() {
	client := consensus.CreateClient()

	var topicIDString string
	var content string
	flag.StringVar(&topicIDString, "t", "", "Specify the topicID")
	flag.StringVar(&content, "c", "", "Specify the content")
	flag.Parse()

	topicID, err := hedera.TopicIDFromString(topicIDString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got TopicID: %s\n", topicID)
	fmt.Printf("Got content: %s\n", content)

	//Create the transaction
	transaction := hedera.NewTopicMessageSubmitTransaction().
		SetTopicID(topicID).
		SetMessage([]byte(content))

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

	//Get the transaction consensus status
	transactionStatus := transactionReceipt.Status

	fmt.Printf("The transaction consensus status is %v\n", transactionStatus)
	fmt.Printf("Receipt: %s", transactionReceipt)
}
