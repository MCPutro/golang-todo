package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DbUser = mustGetEnv("MYSQL_USER")     //postgres
	DbPass = mustGetEnv("MYSQL_PASSWORD") //password
	DbHost = mustGetEnv("MYSQL_HOST")     //localhost
	DbName = mustGetEnv("MYSQL_DBNAME")   //postgres
	DbPort = mustGetEnv("MYSQL_PORT")     //5432
)

func mustGetEnv(k string) string {
	/* if run in localhost will load from variable(file) environment */
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}

	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.", k)
	}
	return v
}
