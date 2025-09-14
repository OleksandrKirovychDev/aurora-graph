package config

import "os"

var (
	AccountUrl     string
	SecretKey      string
	Issuer         string
)

func Init() {
	AccountUrl = os.Getenv("ACCOUNT_URL")
	SecretKey = os.Getenv("SECRET_KEY")
	Issuer = os.Getenv("ISSUER")
}
