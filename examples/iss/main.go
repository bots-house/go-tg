package main

import (
	"context"
	"flag"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/mr-linch/go-tg"
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

func onlyErr(_ *tg.Message, err error) error {
	return err
}

func newHandler(stream <-chan tg.Location) tg.Handler {

	// live location -> expiration
	lives := make(map[*tg.Message]time.Time)
	mutex := sync.RWMutex{}

	go func() {

		for location := range stream {
			mutex.RLock()
			for msg := range lives {
				log.Printf("update live location of message #%d in chat #%d", msg.ID, msg.Chat.ID)
				if _, err := msg.EditLiveLocation(context.Background(), location, nil); err != nil {
					log.Printf("error when update live location: %#v", err)
				}
			}
			mutex.RUnlock()
		}

	}()

	return tg.HandlerFunc(func(ctx context.Context, update *tg.Update) error {
		msg := update.Message

		// skip this update because not text message
		if msg == nil {
			return nil
		}

		switch msg.Text {
		case "/location":
			location, err := getStationLocation(ctx)
			if err != nil {
				return errors.Wrap(err, "get station location")
			}

			return onlyErr(msg.AnswerLocation(ctx, location, nil))
		case "/live":
			location, err := getStationLocation(ctx)
			if err != nil {
				return errors.Wrap(err, "get station location")
			}

			livePeriod := tg.MaxLiveLocationPeriod
			expiresAt := time.Now().Add(livePeriod)

			answer, err := msg.AnswerLocation(ctx, location, &tg.LocationOpts{
				LivePeriod: livePeriod,
			})
			if err != nil {
				return errors.Wrap(err, "send live location")
			}

			mutex.Lock()
			lives[answer] = expiresAt
			mutex.Unlock()
		}

		return nil
	})
}

func main() {
	ctx := context.Background()

	client := tg.NewClient(token)

	me, err := client.GetMe(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.SetMyCommands(ctx, []tg.BotCommand{
		{"location", "get ISS current location"},
		{"live", "get ISS live location"},
	}); err != nil {
		log.Fatal("can't set commands")
	}

	log.Printf("echo bot is %s", me.Username)

	stream := getStationLocationStream(ctx)

	tg.Polling(ctx, client, newHandler(stream))
}
