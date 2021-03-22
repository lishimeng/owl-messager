package main

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl/internal/api"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/etc"
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

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 2)
}

func _main() (err error) {
	configName := "config.toml"

	application := app.New()
	

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, configName, ".")
		if err != nil {
			return err
		}
		dbConfig := persistence.PostgresConfig{
			UserName:  etc.Config.Db.User,
			Password:  etc.Config.Db.Password,
			Host:      etc.Config.Db.Host,
			Port:      etc.Config.Db.Port,
			DbName:    etc.Config.Db.DbName,
			InitDb:    false,
			AliasName: "default",
			SSL: etc.Config.Db.Ssl,
		}
		log.Info("start [%s]", etc.Config.Name)

		builder.EnableDatabase(dbConfig.Build(),
			new(model.MessageInfo),
			new(model.MailMessageInfo),
			new(model.SmsMessageInfo),
			new(model.MailSenderInfo),
			new(model.MessageTask),
			new(model.MessageRunningTask)).
			EnableWeb(etc.Config.Web.Listen, api.Route)//.
			//ComponentBefore(setup.JobClearExpireTask).
			//ComponentBefore(setup.MessageSender)
		return err
	}, func(s string) {
		log.Info(s)
	})
	return
}