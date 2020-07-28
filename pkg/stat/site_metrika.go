package stat

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type SiteYandexMetrika struct {
	CounterID int
	Doer      *http.Client
}

var _ Site = &SiteYandexMetrika{}

func (site *SiteYandexMetrika) doer() *http.Client {
	if site.Doer != nil {
		return site.Doer
	}

	return http.DefaultClient
}

func (site *SiteYandexMetrika) buildUniqueVisitorsPerMonthCallURL() (string, error) {
	u, err := url.Parse("https://api-metrika.yandex.ru/stat/v1/data/bytime")
	if err != nil {
		return "", errors.Wrap(err, "parse url")
	}

	q := u.Query()

	to := time.Now()
	from := to.AddDate(0, -1, 0)

	const dateFormat = "2006-01-02"

	q.Add("ids", strconv.Itoa(site.CounterID))
	q.Add("metrics", "ym:s:users")
	q.Add("date1", from.Format(dateFormat))
	q.Add("date2", to.Format(dateFormat))

	q.Add("group", "month")

	u.RawQuery = q.Encode()

	return u.String(), nil
}

type YandexMetrikaError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`

	Errors []struct {
		ErrorType string `json:"error_type,omitempty"`
		Message   string `json:"message,omitempty"`
	} `json:"errors,omitempty"`
}

func (yme YandexMetrikaError) Error() string {
	return yme.Message
}

func (site *SiteYandexMetrika) GetUniqueVisitorsPerMonth(ctx context.Context) (int, error) {
	endpoint, err := site.buildUniqueVisitorsPerMonthCallURL()
	if err != nil {
		return -1, errors.Wrap(err, "build endpoint url")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return -1, errors.Wrap(err, "build request")
	}

	res, err := site.doer().Do(req)
	if err != nil {
		return -1, errors.Wrap(err, "do request")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var ymErr YandexMetrikaError

		if err := json.NewDecoder(res.Body).Decode(&ymErr); err != nil {
			return -1, errors.Wrap(err, "unmarshal error")
		}

		return -1, ymErr
	}

	data := struct {
		Totals [][]float64 `json:"totals"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return -1, errors.Wrap(err, "unmarshal response")
	}

	total := 0.0

	for _, arr := range data.Totals {
		for _, item := range arr {
			total += item
		}
	}

	return int(math.Round(total)), nil
}
