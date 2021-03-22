package api

import (
	"encoding/json"
	"github.com/kataras/iris"
)

func responseJSON(ctx iris.Context, j interface{}) {
	bs, _ := json.Marshal(j)
	_, _ = ctx.Write(bs)
}
