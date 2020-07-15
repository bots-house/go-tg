package storage

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Space struct {
	Client       *s3.S3
	Bucket       string
	PublicPrefix string
}

var (
	ErrInvalidStatusCode = errors.New("invalid status code")
)

func (s *Space) Add(
	ctx context.Context,
	dir string,
	body io.Reader,
	ext string,
) (string, error) {
	id := uuid.New().String()
	id = strings.Replace(id, "-", "", -1)

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	name := path.Join(dir, id+ext)

	contentType := mime.TypeByExtension(ext)

	content, err := ioutil.ReadAll(body)
	if err != nil {
		return "", errors.Wrap(err, "read body")
	}

	obj := s3.PutObjectInput{
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(name),
		Body:        bytes.NewReader(content),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(contentType),
	}

	_, err = s.Client.PutObjectWithContext(ctx, &obj)
	if err != nil {
		return "", errors.Wrap(err, "put object")
	}

	return name, nil
}

func (s *Space) PublicURL(p string) string {
	return strings.Join([]string{s.PublicPrefix, p}, "/")
}

func (s *Space) AddByURL(ctx context.Context, dir string, url string) (string, error) {
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

	return s.Add(ctx, dir, res.Body, ext)
}
