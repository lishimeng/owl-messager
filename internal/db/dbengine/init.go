package dbengine

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
)

func GetDefault() *persistence.OrmContext {
	return app.GetOrm()
}
