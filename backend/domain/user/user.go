package user

import (
	"net/url"
	"unicode/utf8"

	utilsError "scrapquiz/utils/error"
)

const (
	minNameLength = 1
	maxNameLength = 20
)

type User struct {
	id        string
	name      string
	avatarURL string
}

func NewUser(id string, name string, avatarURL string) (*User, error) {
	if utf8.RuneCountInString(name) < minNameLength || utf8.RuneCountInString(name) > maxNameLength {
		return nil, utilsError.NewBadRequestError("user name length is invalid")
	}

	if _, err := url.ParseRequestURI(avatarURL); err != nil {
		return nil, utilsError.NewBadRequestError("avatarURL is invalid format")
	}

	return &User{
		id:        id,
		name:      name,
		avatarURL: avatarURL,
	}, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) AvatarURL() string {
	return u.avatarURL
}
