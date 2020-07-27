package personal

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/tg"
	tgme "github.com/bots-house/tg-me"
	"github.com/otiai10/opengraph"
	"github.com/pkg/errors"
)

type LotExtraResourceParser struct {
	Telegram *tgme.Parser
}

var (
	ErrWrongContentType            = core.NewError("wrong_content_type", "wrong content type")
	ErrLotExtraResourceSizeIsLarge = core.NewError(
		"lot_extra_resource_size_is_large",
		"lot extra resource size is large (1MB max)",
	)
	ErrWrongResponseCode = core.NewError("wrong_response_code", "wrong response code")
)

const (
	maxLotExtraResourceSizeInBytes = 1 * 1024 * 1024
)

func (parser LotExtraResourceParser) Parse(ctx context.Context, url string) (*core.LotExtraResource, error) {
	_, v := tg.ParseResolveQuery(url)
	if v == "" {
		return parser.parseSiteLotExtraResource(ctx, url)
	}
	return parser.parseTelegramLotExtraResource(ctx, url)
}

func (parser LotExtraResourceParser) getLotExtraResourceOpengraph(ctx context.Context, url string) (*opengraph.OpenGraph, error) {
	if !strings.Contains(url, "http://") && !strings.Contains(url, "https://") {
		url = "https://" + url
	}

	og := opengraph.New(url)
	if og.Error != nil {
		return nil, og.Error
	}

	og.HTTPClient = parser.Telegram.Client
	req, err := http.NewRequestWithContext(ctx, "GET", og.URL.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := og.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, ErrWrongResponseCode
	}

	if !strings.HasPrefix(res.Header.Get("Content-Type"), "text/html") {
		return nil, ErrWrongContentType
	}

	if err := og.Parse(res.Body); err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if len(bytes) > maxLotExtraResourceSizeInBytes {
		return nil, ErrLotExtraResourceSizeIsLarge
	}

	return og, nil
}

func (parser LotExtraResourceParser) parseSiteLotExtraResource(ctx context.Context, url string) (*core.LotExtraResource, error) {
	result, err := parser.getLotExtraResourceOpengraph(ctx, url)
	if err != nil {
		return nil, err
	}

	resource := &core.LotExtraResource{
		URL:         url,
		Title:       result.Title,
		Description: result.Description,
		Domain:      url,
	}
	if len(result.Image) > 0 {
		resource.Image = url + result.Image[0].URL
	}

	return resource, nil
}

func (parser LotExtraResourceParser) parseTelegramLotExtraResource(ctx context.Context, url string) (*core.LotExtraResource, error) {
	result, err := parser.Telegram.Parse(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "fetch telegram resource")
	}

	_, domain := tg.ParseResolveQuery(url)

	switch {
	case result.Channel != nil:
		return &core.LotExtraResource{
			URL:         url,
			Title:       result.Channel.Title,
			Image:       result.Channel.Avatar,
			Description: result.Channel.Description,
			Domain:      domain,
		}, nil
	case result.Chat != nil:
		return &core.LotExtraResource{
			URL:         url,
			Title:       result.Chat.Title,
			Image:       result.Chat.Avatar,
			Description: result.Chat.Description,
			Domain:      domain,
		}, nil
	case result.User != nil:
		return &core.LotExtraResource{
			URL:    url,
			Title:  result.User.Name,
			Image:  result.User.Avatar,
			Domain: domain,
		}, nil
	}
	return nil, nil
}
