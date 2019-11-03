package db

import "database/sql"

var db *sql.DB
var err error

func database() {
	db, err = sql.Open("mysql",   "<user>:<password>@tcp(127.0.0.1:3306)/<dbname>")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()