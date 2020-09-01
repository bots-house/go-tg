package posting

import (
	"bytes"
	"html/template"
	"math"
	"strings"

	"github.com/bots-house/birzzha/core"
	"github.com/lithammer/dedent"
	"github.com/pkg/errors"
	"golang.org/x/text/message"
)

type LotPostText struct {
	Lot                     *core.Lot
	Topics                  core.TopicSlice
	Owner                   *core.User
	Settings                *core.Settings
	SiteWithPathListChannel string
	BotLogin                string
}

var (
	printer = message.NewPrinter(message.MatchLanguage("en"))
)

func (lpt *LotPostText) template() string {
	return `
		<b>Канал:</b> <a href="{{ LotLink }}">{{ .Lot.Name }}</a> {{ $url:= .SiteWithPathListChannel }}
		{{ if .Lot.ExtraResources }}<b>Дополнительные ресурсы:</b>{{ range $extra := .Lot.ExtraResources }} <a href="{{ $extra.URL }}">{{ $extra.Title }}</a>{{ end }}\n{{ end }}<b>Тематика:</b>{{ range .Topics }} <a href="{{ $url }}/lots?topics={{ .ID }}&from=channel">#{{ .MinifiedName }} </a>{{ end }}
		<b>Подписчиков:</b> {{ ApostrophyIntValue .Lot.Metrics.MembersCount }} ({{ .Lot.Metrics.PricePerMember }}₽ / пдп)
		{{ if .Lot.Metrics.PricePerView }}<b>Просмотров на пост:</b> {{ ApostrophyFloat64Value .Lot.Metrics.PricePerView }} ({{ .Lot.Metrics.PriceViewPerPostText }}₽ / просмотр)\n{{ end }}<b>Доход в месяц:</b> {{ ApostrophyIntValue .Lot.Metrics.MonthlyIncome }}₽ {{ if .Lot.Metrics.PaybackPeriod.Valid }}(окупаемость: {{ .Lot.Metrics.PaybackPeriod.Float64 }} {{ Month .Lot.Metrics.PaybackPeriod.Float64 }}){{ end }}

		<b>Комментарий:</b> «{{ .Lot.ShortComment }}»

		<b>Продавец:</b>{{if .Owner.Telegram.Username.Valid }} @{{ .Owner.Telegram.Username.String }}{{else}} <a href="tg://user?{{ .Owner.EscapedQueryUserID }}"> {{ .Owner.FirstName }} {{ .Owner.LastName.String }}</a> {{end}}

		<b>Цена:</b> {{ ApostrophyIntValue .Lot.Price.Current }}₽ {{ PriceLimit .Lot.Price.Current }}

		100% безопасность при сделках в Telegram. Гарант от команды: <a href="https://t.me/zzapusk">Запуск</a> @{{ .Settings.CashierUsername }}	
	`
}

func (lpt *LotPostText) Render() (string, error) {
	t := dedent.Dedent(lpt.template())

	tmpl, err := template.New("rendered lot").Funcs(
		template.FuncMap{
			"ApostrophyIntValue":     apostrophyIntValue,
			"ApostrophyFloat64Value": apostrophyFloat64Value,
			"PriceLimit":             priceInterval,
			"LotLink":                lpt.Lot.Link,
			"Month":                  month,
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

func month(payback float64) string {
	n := int(math.Abs(payback)) % 100

	switch {
	case n >= 2 && n <= 4:
		return "месяца"
	case n == 1:
		return "месяц"
	default:
		return "месяцев"
	}
}

func apostrophyIntValue(v int) string {
	price := printer.Sprintf("%d", v)
	return strings.ReplaceAll(price, ",", "'")
}

func apostrophyFloat64Value(v float64) string {
	price := printer.Sprintf("%.2f", v)
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
