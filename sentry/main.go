package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)

type WarningMessage struct {
	EventType string
	ServerID  string
	Message   string
}

func CaptureMessage(ctx context.Context, eventType, message string) {

	sentryMessage := &WarningMessage{
		Message:   message,
		EventType: eventType,
		ServerID:  "111",
	}

	byteMessage, err := json.Marshal(sentryMessage)
	if err != nil {
		fmt.Println("CaptureMessage err", "err", err.Error())
		return
	}

	fmt.Println("CaptureMessage err", string(byteMessage))

	sentry.CaptureMessage(string(byteMessage))
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://8c7ea8d213914556a5e23cb35cb67323@localhost:9000/2",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	CaptureMessage(context.Background(), "11167553911", "It works!")
}
