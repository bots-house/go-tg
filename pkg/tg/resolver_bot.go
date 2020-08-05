package tg

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"github.com/bots-house/birzzha/core"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/pkg/errors"
)

type BotAPI interface {
	GetChatMembersCount(config tgbotapi.ChatConfig) (int, error)
	GetChat(config tgbotapi.ChatConfig) (tgbotapi.Chat, error)
}

type BotResolver struct {
	Client    BotAPI
	PublicURL func(id string) string
}

var (
	ErrEntityNotFound = core.NewError(
		"tg_entity_not_found",
		"entity not found in Telegram",
	)
	ErrEntityNotFoundOrBotIsNotAdmin = core.NewError(
		"tg_entity_not_found_or_bot_is_not_admin",
		"entity not found or bot is not admin of the channel",
	)

	regexpJoinLinkHex = regexp.MustCompile(`[a-fA-F\d]{32}$`)
)

func (r *BotResolver) Resolve(ctx context.Context, query string) (*ResolveResult, error) {
	qt, val := ParseResolveQuery(query)
	switch qt {
	case queryTypeJoinLink:
		return r.resolveJoinLink(ctx, val)
	case queryTypeUsername:
		return r.resolveUsername(ctx, val)
	default:
		return nil, ErrInvalidQuery
	}
}

func (r *BotResolver) getResolveResultByChat(chat *tgbotapi.Chat) (*ResolveResult, error) {
	var (
		// members count for groups and supergroups
		membersCount int
		err          error
	)

	if chat.IsChannel() || chat.IsGroup() || chat.IsSuperGroup() {
		membersCount, err = r.Client.GetChatMembersCount(tgbotapi.ChatConfig{
			ChatID: chat.ID,
		})

		if err != nil {
			return nil, errors.Wrap(err, "can't get members count")
		}
	}

	switch {
	case chat.IsChannel():
		return r.getChannelChatResult(chat, membersCount), nil
	case chat.IsGroup() || chat.IsSuperGroup():
		return r.getGroupChatResult(chat, membersCount), nil
	}

	return nil, errors.New("unknown type")
}

func (r *BotResolver) getChannelChatResult(chat *tgbotapi.Chat, membersCount int) *ResolveResult {
	result := &ResolveResult{
		Channel: &Channel{
			ID:           chat.ID,
			Name:         chat.Title,
			MembersCount: membersCount,
			Username:     chat.UserName,
			Description:  chat.Description,
		},
	}
	if chat.Photo != nil {
		result.Channel.Avatar = r.PublicURL(chat.Photo.BigFileID)
	}

	return result
}

func (r *BotResolver) getGroupChatResult(chat *tgbotapi.Chat, membersCount int) *ResolveResult {
	result := &ResolveResult{
		Group: &Group{
			ID:           chat.ID,
			Name:         chat.Title,
			MembersCount: membersCount,
			Username:     chat.UserName,
			Description:  chat.Description,
		},
	}
	if chat.Photo != nil {
		result.Group.Avatar = r.PublicURL(chat.Photo.BigFileID)
	}

	return result
}

func (r *BotResolver) ResolveByID(ctx context.Context, id int64) (*ResolveResult, error) {
	chat, err := r.getChat(tgbotapi.ChatConfig{
		ChatID: id,
	})

	if err != nil {
		return nil, errors.Wrap(err, "get chat")
	}

	return r.getResolveResultByChat(chat)
}

func (r *BotResolver) getChat(config tgbotapi.ChatConfig) (*tgbotapi.Chat, error) {
	chat, err := r.Client.GetChat(config)
	if err != nil {
		if err2, ok := err.(*tgbotapi.Error); ok && strings.Contains(err2.Message, "Bad Request: chat not found") {
			return nil, ErrEntityNotFound
		}
		return nil, errors.Wrap(err, "get chat")
	}
	return &chat, nil
}

type joinLinkPayload struct {
	CreatorUserID, GlobalChatID uint32
	RandomID                    uint64
}

func newJoinLinkPayload(buf *bytes.Buffer) (joinLinkPayload, error) {
	var p joinLinkPayload
	if err := binary.Read(buf, binary.BigEndian, &p); err != nil {
		return joinLinkPayload{}, errors.Wrap(err, "failed to read binary and create joinLinkPayload")
	}

	return p, nil
}

func (r *BotResolver) decodeJoinLink(joinLink string) (joinLinkPayload, error) {
	if regexpJoinLinkHex.Match([]byte(joinLink)) {
		decoded, err := hex.DecodeString(joinLink)
		if err != nil {
			return joinLinkPayload{}, errors.Wrap(err, "failed to decode hex of join link")
		}
		return newJoinLinkPayload(bytes.NewBuffer(decoded))
	}

	b := bytes.NewBuffer([]byte(joinLink))
	//ignore number of successfully written bytes
	_, err := base64.URLEncoding.Decode(b.Bytes(), []byte(joinLink+"=="))
	if err != nil {
		return joinLinkPayload{}, errors.Wrap(err, "failed to decode base64 of join link")
	}

	return newJoinLinkPayload(b)
}

func (r *BotResolver) resolveJoinLink(_ context.Context, joinLink string) (*ResolveResult, error) {
	payload, err := r.decodeJoinLink(joinLink)
	if err != nil {
		return nil, err
	}
	chat, err := r.getChat(tgbotapi.ChatConfig{
		SuperGroupUsername: "-100" + fmt.Sprint(payload.GlobalChatID),
	})

	switch err {
	case nil:
		return r.getResolveResultByChat(chat)
	case ErrEntityNotFound:
		//return more generic for this case error
		return nil, ErrEntityNotFoundOrBotIsNotAdmin
	default:
		return nil, err
	}
}

func (r *BotResolver) resolveUsername(_ context.Context, username string) (*ResolveResult, error) {
	chat, err := r.getChat(tgbotapi.ChatConfig{
		SuperGroupUsername: "@" + username,
	})

	if err != nil {
		return nil, errors.Wrap(err, "get chat")
	}

	return r.getResolveResultByChat(chat)
}
