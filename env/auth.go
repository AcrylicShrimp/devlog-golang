package env

import "os"

var JWTSigningKey string

func InitAuthEnvVars() {
	JWTSigningKey = os.Getenv("JWT_SIGNING_KEY")

	if JWTSigningKey == "" {
		panic("the env var JWT_SIGNING_KEY must be set")
	}
}
