package tg

import "regexp"

type queryType int

const (
	queryTypeUnknown queryType = iota
	queryTypeJoinLink
	queryTypeUsername
)

var (
	// https://regex101.com/r/k17Knt/1/
	regexpUsername = regexp.MustCompile(`(?:^|@|/)([a-zA-Z0-9_]{5,32})$`)

	// https://regex101.com/r/WsJx0O/1/
	regexpJoinLink = regexp.MustCompile(`\/joinchat\/([\da-zA-Z_-]+)$`)
)

func parseResolveQuery(query string) (qt queryType, value string) {
	if matches := regexpJoinLink.FindStringSubmatch(query); len(matches) > 0 {
		return queryTypeJoinLink, matches[1]
	} else if matches := regexpUsername.FindStringSubmatch(query); len(matches) > 0 {
		return queryTypeUsername, matches[1]
	} else {
		return queryTypeUnknown, ""
	}
}
