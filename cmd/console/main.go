package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/cmd/console/ddd"
	"github.com/lishimeng/owl-messager/cmd/console/static"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/etc"
	"net/http"
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

		// console的token验证器使用http方式,统一由外部管理,比如passport
		builder = builder.EnableTokenValidator(func(inject app.TokenValidatorInjectFunc) {
			provider := token.HttpStorageConnector{Server: etc.Config.Console.TokenProvider}
			storage := token.NewHttpStorage(provider)
			inject(storage)
		})

		builder.EnableDatabase(dbConfig.Build(),
			model.Tables()...).
			EnableStaticWeb(func() http.FileSystem {
				return http.FS(static.Static)
			}).
			PrintVersion().
			EnableWeb(etc.Config.Web.Listen, ddd.Route)

		return err
	}, func(s string) {
		log.Info(s)
	})

	return
}
