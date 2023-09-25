package backend

import (
	"log"
	"main/backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func Server() {

	// start Gin server
	router := gin.Default()
	routes.AlbumRoute(router)

	// start CORS server
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5173, http://localhost:5173"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
	})
	handler := c.Handler(router)

	err := http.ListenAndServe("localhost:9000", handler)
	if err != nil {
		log.Fatal(err)
		return
	}
}
