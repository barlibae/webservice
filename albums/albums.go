package albums

import (
	"sync"

	"github.com/google/uuid"
)

// run go install github.com/golang/mock/mockgen@v1.6.0 to install mockgen
//go:generate mockgen -source=albums.go -destination=mocks/albums.go -package mocks

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type AlbumRepo interface {
	CreateAlbum(album Album) Album
	Albums() []Album
	AlbumByID(id string) (Album, bool)
}

var demoAlbums = map[string]Album{
	"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type inMemoryRepo struct {
	albums *sync.Map
}

func NewAlbumRepo() AlbumRepo {
	albumsMap := sync.Map{}

	for k, v := range demoAlbums {
		albumsMap.Store(k, v)
	}

	return &inMemoryRepo{
		albums: &albumsMap,
	}
}

func (r *inMemoryRepo) CreateAlbum(album Album) Album {
	album.ID = uuid.New().String()
	r.albums.Store(album.ID, album)
	return album
}

func (r *inMemoryRepo) Albums() []Album {
	albums := make([]Album, 0)
	r.albums.Range(func(_, value any) bool {
		albums = append(albums, value.(Album))
		return true
	})
	return albums
}

func (r *inMemoryRepo) AlbumByID(id string) (Album, bool) {
	a, found := r.albums.Load(id)
	if !found {
		return Album{}, false
	}
	return a.(Album), true
}
