package statisticsApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/consoleorg"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"time"
)

type providersStat struct {
	Provider string `json:"provider"`
	Count    int    `json:"count"`
}

type dailyStat struct {
	Date  string `json:"date"`
	Mail  int    `json:"mail"`
	Sms   int    `json:"sms"`
	Im    int    `json:"im"`
	Total int    `json:"total"`
}

type respProvidersStat struct {
	app.Response
	Providers map[string][]providersStat `json:"providers"`
}

type respDailyStat struct {
	app.Response
	Stats []dailyStat `json:"stats"`
}

func GetProvidersStat(ctx server.Context) {
	var resp respProvidersStat
	var stat []model.ProviderStats

	orgID := consoleorg.ID(ctx)
	err := app.GetOrm().Model(&model.ProviderStats{}).
		Equal("org", orgID).
		Find(&stat)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	result := make(map[string][]providersStat)
	for _, item := range stat {
		c, exists := result[item.Category.String()]
		if !exists {
			c = make([]providersStat, 0)
		}
		result[item.Category.String()] = append(c, providersStat{item.Provider.String(), item.Value})
	}
	resp.Providers = result
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func GetDailyStat(ctx server.Context) {
	var resp respDailyStat

	startString := ctx.C.URLParam("date")
	batch := ctx.C.URLParamIntDefault("batch", 30)
	minBatch := ctx.C.URLParamIntDefault("min", 0)

	startDate, err := time.Parse("2006-01-02", startString)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	startDate = startDate.Local()

	orgID := consoleorg.ID(ctx)
	var earliest model.DailySummary
	err = app.GetOrm().Model(&model.DailySummary{}).
		Equal("org", orgID).
		Order("date").
		First(&earliest)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	diff := int(startDate.Sub(earliest.Date).Hours()/24) + 1
	if batch > diff {
		batch = diff
	}
	if batch < minBatch {
		batch = minBatch
	}

	endDate := startDate.AddDate(0, 0, -batch)

	result := make([]dailyStat, batch)
	var records []model.DailySummary
	err = app.GetOrm().Model(&model.DailySummary{}).
		Equal("org", orgID).
		Where("date > ?", endDate).
		Where("date <= ?", startDate).
		Order("-date").
		Find(&records)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	for i, j := 0, 0; i < batch; startDate = startDate.AddDate(0, 0, -1) {
		result[i].Date = startDate.Format("2006-01-02")
		if j < len(records) && startDate.Equal(records[j].Date) {
			result[i].Im = records[j].Im
			result[i].Sms = records[j].Sms
			result[i].Mail = records[j].Mail
			result[i].Total = result[i].Im + result[i].Sms + result[i].Mail
			j++
		}
		i++
	}

	resp.Stats = result
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
