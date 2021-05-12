package openapi

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/lishimeng/owl/internal/etc"
	"log"
)
import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	oredis "github.com/go-oauth2/redis/v4"
)

type Req struct {
	AccessToken  string
	RefreshToken string
	AppKey       string
	AppSecret    string
}

func Init(ctx context.Context) {

	manager := manage.NewDefaultManager()
	clientStore := NewClientStore()

	manager.MapClientStorage(clientStore)
	manager.MustTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr:               etc.Config.Redis.Addr,
	}), nil)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
}