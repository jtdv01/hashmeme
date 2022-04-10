package consensus

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
	"log"
)

func CreateClient(_operatorID string, _operatorKey string) *hedera.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load environment variables from .env file. Error:\n%s\n", err)
	}

	operatorAccountID, err := hedera.AccountIDFromString(_operatorId)
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
