package tg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)


// Doer define interface for mock client calls.
type Doer interface {
	// Do HTTP request.
	Do(r *http.Request) (*http.Response, error)
}

// ClientOption define type for customize client.
type ClientOption func(*Client)

// WithDoer sets client http doer.
func WithDoer(doer Doer) ClientOption {
	return func(client *Client) {
		client.doer = doer
	}
}

// Client it's Telegram Bot API client.
type Client struct {
	token string
	doer  Doer
}


// NewClient creates new Telegram Bot API client with provided token.
// Additional options can be specified using ClientOption.
func NewClient(token string, opts ...ClientOption) *Client {
	// create client with default values
	client := &Client{
		token: token,
		doer:  http.DefaultClient,
	}

	// customize client via options
	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Do Telegram Bot API request. Just execute request and return response as is.
func (client *Client) Do(ctx context.Context, r *Request) (*Response, error) {
	r.SetToken(client.token)

	if r.hasInputFiles() {
		return client.doMultipart(ctx, r)
	}

	return client.doURLEncoded(ctx, r)
}


func (client *Client) getCallURL(r *Request) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", r.token, r.method)
}

// doMultipart execute request using multipart encoding and streaming upload.
func (client *Client) doMultipart(ctx context.Context, r *Request) (*Response, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pr, pw := io.Pipe()
	resChan, errChan := make(chan *Response), make(chan error)

	writer := multipart.NewWriter(pw)

	// encode and copy
	go func() {
		defer pw.Close()
		defer writer.Close()

		for k, v := range r.args {
			if err := writer.WriteField(k, v); err != nil {
				errChan <- errors.Wrapf(err, "write field '%s'", k)
				return
			}
		}

		for k, v := range r.files {
			file, err := writer.CreateFormFile(k, v.Name)
			if err != nil {
				errChan <- errors.Wrapf(err, "write file '%s'", k)
				return
			}

			if _, err := io.Copy(file, v.Body); err != nil {
				errChan <- errors.Wrapf(err, "copy file content '%s'", k)
				return
			}
		}
	}()

	// upload
	go func() {
		req, err := http.NewRequestWithContext(ctx,
			http.MethodPost,
			client.getCallURL(r),
			pr,
		)
		if err != nil {
			errChan <- errors.Wrapf(err, "build http request")
			return
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())

		res, err := client.doer.Do(req)
		if err != nil {
			errChan <- errors.Wrap(err, "execute request")
			return
		}
		defer res.Body.Close()

		// TODO: check content type
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			errChan <- errors.Wrap(err, "read response body")
			return
		}

		response := &Response{}

		if err := json.Unmarshal(body, response); err != nil {
			errChan <- errors.Wrap(err, "unmarshal body")
			return
		}

		resChan <- response
	}()

	select {
	case err := <-errChan:
		return nil, err
	case res := <-resChan:
		return res, nil
	}
}

// doURLEncoded execute request using urlencoded encoding and simple upload.
func (client *Client) doURLEncoded(ctx context.Context, r *Request) (*Response, error) {
	vs := url.Values{}

	for k, v := range r.args {
		vs.Set(k, v)
	}

	var body io.Reader

	if len(vs) > 0 {
		body = strings.NewReader(vs.Encode())
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		client.getCallURL(r),
		body,
	)

	if err != nil {
		return nil, errors.Wrap(err, "build request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.doer.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "execute request")
	}
	defer res.Body.Close()

	// TODO: check content type
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	response := &Response{}

	if err := json.Unmarshal(data, response); err != nil {
		return nil, errors.Wrap(err, "unmarshal body")
	}

	return response, nil
}

// Invoke call to Telegram Bot API. Also handle unmarshalling of response to dst and errors.
func (client *Client) Invoke(ctx context.Context, r *Request, dst interface{}) error {
	res, err := client.Do(ctx, r)
	if err != nil {
		return errors.Wrap(err, "do request")
	}

	// TODO: handle error here

	if dst != nil {
		if err := res.UnmarshalResult(dst); err != nil {
			return errors.Wrap(err, "unmarshal result")
		}
	}

	return nil
}

var ErrUnsuccessfulInvoke = errors.New("unsuccessful invoke")

func (client *Client) invokeExceptedTrue(ctx context.Context, r *Request) error {
	var ok bool

	if err := client.Invoke(ctx, r, &ok); err != nil {
		return errors.Wrap(err, "invoke")
	}

	if !ok {
		return ErrUnsuccessfulInvoke
	}

	return nil
}