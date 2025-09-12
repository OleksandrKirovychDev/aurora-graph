package config

import "os"

var (
	AccountUrl     string
	SecretKey      string
	Issuer         string
)

func init() {
	AccountUrl = os.Getenv("ACCOUNT_SERVICE_URL")
	SecretKey = os.Getenv("SECRET_KEY")
	Issuer = os.Getenv("ISSUER")
}
