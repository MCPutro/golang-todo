package database

import (
	"fmt"
	"github.com/MCPutro/golang-todo/config"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestInitConnection(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	config.DbHost = os.Getenv("MYSQL_HOST")
	config.DbPass = os.Getenv("MYSQL_PASSWORD")
	config.DbUser = os.Getenv("MYSQL_USER")
	config.DbName = os.Getenv("MYSQL_DBNAME")
	config.DbPort = os.Getenv("MYSQL_PORT")

	instance, err := GetInstance()

	fmt.Println("1 >>", err)
	fmt.Println("2 >>", instance)
}
