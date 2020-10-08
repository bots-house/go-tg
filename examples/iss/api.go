package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/bots-house/go-tg"
)

const issLocationURL = "http://api.open-notify.org/iss-now.json"

func getStationLocation(ctx context.Context) (tg.Location, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, issLocationURL, nil)
	if err != nil {
		return tg.Location{}, errors.Wrap(err, "build request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return tg.Location{}, errors.Wrap(err, "do request")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return tg.Location{}, errors.Wrap(err, "read response")
	}

	response := struct {
		Message     string `json:"message"`
		ISSPosition struct {
			Longitude float64 `json:"longitude,string"`
			Latitude  float64 `json:"latitude,string"`
		} `json:"iss_position"`
	}{}

	if err := json.Unmarshal(body, &response); err != nil {
		return tg.Location{}, errors.Wrap(err, "unmarshal response error")
	}

	if response.Message != "success" {
		return tg.Location{}, errors.Wrap(err, "response is unsuccessful")
	}

	return tg.Location{
		Longitude: response.ISSPosition.Longitude,
		Latitude:  response.ISSPosition.Latitude,
	}, nil
}

func getStationLocationStream(ctx context.Context) <-chan tg.Location {
	stream := make(chan tg.Location)

	go func() {
	LOOP:
		for {
			ticker := time.NewTicker(time.Second * 3)

			select {
			case <-ctx.Done():
				break LOOP
			case <-ticker.C:
				location, err := getStationLocation(ctx)
				if err != nil {
					log.Printf("get station location failed: %v", err)
					time.Sleep(time.Second)
					continue LOOP
				}
				stream <- location
			}
		}

		close(stream)
	}()
	return stream
}
