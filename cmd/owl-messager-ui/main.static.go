package main

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/cmd"
	"github.com/lishimeng/owl/internal/etc"
	"github.com/lishimeng/owl/static"
	"time"
)
import _ "github.com/lib/pq"

func main() {
	orm.Debug = true

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(cmd.AppName)
	fmt.Println(cmd.Version)

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 2)
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

		builder.SetWebLogLevel("debug").
			EnableWeb(etc.Config.Web.Listen).
			EnableStaticWeb(
				"ui/dist",
				"index.html",
				static.AssetInfo,
				static.Asset,
				static.AssetNames)
		return err
	}, func(s string) {
		log.Info(s)
	})
	return
}
