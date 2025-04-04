package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const DBKindMySql = "mysql"

var (
	MySqlDbHost string
	MySqlDbName string
	MySqlDbPass string
	MySqlDbPort int
	MySqlDbUser string

	HttpPort         int
	HttpReadTimeout  int
	HttpWriteTimeout int

	CorsAllowedOrigins string
	Environment        string
	Version            string
	JwtSecretKey       string
)

func LoadEnvs() {
	godotenv.Load()

	MySqlDbHost = os.Getenv("DB_MYSQL_HOST")
	MySqlDbName = os.Getenv("DB_MYSQL_NAME")
	MySqlDbPass = os.Getenv("DB_MYSQL_PASS")
	MySqlDbPort, _ = strconv.Atoi(os.Getenv("DB_MYSQL_PORT"))
	MySqlDbUser = os.Getenv("DB_MYSQL_USER")

	HttpPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))
	HttpReadTimeout, _ = strconv.Atoi(os.Getenv("HTTP_READ_TIMEOUT"))
	HttpWriteTimeout, _ = strconv.Atoi(os.Getenv("HTTP_WRITE_TIMEOUT"))

	CorsAllowedOrigins = os.Getenv("CORS_ALLOWED_ORIGINS")
	Environment = os.Getenv("ENV")
	Version = os.Getenv("VERSION")
}
