package main

import (
	"github.com/barlibae/webservice/albums"
	"github.com/gin-gonic/gin"
	"net/http"
)

func newRouter(repo albums.AlbumRepo) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.GET("/albums", getAlbums(repo))
	ginRouter.GET("/albums/:id", getAlbumByID(repo))
	ginRouter.POST("/albums", postAlbums(repo))

	return ginRouter
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(albumRepo albums.AlbumRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, albumRepo.Albums())
	}
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(albumRepo albums.AlbumRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newAlbum albums.Album

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}

		created := albumRepo.CreateAlbum(newAlbum)
		c.IndentedJSON(http.StatusCreated, created)
	}
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(albumRepo albums.AlbumRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		a, found := albumRepo.AlbumByID(id)
		if !found {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}

		c.IndentedJSON(http.StatusOK, a)
	}
}
