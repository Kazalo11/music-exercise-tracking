package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start() {

	router := gin.Default()
	v1 := router.Group("/v1")
	AddRoutes(v1)
	fmt.Println("Server is running at http://localhost:8080")
	router.Run()
}

func AddRoutes(superRoute *gin.RouterGroup) {
	SpotifyRoutes(superRoute)
	StravaRoutes(superRoute)
}
