package main

import (
         "encoding/json"
         "log"
         "net/http"
         "strings"
         )
         
//helper - sends json response

func sendJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("JSON encode error: %v", err)
	}
}

//helper - send error message 
func sendError(w http.ResponseWriter, message string, status int) {
	sendJSON(w, map[string]string{"error": message}, status)
}

//Get
func handleIndex(w http.ResponseWriter, r*http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

//Get f(for searching)
func handleSearch(w http.ResponseWriter, r*http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query =="" {
		sendError(w, "query required", http.StatusBadRequest)
		
	}
	
	contentType := r.URL.Query().Get("type")
	if contentType == "" {
	 contentType = "movie"
 }
 
 //call your api i.e. tmdb or jikan, on type
 
 sendJSON(w, map[string]string{"status": "ok", "query": query}, http.StatusOK)
}

//GET movie name

func handleMovie(w http.ResponseWriter, r*http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/movie/")
	if name =="" {
	 sendError(w, "movie name required", http.StatusBadRequest)
	 return
 }
// fetch from tmdb
 
 sendJSON(w, map[string]string{"status": "ok", "name": name }, http.StatusOK)
}

//FOR TV NAME FETCH
 
 func handleTV(w http.ResponseWriter, r*http.Request) {
	 name := strings.TrimPrefix(r.URL.Path, "/api/tv/")
	 if name == "" {
		 sendError(w, "tv show name required", http.StatusBadRequest) 
			 return
		 }
		 //fetch it from tmdb
		 
		 sendJSON(w, map[string]string{"status": "ok", "name": name}, http.StatusOK)
	 }
 
 
 //get anime name
 
 func handleAnime(w http.ResponseWriter, r*http.Request) {
	 name := strings.TrimPrefix(r.URL.Path, "api/anime/")
	 if name == "" {
		 sendError(w, "anime name required", http.StatusBadRequest)
		 return 
	 }
	 
	 sendJSON(w, map[string]string{"status": "ok", "name": name}, http.StatusOK)
 }
