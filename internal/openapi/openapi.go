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
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
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
	clientStore := store.NewClientStore()
	_ = clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	}) // TODO replace the client store: wait implementation of postgres
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