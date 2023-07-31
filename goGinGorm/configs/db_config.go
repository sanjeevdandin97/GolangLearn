package configs

type DBConfigStruct struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var DBConfig = DBConfigStruct{
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "1234",
	DBName:   "db_1",
}
