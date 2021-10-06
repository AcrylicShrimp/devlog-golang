package env

func InitEnvVars() {
	InitAuthEnvVars()
	InitAWSEnvVars()
	InitDBEnvVars()
	InitServerEnvVars()
}
