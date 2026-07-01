package msg

import (
	"reflect"
	"strings"
)

// ProviderCatalogEntry 通讯方式 + 平台 + 配置结构，系统唯一数据字典。
type ProviderCatalogEntry struct {
	Category MessageCategory
	Provider MessageProvider
	Config   any
}

var providerCatalog []ProviderCatalogEntry

func init() {
	providerCatalog = []ProviderCatalogEntry{
		{MailMessage, Smtp, SmtpConfig{}},
		{MailMessage, Microsoft, GraphConfig{}},
		{MailMessage, Tencent, TencentConfig{}},
		{SmsMessage, Ali, AliSmsConfig{}},
		{SmsMessage, Tencent, TencentSmsConfig{}},
		{SmsMessage, Huawei, HuaweiSmsConfig{}},
		{SmsMessage, Baidu, struct{}{}},
		{SmsMessage, QiNiu, struct{}{}},
		{ImMessage, FastMsg, FastMsgConfig{}},
		{ApnsMessage, Apns, ApnsConfig{}},
	}
	initProviderMaps()
}

func initProviderMaps() {
	MailProviders = make(map[MessageProvider]byte)
	SmsProviders = make(map[MessageProvider]byte)
	ImProviders = make(map[MessageProvider]byte)
	ApnsProviders = make(map[MessageProvider]byte)

	for _, e := range providerCatalog {
		switch e.Category {
		case MailMessage:
			MailProviders[e.Provider] = VendorEnable
		case SmsMessage:
			if e.Provider == UpYun {
				continue
			}
			SmsProviders[e.Provider] = VendorEnable
		case ImMessage:
			ImProviders[e.Provider] = VendorEnable
		case ApnsMessage:
			ApnsProviders[e.Provider] = VendorEnable
		}
	}

	Providers = make(map[MessageCategory]map[MessageProvider]byte)
	Providers[MailMessage] = MailProviders
	Providers[SmsMessage] = SmsProviders
	Providers[ImMessage] = ImProviders
	Providers[ApnsMessage] = ApnsProviders
}

// ListProviders 返回某通讯方式下支持的 provider（顺序与目录一致）。
func ListProviders(category MessageCategory) []MessageProvider {
	var out []MessageProvider
	seen := make(map[MessageProvider]bool)
	for _, e := range providerCatalog {
		if e.Category != category || seen[e.Provider] {
			continue
		}
		seen[e.Provider] = true
		out = append(out, e.Provider)
	}
	return out
}

// ListCatalog 返回目录条目，category 为空时返回全部。
func ListCatalog(category MessageCategory) []ProviderCatalogEntry {
	if category == "" {
		return append([]ProviderCatalogEntry(nil), providerCatalog...)
	}
	var out []ProviderCatalogEntry
	for _, e := range providerCatalog {
		if e.Category == category {
			out = append(out, e)
		}
	}
	return out
}

// ProviderConfig 返回 (category, provider) 对应的配置结构原型。
func ProviderConfig(category MessageCategory, provider MessageProvider) (any, bool) {
	p := NormalizeProvider(provider)
	for _, e := range providerCatalog {
		if e.Category == category && e.Provider == p {
			return e.Config, true
		}
	}
	return nil, false
}

// ProviderConfigFields 返回配置 JSON 字段名 -> Go 类型名。
func ProviderConfigFields(category MessageCategory, provider MessageProvider) map[string]string {
	cfg, ok := ProviderConfig(category, provider)
	if !ok {
		return nil
	}
	return configJSONFields(cfg)
}

// NormalizeProvider 将历史别名归一为目录中的 provider 值。
func NormalizeProvider(provider MessageProvider) MessageProvider {
	switch provider {
	case "tencent_yun":
		return Tencent
	default:
		return provider
	}
}

func configJSONFields(v any) map[string]string {
	m := make(map[string]string)
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return m
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag, ok := f.Tag.Lookup("json")
		if !ok {
			continue
		}
		name := strings.Split(tag, ",")[0]
		if name == "" || name == "-" {
			continue
		}
		m[name] = f.Type.Name()
	}
	return m
}
