package papi

import "time"

type TokenStorage struct {
	tokens map[string]TokenItem
}

type TokenItem struct {
	Id string
	AccessToken string
	AccessTokenExpire time.Time
	RefreshToken string
	RefreshTokenExpireAt time.Time
}

func (t TokenItem) AccessExpired() bool {
	return t.AccessTokenExpire.Before(time.Now())
}

func (t TokenItem) RefreshExpired() bool {
	return t.RefreshTokenExpireAt.Before(time.Now())
}

type TokenCenter interface {
	RegisterToken(id string, t TokenItem)
	GetToken(id string) (t TokenItem, err error)
}
