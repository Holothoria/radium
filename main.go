package main

import (
        "log"
        "net/http"
        "os"
        
        _"github.com/mattn/go-sqlite3"
)

func main () {
	//initialized the database
	if err := initDB(); err != nil {
		log.Fatalf("Failed to init DB: %v", err)
	}
	
	//static files
	
	http.Handle("/static",
	 http.StripPrefix("/static/",
	  http.FileServer(http.Dir("static"))))
	  
	  //API routes
	  
	  http.HandleFunc("/api/search", handleSearch)
	  http.HandleFunc("/api/movie/", handleMovie)
	  http.HandleFunc("api/tv/"    , handleTV)
	  http.HandleFunc ("/api/anime/", handleAnime)
	  
	  //Frontend
	  
	  http.HandleFunc("/", handleIndex)
	  
	  //Port from env or default
	  
	  port := os.Getenv("PORT")
	  if port == "" {
		  port = "5000"
	  }
	  
	  log.Printf("Radium running on :%s", port )
	  if err := http.ListenAndServe(":" + port, nil); err != nil {
		  log.Fatalf("Server failed to start: %v", err)
	  }
  }
	
