package posting

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/bots-house/birzzha/core"
	"github.com/lithammer/dedent"
	"github.com/pkg/errors"
	"golang.org/x/text/message"
)

type LotPostText struct {
	Lot      *core.Lot
	Topics   core.TopicSlice
	Owner    *core.User
	Settings *core.Settings
}

var (
	printer = message.NewPrinter(message.MatchLanguage("en"))
)

func (lpt *LotPostText) template() string {
	return `
		<b>Канал на продажу:</b> <a href="{{ LotLink }}">{{ .Lot.Name }}</a>	
		{{ if .Lot.ExtraResources }}<b>Дополнительные ресурсы:</b>{{ range $extra := .Lot.ExtraResources }} <a href={{ $extra.URL }}>{{ $extra.Title }}</a>{{ end }}\n{{ end }}
		<b>Тематика:</b>{{ range .Topics }} #{{ .Slug }}{{ end }}

		<b>Подписчиков:</b> {{ .Lot.Metrics.MembersCount }}

		<b>Комментарий:</b>
		{{ .Lot.Comment }}

		<b>Цена:</b> {{ PriceHashTag .Lot.Price.Current }}₽ {{ PriceLimit .Lot.Price.Current }}

		<b>Продавец:</b> @{{ .Owner.Telegram.Username.String }}

		100% безопасность при сделках в Telegram.
		<b>Гарант от команды:</b> @{{ .Settings.CashierUsername }}
	`
}

func (lpt *LotPostText) Render() (string, error) {
	t := dedent.Dedent(lpt.template())

	tmpl, err := template.New("rendered lot").Funcs(
		template.FuncMap{
			"PriceHashTag": priceHashTag,
			"PriceLimit":   priceInterval,
			"LotLink":      lpt.Lot.Link,
		},
	).Parse(t)
	if err != nil {
		return "", errors.Wrap(err, "create template")
	}

	res := &bytes.Buffer{}

	if err := tmpl.Execute(res, lpt); err != nil {
		return "", errors.Wrap(err, "execute template")
	}
	return res.String(), nil
}

func priceHashTag(p int) string {
	price := printer.Sprintf("%d", p)
	return strings.ReplaceAll(price, ",", "'")
}

func priceInterval(price int) string {
	switch {
	case price <= 50000:
		return "#До50К"
	case price <= 100000:
		return "#До100К"
	case price <= 500000:
		return "#До500К"
	case price <= 1000000:
		return "#До1КК"
	case price > 1000000:
		return "#Выше1КК"
	}
	return ""
}
