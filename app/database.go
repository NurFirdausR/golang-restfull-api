package app

import (
	"database/sql"
	"time"

	"github.com/NurFirdausR/golang-restfull-api/helper"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:Bismillah@123@tcp(localhost:3306)/golang_api")

	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
