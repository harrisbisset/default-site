package host

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/gofor-little/env"
)

func loadConfig() mysql.Config {

	if err := env.Load("./internal/site/.env"); err != nil {
		log.Panic(err)
	}

	var (
		host     = env.Get("DB_SITE_HOST", "")
		port     = env.Get("DB_SITE_PORT", "")
		user     = env.Get("DB_SITE_ROOT", "")
		password = env.Get("DB_SITE_PASS", "")
		dbname   = env.Get("DB_SITE_NAME", "")
	)

	var cfg = mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 host + ":" + port,
		DBName:               dbname,
		AllowNativePasswords: true,
	}

	return cfg
}

func getDatabase(cfg mysql.Config) *sql.DB {

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to database!")

	return db
}
