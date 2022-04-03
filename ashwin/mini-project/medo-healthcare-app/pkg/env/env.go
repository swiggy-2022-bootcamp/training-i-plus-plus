package env

import (
	"medo-healthcare-app/pkg/err"
	"os"

	"github.com/joho/godotenv"
)

//GoDotEnvVariable ...
func GoDotEnvVariable(key string) string {
	cancel := godotenv.Load(".env")
	err.CheckNilErr(cancel)
	return os.Getenv(key)
}
