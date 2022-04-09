package consensus

import (
	"log"
	"testing"
)

// func TestClient(t *testing.T) {
// 	createClient()
// }

func TestCreateMessage(t *testing.T) {
	message := NewMessage("me", "some random meme")
	log.Println(message)
}
