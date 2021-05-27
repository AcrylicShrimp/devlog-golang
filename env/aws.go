package env

import "os"

var AWSS3Region string
var AWSS3Bucket string

func InitAWSEnvVars() {
	AWSS3Region = os.Getenv("AWS_S3_REGION")
	AWSS3Bucket = os.Getenv("AWS_S3_BUCKET")
}
