package tg

func GetLinkByUsername(un string) string {
	return "https://t.me/" + un
}

func IsMessageIsNotModified(s string) bool {
	return s == "Bad Request: message is not modified: specified new message content and reply markup are exactly the same as a current content and reply markup of the message"
}
