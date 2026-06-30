package etc

type Configuration struct {
	Db             db             `toml:"db"`
	Web            web            `toml:"web"`
	MailAttachment mailAttachment `toml:"mail-attachment"`
	LogLevel       string         `toml:"log-level"`
	Sqlite         sqlite         `toml:"sqlite"`
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
