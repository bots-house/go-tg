package tg

import (
	"context"
	"math/rand"
	"strings"

	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/core"
)

type BotResolver struct {
	Client    *tgbotapi.BotAPI
	PublicURL func(id string) string
}

var (
	ErrJoinLinkIsNotSupported = core.NewError(
		"tg_join_link_is_not_supported",
		"private links is not supported (while under development)",
	)
	ErrEntityNotFound = core.NewError(
		"tg_entity_not_found",
		"entity not found in Telegram",
	)
)

func (r *BotResolver) Resolve(ctx context.Context, query string) (*ResolveResult, error) {
	qt, val := parseResolveQuery(query)
	switch qt {
	case queryTypeJoinLink:
		return nil, ErrJoinLinkIsNotSupported
	case queryTypeUsername:
		return r.resolveUsername(ctx, val)
	default:
		return nil, ErrInvalidQuery
	}
}

func (r *BotResolver) resolveUsername(ctx context.Context, username string) (*ResolveResult, error) {
	chat, err := r.Client.GetChat(tgbotapi.ChatConfig{
		SuperGroupUsername: "@" + username,
	})
	if err != nil {
		if err2, ok := err.(*tgbotapi.Error); ok && strings.Contains(err2.Message, "Bad Request: chat not found"){
			return nil, ErrEntityNotFound
		}
		return nil, errors.Wrap(err, "get chat")
	}

	// common fields

	// avatar (for all types)
	var avatarFileID string
	if chat.Photo != nil {
		avatarFileID = chat.Photo.BigFileID
	}

	// members count for groups and supergroups
	var membersCount int
	if chat.IsChannel() || chat.IsGroup() || chat.IsSuperGroup() {
		membersCount, err = r.Client.GetChatMembersCount(tgbotapi.ChatConfig{
			ChatID: chat.ID,
		})

		if err != nil {
			return nil, errors.Wrap(err, "can't get members count")
		}
	}

	if chat.IsChannel() {
		return &ResolveResult{
			Channel: &Channel{
				ID:           chat.ID,
				Name:         chat.Title,
				Avatar:       r.PublicURL(avatarFileID),
				MembersCount: membersCount,
				Username:     chat.UserName,
				Description:  chat.Description,
				// TODO: real data
				DailyCoverage: 1 + rand.Intn(membersCount),
			},
		}, nil
	} else if chat.IsGroup() || chat.IsSuperGroup() {
		return &ResolveResult{
			Group: &Group{
				ID:           chat.ID,
				Name:         chat.Title,
				Avatar:       r.PublicURL(avatarFileID),
				MembersCount: membersCount,
				Username:     chat.UserName,
				Description:  chat.Description,
			},
		}, nil
	}

	return nil, errors.New("unknown type")
}
