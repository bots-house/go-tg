package stat

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

type TelegramTelemetr struct {
	Doer *http.Client
}

var (
	_ Telegram = &TelegramTelemetr{}
)

type TelemetrError struct {
	Message string
}

func (err *TelemetrError) Error() string {
	return err.Message
}

func (tm *TelegramTelemetr) doer() *http.Client {
	if tm.Doer != nil {
		return tm.Doer
	}
	return http.DefaultClient
}

func (tm *TelegramTelemetr) buildGetRequest(ctx context.Context, query string) (*http.Request, error) {
	u, err := url.Parse("https://telemetr.me/api_ev.php")
	if err != nil {
		return nil, errors.Wrap(err, "parse url")
	}

	q := u.Query()
	q.Add("name", query)

	u.RawQuery = q.Encode()

	return http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
}

func (tm *TelegramTelemetr) Get(ctx context.Context, query string) (*TelegramStats, error) {
	req, err := tm.buildGetRequest(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "build request")
	}

	res, err := tm.doer().Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}

	responseErr := struct {
		Err string `json:"error"`
	}{}

	if err := json.Unmarshal(body, &responseErr); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal possible error")
	}

	if responseErr.Err != "" {
		switch responseErr.Err {
		case "channel not found":
			return nil, ErrChannelNotFound
		default:
			return nil, &TelemetrError{
				Message: responseErr.Err,
			}
		}
	}

	responseOk := struct {
		Res struct {
			ViewsPerPostAvg interface{} `json:"views_per_post_avg"`
			ViewsPerPost24  interface{} `json:"views_per_post_24"`
		} `json:"res"`
	}{}

	if err := json.Unmarshal(body, &responseOk); err != nil {
		return nil, errors.Wrap(err, "unmarshal response")
	}
	perPost24, err := convertTelemetrRespField(responseOk.Res.ViewsPerPost24)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal view per post 24 field")
	}
	perPostAvg, err := convertTelemetrRespField(responseOk.Res.ViewsPerPostAvg)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal view per post avg field")
	}

	return &TelegramStats{
		ViewsPerPostDaily: perPost24,
		ViewsPerPostAvg:   perPostAvg,
	}, nil
}

func convertTelemetrRespField(v interface{}) (int, error) {
	switch v := v.(type) {
	case float64:
		return int(v), nil
	case int:
		return v, nil
	case string:
		converted, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return converted, nil
	default:
		return 0, fmt.Errorf("uknown type")
	}
}
