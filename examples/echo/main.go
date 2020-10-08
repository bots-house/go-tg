package main

import (
	"context"
	"flag"
	"log"

	"github.com/bots-house/go-tg"
)

var (
	token string
)

func init() {
	flag.StringVar(
		&token,
		"token",
		"",
		"Telegram Bot API token",
	)

	flag.Parse()
}

func handler(ctx context.Context, update *tg.Update) error {
	msg := update.Message

	// skip this update because not text message
	if msg == nil {
		return nil
	}

	_, err := msg.ReplyText(ctx, msg.Text, nil)

	return err
}

func main() {
	ctx := context.Background()

	client := tg.NewClient(token)

	me, err := client.GetMe(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("echo bot is %s", me.Username)

	tg.Polling(ctx, client, tg.HandlerFunc(handler))
}
