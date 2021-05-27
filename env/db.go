package env

import "os"

var DatabaseDriver string
var DatabaseDSN string

func InitDBEnvVars() {
	DatabaseDriver = os.Getenv("DB_DRIVER")
	DatabaseDSN = os.Getenv("DB_DSN")
}
