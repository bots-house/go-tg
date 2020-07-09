package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/friendsofgo/errors"
)

type TelegramWidgetInfo struct {
	// ID of user in Telegram
	ID int

	// First name of user in Telegram
	FirstName string

	// Last name of user in Telegram
	LastName string

	// Username of user in Telegram
	Username string

	// Photo URL of user from Telegram
	PhotoURL string

	// Time of user authorization
	AuthDate int64

	// Signature of data
	Hash string
}

func (twi *TelegramWidgetInfo) Encode(botTokenHash []byte) url.Values {
	vs := url.Values{}

	vs.Set("id", strconv.Itoa(twi.ID))

	vs.Set("first_name", twi.FirstName)

	if twi.LastName != "" {
		vs.Set("last_name", twi.LastName)
	}

	if twi.Username != "" {
		vs.Set("username", twi.Username)
	}

	if twi.PhotoURL != "" {
		vs.Set("photo_url", twi.PhotoURL)
	}

	vs.Set("auth_date", strconv.FormatInt(twi.AuthDate, 10))

	mac := hmac.New(sha256.New, botTokenHash)
	_, _ = mac.Write(twi.getCheckString())
	hash := mac.Sum(nil)

	vs.Set("hash", hex.EncodeToString(hash))

	return vs
}

func (twi *TelegramWidgetInfo) getCheckString() []byte {
	kv := make(map[string]string, 7)

	kv["id"] = strconv.Itoa(twi.ID)
	kv["first_name"] = twi.FirstName

	if twi.LastName != "" {
		kv["last_name"] = twi.LastName
	}

	if twi.Username != "" {
		kv["username"] = twi.Username
	}

	if twi.PhotoURL != "" {
		kv["photo_url"] = twi.PhotoURL
	}

	kv["auth_date"] = strconv.FormatInt(twi.AuthDate, 10)

	keys := make([]string, 0, len(kv))
	for k := range kv {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	lines := make([]string, len(keys))
	for i, k := range keys {
		lines[i] = k + "=" + kv[k]
	}

	return []byte(strings.Join(lines, "\n"))
}

func (twi *TelegramWidgetInfo) Check(botTokenHash []byte) (bool, error) {
	mac := hmac.New(sha256.New, botTokenHash)
	_, _ = mac.Write(twi.getCheckString())
	exceptedMAC := mac.Sum(nil)

	actualMAC, err := hex.DecodeString(twi.Hash)
	if err != nil {
		return false, errors.Wrap(err, "hash is hex encoded")
	}

	return hmac.Equal(exceptedMAC, actualMAC), nil
}

func (twi *TelegramWidgetInfo) AuthDateTime() time.Time {
	return time.Unix(twi.AuthDate, 0)
}
