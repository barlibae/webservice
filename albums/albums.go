package albums

// run go install github.com/golang/mock/mockgen@v1.6.0 to install mockgen
//go:generate mockgen -source=albums.go -destination=mocks/albums.go -package mocks

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Repo interface {
	CreateAlbum(album Album) Album
	Albums() []Album
	AlbumByID(id string) (Album, bool)
}

type inMemoryRepo struct {
	albums []Album
}

func NewRepo() Repo {
	// albums slice to seed record album data.
	var albums = []Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	return &inMemoryRepo{
		albums: albums,
	}
}

func (r *inMemoryRepo) CreateAlbum(album Album) Album {
	// Add the new album to the slice.
	r.albums = append(r.albums, album)
	return album
}

func (r *inMemoryRepo) Albums() []Album {
	return r.albums
}

func (r *inMemoryRepo) AlbumByID(id string) (Album, bool) {
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range r.albums {
		if a.ID == id {
			return a, true
		}
	}
	return Album{}, false
}
