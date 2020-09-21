package variable

import (
	"os"

	"github.com/joho/godotenv"
)

// MinioConfig Variable
var MinioConfig struct {
	EndPoint  string
	AccessKey string
	SecretKey string
}

// WeeberStoragesConfig Variable
var WeeberStoragesConfig struct {
	URI string
}

// Initialization for getting variable environment
func Initialization() {
	godotenv.Load("devel.env")

	MinioConfig.EndPoint = os.Getenv("ENDPOINT")
	MinioConfig.AccessKey = os.Getenv("ACCESS_KEY")
	MinioConfig.SecretKey = os.Getenv("SECRET_KEY")

	WeeberStoragesConfig.URI = os.Getenv("WEEBER_STORAGES_URI")
}
