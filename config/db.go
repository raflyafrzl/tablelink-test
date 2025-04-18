package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() *sqlx.DB {
	var err error
	var db *sqlx.DB
	// Should be hide inside configmap/env variable
	db, err = sqlx.Connect("postgres", "user=raflyafrzl password=raflyafrzl dbname=test_tablelink sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
