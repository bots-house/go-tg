package storage

import (
	"context"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type FS struct {
	Path string
}

// Add file to storage.
// Assign it random name.
func (fs *FS) Add(ctx context.Context, dir string, body io.Reader, ext string) (string, error) {
	pth := fs.getPath(dir, ext)

	if err := fs.ensureDirExists(dir); err != nil {
		return "", errors.Wrap(err, "ensure dir exists")
	}

	if err := fs.saveFile(pth, body); err != nil {
		return "", errors.Wrap(err, "save file")
	}

	return pth, nil
}

func (fs *FS) saveFile(pth string, body io.Reader) error {
	file, err := os.Create(path.Join(fs.Path, pth))
	if err != nil {
		return errors.Wrap(err, "create file")
	}
	defer file.Close()

	if _, err := io.Copy(file, body); err != nil {
		return errors.Wrap(err, "write file")
	}

	return nil
}

func (fs *FS) ensureDirExists(dir string) error {
	pth := path.Join(fs.Path, dir)
	return os.MkdirAll(pth, os.ModePerm)
}

func (fs *FS) getPath(dir, ext string) string {
	id := uuid.New().String()
	id = strings.Replace(id, "-", "", -1)

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	name := id + ext

	return path.Join(dir, name)
}

var (
	ErrInvalidStatusCode = errors.New("invalid status code")
)

// Add file to storage from remote URL
func (fs *FS) AddByURL(ctx context.Context, dir string, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", errors.Wrap(err, "create request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "execute request")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", ErrInvalidStatusCode
	}

	ext := filepath.Ext(url)
	if ext == "" {
		contentType := res.Header.Get("Content-Type")
		exts, err := mime.ExtensionsByType(contentType)
		if err != nil {
			return "", errors.Wrap(err, "extension by type")
		}

		if len(exts) == 0 {
			ext = ".file"
		} else {
			ext = exts[0]
		}
	}

	return fs.Add(ctx, dir, res.Body, ext)
}
