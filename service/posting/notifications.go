package posting

type LotPublishedNotification struct {
	PublishedAt string
}

func (l LotPublishedNotification) Build() string {
	return `
		 📅 <b>Лот <a href="{{ .Lot.Link }}">{{ .Lot.Name }}</a> (#{{ .Lot.ID }})</b> прошел модерацию и доступен на <a href="https://birzzha.me/lots/{{ .Lot.ID }}">сайте</a>
		
В <a href="https://t.me/birzzha">канале</a> объявление будет опубликовано <b>` + l.PublishedAt + ` по МСК</b>.
	`
}
