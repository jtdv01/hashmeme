This directory contains adhoc scripts to bootstrap the environment for consensus.

1. Create the topic with `go run create_topic.go`. This will print out the `topic-id`
2. Submit a new message with `go run bootstrap/submit_message.go -t="<topid-id>" -c="hello future"`
3. Read message by subscribing `go run bootstrap/get_message.go -t="<topid-id>" -c="hello future"`