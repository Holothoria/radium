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
	
	http.Handler("/static",
	 http.StripPrefix("/static/",
	  http.FileServer(http.Dir("static"))))
	  
	  //API routes
	  
	  http.Handler("/api/search", handleSearch)
	  http.Handler("/api/movie/", handleMovie)
	  http.Handler("api/tv/"    , handleTv)
	  http.Handler ("/api/anime/", handleAnime)
	  
	  //Frontend
	  
	  http.Handler("/", handleIndex)
	  
	  //Port from env or default
	  
	  port := os.Getenv("PORT")
	  if port == "" {
		  port = "5000"
	  }
	  
	  log.printf("Radium running on :%s", port )
	  if err := http.ListenAndServe(":" + port, nil); err != nil {
		  log.Fatalf("Server failed to start: %v", err)
	  }
  }
	
