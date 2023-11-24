package postgresql

type Config struct {
	Host     string `json:"host`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
}