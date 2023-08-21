package etc

type Configuration struct {
	Db      db      `toml:"db"`
	Web     web     `toml:"web"`
	Token   token   `toml:"token"`
	Redis   redis   `toml:"redis"`
	Console console `toml:"console"`
	Sender  sender  `toml:"sender"`
}

type web struct {
	Listen string `toml:"listen"`
}

type redis struct {
	Enable   bool   `toml:"enable"`
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	Db       int    `toml:"db"`
}

type token struct {
	Enable bool   `toml:"enable"`
	Issuer string `toml:"issuer"`
	Key    string `toml:"key"`
}

type sender struct {
	Strategy int `toml:"strategy"` // task.Strategy
	Buff     int `toml:"buff"`
}

type db struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Database string `toml:"database"`
	Ssl      string `toml:"ssl"`
}

// console独有的配置
type console struct {
	TokenProvider string `toml:"tokenProvider"`
}
