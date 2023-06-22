package main

import (
	"github.com/barlibae/webservice/albums"
)

func main() {
	repo := albums.NewAlbumRepo()
	router := newRouter(repo)
	router.Run("localhost:8080")
}
