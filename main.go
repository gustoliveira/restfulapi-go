package main

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	fmt.Println("Hello world API - Jazz Music")

	router := gin.Default()

	router.GET("/albums", getAlbums)

	router.POST("/albums", addAlbum)

	router.GET("/albums/:id", getAlbum)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	index := slices.IndexFunc(albums, func(a album) bool {
		return a.ID == string(id)
	})

	if index != -1 {
		c.IndentedJSON(http.StatusOK, albums[index])
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album not found"})
}

func addAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, albums)
}
