package core

import "fmt"

type BotLinkBuilder struct {
	Useraname     string
	LoginPrefix   string
	ContactPrefix string
}

func (lb *BotLinkBuilder) LoginURL(id string) string {
	return fmt.Sprintf("https://t.me/%s?start=%s%s", lb.Useraname, lb.LoginPrefix, id)
}

func (lb *BotLinkBuilder) ContactURL(id int) string {
	return fmt.Sprintf("https://t.me/%s?start=%s%d", lb.Useraname, lb.ContactPrefix, id)
}
