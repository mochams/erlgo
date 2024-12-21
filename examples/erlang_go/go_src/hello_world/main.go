package main

import (
	"encoding/json"
	"log"

	"github.com/mochams/erlgo"
)

type Message struct {
	Arguments []string    `json:"arguments"`
	Action    interface{} `json:"action"`
}

func main() {
	messageBytes, err := erlgo.ReadFromErlang()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received message from Erlang: %s", messageBytes)

	var message Message
	err = json.Unmarshal(messageBytes, &message)
	if err != nil {
		log.Fatal("failed to unmarshal message: %w", err)
	}

	log.Printf("Format message: %v", message)

	response := &Message{
		Arguments: []string{"ok"},
		Action:    "response",
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	err = erlgo.WriteToErlang(jsonBytes)
	if err != nil {
		log.Fatal(err)
	}
}
