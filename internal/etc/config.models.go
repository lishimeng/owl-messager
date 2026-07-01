package etc

type Configuration struct {
	Db             db             `toml:"db"`
	Web            web            `toml:"web"`
	Messager       messager       `toml:"messager"`
	MailAttachment mailAttachment `toml:"mail-attachment"`
	LogLevel       string         `toml:"log-level"`
	Sqlite         sqlite         `toml:"sqlite"`
}

type messager struct {
	// Host owl-messager Open API 根地址，Console 测试发送默认转发目标，如 http://127.0.0.1:81
	Host string `toml:"host"`
}

type web struct {
	Listen string `toml:"listen"`
}

type sqlite struct {
	Db string `toml:"db"`
}

type mailAttachment struct {
	Dir             string `toml:"dir"`
	MaxFileSize     int64  `toml:"maxFileSize"`
	MaxTotalSize    int64  `toml:"maxTotalSize"`
	MaxCount        int    `toml:"maxCount"`
	StagingTTL      string `toml:"stagingTTL"`
	RetainAfterSend string `toml:"retainAfterSend"`
}

type db struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Database string `toml:"database"`
	Ssl      string `toml:"ssl"`
}
