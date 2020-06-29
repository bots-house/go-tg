package tg

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/bots-house/birzzha/core"
	"github.com/pkg/errors"
)

type AvatarResolver struct {
	Client *http.Client
}

const (
	webUserPageURL = "https://t.me/"
)

var (
	ErrCantDownloadAvatar = core.NewError("cant_download_avatar", "cant download avatar")
)

func (ar *AvatarResolver) ResolveAvatar(ctx context.Context, username string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", webUserPageURL+username, nil)
	if err != nil {
		return "", errors.Wrap(err, "build request")
	}

	resp, err := ar.Client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "execute request")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", ErrCantDownloadAvatar
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", ErrCantDownloadAvatar
	}

	photoURL, exists := doc.Find(".tgme_page_photo_image").Attr("src")
	if !exists {
		return "", ErrCantDownloadAvatar
	}

	return photoURL, nil
}
