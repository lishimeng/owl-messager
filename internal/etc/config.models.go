package etc

type Configuration struct {
	Db  db
	Web web
}

type web struct {
	Listen string
}

type db struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Ssl      string
}
