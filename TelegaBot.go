package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/mymmrac/telego"
)

func main() {
	ctx := context.Background()
	botToken := os.Getenv("8296637708:AAHPMdrAle-k8eSPoebDbWP1q5oMBZ3MSU8")

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set up a webhook on Telegram side
	_ = bot.SetWebhook(ctx, &telego.SetWebhookParams{
		URL:         "https://example.com/bot",
		SecretToken: bot.SecretToken(),
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo(ctx)
	fmt.Printf("Webhook Info: %+v\n", info)

	// Create http serve mux
	mux := http.NewServeMux()

	// Get an update channel from webhook.
	// (more on configuration in examples/updates_webhook/main.go)
	updates, _ := bot.UpdatesViaWebhook(ctx, telego.WebhookHTTPServeMux(mux, "/bot", bot.SecretToken()))

	// Start server for receiving requests from the Telegram
	go func() {
		_ = http.ListenAndServe(":443", mux)
	}()

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %+v\n", update)
	}
}
