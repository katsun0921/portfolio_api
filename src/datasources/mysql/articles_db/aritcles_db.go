package articles_db

import (
  "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	Client *sql.DB

	username = os.Getenv("MYSQL_BLOG_USER")
	password = os.Getenv("MYSQL_BLOG_PASSWORD")
	host     = os.Getenv("MYSQL_BLOG_HOST")
	schema   = os.Getenv("MYSQL_BLOG_SCHEMA")
)

func init() {
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
