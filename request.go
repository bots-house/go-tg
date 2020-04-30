package tg

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Request represents RPC request to Telegram Bot API.
// It contains following info: method, token, parameters, files.
type Request struct {
	token  string
	method string

	args          map[string]string
	files         map[string]*InputFile
	attachmentIdx int
}

// NewRequest creates request with provided method.
func NewRequest(method string) *Request {
	return &Request{
		method: method,
		args: map[string]string{},
		files: map[string]*InputFile{},
	}
}

// Token returns request token.
func (r *Request) Token() string {
	return r.token
}

// Method returns request method.
func (r *Request) Method() string {
	return r.method
}

// SetToken sets request token.
func (r *Request) SetToken(token string) {
	r.token = token
}

// SetString sets request string argument.
func (r *Request) SetString(k string, v string) {
	r.args[k] = v
}

// SetOptString sets request string argument, if value is not zero.
func (r *Request) SetOptString(k string, v string) {
	if v != "" {
		r.args[k] = v
	}
}

// hasInputFiles returns true if request contains files for upload.
func (r *Request) hasInputFiles() bool {
	return len(r.files) > 0
}

// SetInputFile add file to request
func (r *Request) SetInputFile(k string, file *InputFile) {
	name := fmt.Sprintf("attachment_%d", r.attachmentIdx)
	r.SetString(k, "attach://"+name)
	r.files[name] = file
	r.attachmentIdx++
}

func (r *Request) SetJSON(k string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "marshal v")
	}

	r.args[k] = string(data)

	return nil
}