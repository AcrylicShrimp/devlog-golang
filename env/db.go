package env

import "os"

var DatabaseDriver = os.Getenv("DB_DRIVER")
var DatabaseDSN = os.Getenv("DB_DSN")
