package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
)

func orm() *persistence.OrmContext {
	return app.GetOrm()
}

func create(value interface{}) error {
	return app.Transaction(func(tx persistence.TxContext) error {
		return tx.Create(value)
	})
}

func updateSelect(ctx persistence.TxContext, m interface{}, cols ...string) error {
	if len(cols) == 0 {
		return nil
	}
	args := make([]interface{}, len(cols)-1)
	for i, c := range cols[1:] {
		args[i] = c
	}
	return ctx.Model(m).Select(cols[0], args...).Updates(m)
}
