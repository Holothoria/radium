package main

import (
         "encoding/json"
         "log"
         "net/http"
         "strings"
         "fmt"
         neturl "net/url"
         "net"
         "os"
         )
        
func searchTMDB(query, contentType string) ([]SearchResult, error){
 apiKey := os.Getenv("TMDB_API_KEY")
 
 url := fmt.Sprintf(
   "https://api.themoviedb.org/3/search/%s?query=%s&api_key=%s",
   contentType,
   neturl.QueryEscape(query),
   apiKey,
   )
   
   transport := &http.Transport {
	   DialContext: (&net.Dialer{
		   DualStack: false,
	   }).DialContext,
   }
   
   client := &http.Client{Transport: transport}
   resp, err := client.Get(url)
   if err != nil {
	   return nil, err
   }
   
   defer resp.Body.Close()
   
   var result struct {
	   Results []struct {
		   ID                     int         `json:"id"`
		   Title                  string      `json:"title"`
		   Name                   string      `json:"name"`
		   Overview               string      `json:"overview"`
		   PosterPath             string      `json:"poster_path"`
           ReleaseDate            string      `json:"release_date"`
           FirstAirDate           string      `json:"first_air_date"`
	   } `json:"results"`
   }
   
   if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	   return nil, err
   }
   
   var movies []SearchResult
   for _, r := range result.Results {
	   title := r.Title
	   if title == "" {
		   title = r.Name 
	   }                                                       
	  year := r.ReleaseDate
	  if year == "" {
		  year = r.FirstAirDate
	  }
	  
	  if len(year) >= 4 {
		  year = year[:4]
	  }
	   movies = append(movies, SearchResult{
		   ID:             r.ID,
		   Title:          title,
		   Year:           year,
		   Poster:         "https://image.tmdb.org/t/p/w500" + r.PosterPath,
		   Type:           contentType,
	   })
   } 
   return movies, nil
} 
	   
	
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
 
 results, err := searchTMDB(query, contentType)
 if err != nil {
	 log.Printf("TMDB error: %v", err)
	 sendError(w, "search failed", http.StatusInternalServerError)
	 return
 }
 sendJSON(w, results, http.StatusOK)
}

//GET movie name

func handleMovie(w http.ResponseWriter, r*http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/movie/")
	if name =="" {
	 sendError(w, "movie name required", http.StatusBadRequest)
	 return
 }
// fetch from tmdb
 
 results, err := searchTMDB(name, "movie")
 if err != nil {
	 log.Printf("TMDB errro: %v", err )
	 sendError(w, "search failed", http.StatusInternalServerError)
	 return
 }
 sendJSON(w, results, http.StatusOK)
}

//FOR TV NAME FETCH
 
 func handleTV(w http.ResponseWriter, r*http.Request) {
	 name := strings.TrimPrefix(r.URL.Path, "/api/tv/")
	 if name == "" {
		 sendError(w, "tv show name required", http.StatusBadRequest) 
			 return
		 }
		 //fetch it from tmdb

results, err := searchTMDB(name, "tv")
if err != nil {
	log.Printf("TMDB error: %v", err)
	sendError(w, "search Failed", http.StatusInternalServerError)
	return
}
		 
sendJSON(w, results, http.StatusOK)
}
 
 
 //get anime name
 
 func handleAnime(w http.ResponseWriter, r*http.Request) {
	 name := strings.TrimPrefix(r.URL.Path, "/api/anime/")
	 if name == "" {
		 sendError(w, "anime name required", http.StatusBadRequest)
		 return 
	 }
  results, err := searchTMDB(name, "multi")
  if err != nil {
	  log.Printf("TMDB error: %v", err)
	  sendError(w, "search failed", http.StatusInternalServerError)
	  return
  }
  sendJSON(w, results, http.StatusOK)
} 	 
