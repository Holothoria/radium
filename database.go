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

func createTables() error {
	tables := []string{
	`CREATE TABLE IF NOT EXISTS movies (
	
	 id                          INTEGER                 PRIMARY KEY        AUTOINCREMENT,
	 tmdb_id                     INTEGER                 UNIQUE,
	 title                       TEXT,
	 description                 TEXT,
	 rating                      TEXT,
	 year                        TEXT,
	 poster                      TEXT,
	 type                        TEXT
	)`,
	
	`CREATE TABLE IF NOT EXISTS anime (
	
     id                          INTEGER                 PRIMARY KEY        AUTOINCREMENT,
	 mal_id                      INTEGER                 UNIQUE,
	 title                       TEXT,
	 description                 TEXT,
	 rating                      TEXT,
	 year                        TEXT,                                            
	 poster                      TEXT,
	 episodes                    INTEGER
	 )`, 
 }
  for _, q := range tables {
	  if _, err := db.Exec(q); err != nil {
		  return err 
 }
}
 return nil
}
 
 //save movie to cache
 
 func saveMovieToDB(m Movie) error {
	 _, err := db.Exec(`
	 
	 INSERT OR REPLACE INTO movies
	 (tmdb_id, title, description, rating, year, poster, type)
	 VALUES (?,?,?,?,?,?,?)`,
	 m.TMDBID, m.Title, m.Description, m.Rating, m.Year, m.Poster, m.Type)
	 
	 if err != nil {
		 log.Printf("DB save error: %v", err)
	 }
	 return err
 }
 
 //Request movie from cache
 
 func getMovieFromDB(tmdbID int) (Movie, error){
	 var m Movie
	 err := db.QueryRow(`
	 SELECT tmdb_id, title, description, rating, year, poster, type,
	 FROM movies WHERE tmdb_id = ?`, tmdbID).
	 Scan(&m.TMDBID, &m.Title, &m.Description, &m.Rating, &m.Year, &m.Poster, &m.Type)
	 return m, err
 }	

