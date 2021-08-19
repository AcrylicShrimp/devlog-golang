package env

import "os"

var Port string

func InitServerEnvVars() {
	Port = os.Getenv("PORT")
}
