package env

import "os"

var AWSS3Region = os.Getenv("AWS_S3_REGION")
var AWSS3Bucket = os.Getenv("AWS_S3_BUCKET")
