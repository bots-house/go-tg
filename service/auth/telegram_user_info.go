package auth

import "context"

type AvatarProvider func(ctx context.Context) (string, error)

type TelegramUserInfo struct {
	ID           int
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string

	// GetAvatar returns user avatar url or error
	GetAvatar AvatarProvider
}
