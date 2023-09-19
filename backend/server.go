package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// albums
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album slice
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Mingus Ah Um", Artist: "Charles Mingus", Price: 25.99},
	{ID: "5", Title: "The Clown", Artist: "Charles Mingus", Price: 29.99},
	{ID: "6", Title: "The Bridge", Artist: "Sonny Rollins", Price: 19.99},
	{ID: "7", Title: "Somethin' Else", Artist: "Cannonball Adderley", Price: 15.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func Server() {

	// start Gin server
	router := gin.Default()
	router.GET("/albums", getAlbums)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5173"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
	})
	handler := c.Handler(router)

	http.ListenAndServe("localhost:9000", handler)

}
