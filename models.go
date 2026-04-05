package main






//Movie represents a movie or Tv show
;
type Movie struct {
	ID                 int             `json:"id"`
	Title              string          `json:"title"`
	Description        string          `json:"description"`
	Rating             string          `json:"rating"`
	Year               string          `json:"year"`
	Poster             string          `json:"poster"`
	Type               string          `json:"type"`  // anythng it can be (movie,tv shows or {anime for  now} )
	TMDBID             int             `json:"tmdb_id"` 
}

//for anime

type Anime struct (
	ID                 int             `json:"id"`
	Title              string          `json:"title"`
	Description        string          `json:"description"`
	Rating             string          `json:"rating"`
	Year               string          `json:"year"`
	Poster             string          `json:"poster"`
	Episodes           int             `json:"episodes"`
    MalID              int             `json:"mal_id"`
}

//search results for suggesstions dropdown

type SearchResult struct {
    Title             string           `json:"title"`
    Year              string           `json:"year"`
    Poster            string           `json:"poster"`
    Type              string           `json:"type"` // anything provided by you
    ID                int              `json:"id"`
}

//stream source is for your streaming source

type StreamSource struct {
	Name            string             `json:"name"`
	URL             string             `json:"url"`
}
