package main

import (
	"log"

	"github.com/deyring/webhooks/handler"
	webhookserver "github.com/deyring/webhooks/webhook-server"
)

func main() {
	handler := handler.New()

	server := webhookserver.New(handler)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("cannot serve: %v", err)
	}
}
