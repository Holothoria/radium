package main

import (
	"database/sql"
	"log"
	
	_"github.com/mattn/go-sqlite3"
)

var db *sql.DB 

func initDB() error {
	var err error 
	db, err = sql.Open("sqlite3", "./radium.db")
	if err != nil {
		return err
	}
	
	//Test connection
	
	if err = db.Ping(); err != nil {
		return err
	}
	
	//Create Tables
	
	if err = createTables(); err != nil {
		return err
	}
	
	log.Printf("Database started!")
	return nil
}

