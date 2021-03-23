package etc

type Configuration struct {
	Db      db     `toml:"db"`
	Web     web    `toml:"web"`
}

type web struct {
	Listen string `toml:"listen"`
}

type db struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DbName   string `toml:"database"`
	Ssl      string `toml:"ssl"`
}
