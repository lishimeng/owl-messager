package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/persistence/driver/postgres"
	"github.com/lishimeng/app-starter/persistence/driver/sqlite"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/process"
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

		if len(etc.Config.LogLevel) > 0 {
			var lvl log.Level
			lvl, err = log.FormatLevel(etc.Config.LogLevel)
			if err != nil {
				return err
			}
			log.SetLevelAll(lvl)
		}
		var dbConfig persistence.BaseConfig
		if len(etc.Config.Db.Host) > 0 {
			c := postgres.Config{
				UserName:  etc.Config.Db.User,
				Password:  etc.Config.Db.Password,
				Host:      etc.Config.Db.Host,
				Port:      etc.Config.Db.Port,
				DbName:    etc.Config.Db.Database,
				InitDb:    true,
				AliasName: "default",
				SSL:       etc.Config.Db.Ssl,
			}
			dbConfig = c.Build()
		} else if len(etc.Config.Sqlite.Db) > 0 {
			c := sqlite.Config{
				Database:  etc.Config.Sqlite.Db,
				AliasName: "default",
				InitDb:    true,
			}
			dbConfig = c.Build()
		} else {
			panic("no db config")
		}

		builder.EnableDatabase(dbConfig,
			ddd.Tables()...).
			PrintVersion().
			EnableDatabaseLog().
			EnableWeb(etc.Config.Web.Listen, ddd.Route).
			ComponentBefore(process.BeforeStarted).
			ComponentAfter(process.AfterStarted)

		return err
	}, func(s string) {
		log.Info(s)
	})

	return
}
