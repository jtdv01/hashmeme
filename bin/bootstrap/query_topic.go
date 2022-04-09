package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/jtdv01/hashmeme/consensus"
	"time"
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

	//Create the account info query
	fmt.Println("Info about topic")
	query := hedera.NewTopicInfoQuery().
		SetTopicID(topicID)

	//Submit the query to a Hedera network
	info, err := query.Execute(client)
	if err != nil {
		panic(err)
	}

	//Print the account key to the console
	fmt.Printf("Memo: %s Expiration: %s\n", info.TopicMemo, info.ExpirationTime)

	fmt.Println("Waiting for new messages...\n\n")
	// Create the query for messages
	wait := true
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

	for wait {
		time.Sleep(4 * time.Second)
	}
	if err != nil {
		panic(err)
	}

}
