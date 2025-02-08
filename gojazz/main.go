package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// add albums with post request
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call bindJSON to bind Json received to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// add album to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GET creates JSON from slice of album structs, writing JSON to response
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// retrieve album by id
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over albums
	// for loops in go return the index and value
	// we dont want the index but the value
	for _, a := range albums {
		// check if id matches
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// albums slice for data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// sets up association which getAlbums handles requests to /albums endpoint
func main() {
	router := gin.Default()
	// pass in ref to the func, not the result (getAlbums())
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	// Start the server and listen on port 8080
	router.Run("localhost:8080")
}
