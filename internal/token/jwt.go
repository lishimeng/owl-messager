package token

import (
	"fmt"
	proxy "github.com/golang-jwt/jwt/v4"
	"time"
)

type BaseToken struct {
	UID  int   `json:"uid"`  // id
	OID  int   `json:"oid"`  // organization
	Type int32 `json:"type"` // 登录方式
}

type Req struct {
	BaseToken
	Audience string
	Subject  string
	Expire   time.Duration
}

type Claims struct {
	BaseToken
	proxy.RegisteredClaims
}

type Handler struct {
	key    []byte        // key
	expire time.Duration // 有效时间
	issuer string
}

func New(key []byte, issuer string, expire time.Duration) Handler {

	return Handler{key: key, issuer: issuer, expire: expire}
}

func (h *Handler) GenToken(t Req) (claims *Claims, expire time.Duration, signedToken string, success bool) {

	claims = &Claims{
		BaseToken: t.BaseToken,
		RegisteredClaims: proxy.RegisteredClaims{
			Issuer: h.issuer,
		},
	}
	if len(t.Audience) > 0 {
		claims.RegisteredClaims.Audience = []string{t.Audience}
	}
	if len(t.Subject) > 0 {
		claims.RegisteredClaims.Subject = t.Subject
	}
	if t.Expire > 0 {
		expire = t.Expire
	} else {
		expire = h.expire
	}
	claims.ExpiresAt = proxy.NewNumericDate(time.Now().Add(expire))
	signedToken, success = h.CreateToken(claims)
	return
}

func (h *Handler) VerifyToken(signedToken string) (claims *Claims, success bool) {

	claims, success = h.ValidateToken(signedToken)
	if success {
		success = claims.VerifyIssuer(h.issuer, false)
	}
	if success {
		success = claims.VerifyExpiresAt(time.Now(), true)
	}
	return
}

// CreateToken create tokenApi
func (h *Handler) CreateToken(claims *Claims) (signedToken string, success bool) {
	token := proxy.NewWithClaims(proxy.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(h.key)
	if err != nil {
		return
	}

	success = true
	return
}

// ValidateToken validate tokenApi
func (h *Handler) ValidateToken(signedToken string) (claims *Claims, success bool) {
	token, err := proxy.ParseWithClaims(signedToken, &Claims{},
		func(token *proxy.Token) (interface{}, error) {
			if _, ok := token.Method.(*proxy.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return h.key, nil
		})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		success = true
		return
	}

	return
}
