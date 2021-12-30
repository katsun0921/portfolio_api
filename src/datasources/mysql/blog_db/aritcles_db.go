package blog_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Client *sql.DB
)

func init() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("MYSQL_BLOG_USERNAME")
	password := os.Getenv("MYSQL_BLOG_PASSWORD")
	host := os.Getenv("MYSQL_BLOG_HOST")
	schema := os.Getenv("MYSQL_BLOG_SCHEMA")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	log.Println(fmt.Sprintf("about to connect to %s", dataSourceName))
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
