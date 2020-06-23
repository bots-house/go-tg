package tg

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path"

	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/pkg/errors"
)

type FileProxy struct {
	client *tgbotapi.BotAPI
	path   string
}

func NewFileProxy(path string, client *tgbotapi.BotAPI) (*FileProxy, error) {
	if err := os.MkdirAll(path, 0700); err != nil {
		return nil, errors.Wrap(err, "make cache dir")
	}

	return &FileProxy{
		client: client,
		path:   path,
	}, nil
}

func (fp *FileProxy) getHash(id string) string {
	hash := sha256.New()
	hash.Write([]byte(id))
	return hex.EncodeToString(hash.Sum(nil))
}

func (fp *FileProxy) getPath(hash string) string {
	return path.Join(fp.path, hash)
}

func (fp *FileProxy) exists(pth string) bool {
	_, err := os.Stat(pth)
	return err == nil
}

func (fp *FileProxy) saveFile(ctx context.Context, id, path string) error {
	obj, err := fp.client.GetFile(tgbotapi.FileConfig{FileID: id})
	if err != nil {
		return errors.Wrap(err, "get file")
	}

	url := obj.Link(fp.client.Token)

	// create download request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return errors.Wrap(err, "create download req")
	}

	// execute download request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "execute download req")
	}
	defer res.Body.Close()

	// create file
	file, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "create file")
	}
	defer file.Close()

	// copy content
	if _, err := io.Copy(file, res.Body); err != nil {
		return errors.Wrap(err, "write file")
	}

	return nil
}

func (fp *FileProxy) Get(ctx context.Context, id string) (string, error) {
	hash := fp.getHash(id)

	pth := fp.getPath(hash)

	if !fp.exists(pth) {
		if err := fp.saveFile(ctx, id, pth); err != nil {
			return "", errors.Wrap(err, "save file")
		}
	}

	return pth, nil
}
