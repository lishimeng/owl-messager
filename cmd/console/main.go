package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence/driver/postgres"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/cmd/console/ddd"
	"github.com/lishimeng/owl-messager/cmd/console/process"
	"github.com/lishimeng/owl-messager/cmd/console/static"
	"github.com/lishimeng/owl-messager/internal/etc"
)
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

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, nil)
		if err != nil {
			return err
		}
		dbConfig := postgres.Config{
			UserName:  etc.Config.Db.User,
			Password:  etc.Config.Db.Password,
			Host:      etc.Config.Db.Host,
			Port:      etc.Config.Db.Port,
			DbName:    etc.Config.Db.Database,
			InitDb:    true,
			AliasName: "default",
			SSL:       etc.Config.Db.Ssl,
		}

		builder.EnableDatabase(dbConfig.Build(),
			ddd.Tables()...).
			ComponentBefore(process.BeforeStarted).
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
