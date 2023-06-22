package main

import (
	"encoding/json"
	"github.com/barlibae/webservice/albums"
	"github.com/barlibae/webservice/albums/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Albums web service tests", func() {
	var mockCtrl *gomock.Controller
	var repo *mocks.MockRepo
	var router *gin.Engine
	var album = albums.Album{
		ID:     "100",
		Title:  "Some title",
		Artist: "Some artist",
		Price:  100,
	}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repo = mocks.NewMockRepo(mockCtrl)
		router = newRouter(repo)
	})

	Context("When retrieving an album by id", func() {

		It("Returns the album when it exists", func() {
			w := httptest.NewRecorder()

			repo.EXPECT().AlbumByID(album.ID).Return(album, true)

			req, _ := http.NewRequest("GET", "/albums/"+album.ID, nil)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(200))
			var actual albums.Album
			err := json.Unmarshal([]byte(w.Body.String()), &actual)
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(Equal(album))
		})

		It("Returns not found error", func() {
			w := httptest.NewRecorder()

			id := "non-existent-album"
			repo.EXPECT().AlbumByID(id).Return(albums.Album{}, false)

			req, _ := http.NewRequest("GET", "/albums/"+id, nil)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(404))
			Expect(w.Body.String()).To(ContainSubstring("album not found"))
		})
	})
})
