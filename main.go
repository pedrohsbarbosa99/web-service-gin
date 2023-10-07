package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func removeAlbum(a []album, index int) []album {
	return append(a[:index], a[index+1:]...)
}

func getAlmbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album
	if err := context.BindJSON(&newAlbum); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "Invalid Payload",
				"message": err})
		return
	}
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(context *gin.Context) {
	id := context.Param("id")

	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}

	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deletAlbumById(context *gin.Context) {
	id := context.Param(("id"))

	for index, album := range albums {
		if album.ID == id {
			albums = removeAlbum(albums, index)
			context.JSON(http.StatusNoContent, "")
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlmbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deletAlbumById)
	router.Run("localhost:8080")
}
