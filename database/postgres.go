package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func OpenPg() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", viper.GetString("postgres.user"), viper.GetString("postgres.password"), viper.GetString("base.host"), viper.GetString("postgres.port"), viper.GetString("postgres.database"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	return db
}
