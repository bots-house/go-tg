package tg

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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
		args:   map[string]string{},
		files:  map[string]*InputFile{},
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
func (r *Request) SetInputFile(k string, file *InputFile) (addr string) {
	name := fmt.Sprintf("attachment_%d", r.attachmentIdx)
	addr = "attach://" + name
	if k != "" {
		r.SetString(k, addr)
	}
	r.files[name] = file
	r.attachmentIdx++
	return addr
}

// SetJSON value k.
func (r *Request) SetJSON(k string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "marshal v")
	}

	r.args[k] = string(data)

	return nil
}

// SetOptJSON value k.
func (r *Request) SetOptJSON(k string, v interface{}) error {
	if v != nil {
		return r.SetJSON(k, v)
	}
	return nil
}

// SetInt sets request int argument k to value v.
func (r *Request) SetInt(k string, v int) {
	r.args[k] = strconv.Itoa(v)
}

// SetFloat sets request float argument k to value v.
func (r *Request) SetFloat64(k string, v float64) {
	r.args[k] = strconv.FormatFloat(v, 'f', -1, 64)
}

// SetOptInt sets request int argument k to value v, if v is not zero.
func (r *Request) SetOptInt(k string, v int) {
	if v != 0 {
		r.args[k] = strconv.Itoa(v)
	}
}

func (r *Request) SetOptDuration(k string, v time.Duration) {
	if v != 0 {
		r.args[k] = strconv.Itoa(int(v.Seconds()))
	}
}

func (r *Request) SetPeer(k string, v Peer) {
	r.args[k] = v.Peer()
}

func (r *Request) SetOptBool(k string, v bool) {
	if v {
		r.args[k] = "true"
	}
}
