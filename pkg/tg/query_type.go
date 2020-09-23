package tg

import "regexp"

type QueryType int

const (
	QueryTypeUnknown QueryType = iota
	QueryTypeJoinLink
	QueryTypeUsername
)

var (
	// https://regex101.com/r/k17Knt/1/
	regexpUsername = regexp.MustCompile(`(?:^|@|/)([a-zA-Z0-9_]{5,32})$`)

	// https://regex101.com/r/WsJx0O/1/
	regexpJoinLink = regexp.MustCompile(`\/joinchat\/([\da-zA-Z_-]+)$`)
)

func ParseResolveQuery(query string) (qt QueryType, value string) {
	if matches := regexpJoinLink.FindStringSubmatch(query); len(matches) > 0 {
		return QueryTypeJoinLink, matches[1]
	} else if matches := regexpUsername.FindStringSubmatch(query); len(matches) > 0 {
		return QueryTypeUsername, matches[1]
	} else {
		return QueryTypeUnknown, ""
	}
}
