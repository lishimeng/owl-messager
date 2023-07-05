package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/process"
	"github.com/lishimeng/owl-messager/internal/etc"
	"time"
)
import _ "github.com/lib/pq"
import _ "github.com/lishimeng/owl-messager/providers"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 50)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}
		dbConfig := persistence.PostgresConfig{
			UserName:  etc.Config.Db.User,
			Password:  etc.Config.Db.Password,
			Host:      etc.Config.Db.Host,
			Port:      etc.Config.Db.Port,
			DbName:    etc.Config.Db.Database,
			InitDb:    true,
			AliasName: "default",
			SSL:       etc.Config.Db.Ssl,
		}

		issuer := etc.Config.Token.Issuer
		tokenKey := []byte(etc.Config.Token.Key)
		builder = builder.EnableTokenValidator(func(inject app.TokenValidatorInjectFunc) {
			provider := token.NewJwtProvider(issuer,
				token.WithKey(tokenKey, tokenKey), // hs256的秘钥必须是[]byte
				token.WithAlg("HS256"),
				token.WithDefaultTTL(etc.TokenTTL),
			)
			storage := token.NewLocalStorage(provider)
			factory.Add(provider)
			inject(storage)
		})

		builder.EnableDatabase(dbConfig.Build(),
			ddd.Tables()...).
			PrintVersion().
			EnableWeb(etc.Config.Web.Listen, ddd.Route).
			ComponentBefore(process.BeforeStarted).
			ComponentAfter(process.AfterStarted)

		return err
	}, func(s string) {
		log.Info(s)
	})

	return
}
