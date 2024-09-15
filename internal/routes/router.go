package routes

import (
	"fmt"
	"net/http"

	session "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	clientManager "music-exercise-tracking/internal/client"
	middleware "music-exercise-tracking/middleware"

	"github.com/gin-gonic/gin"
)

func Start() {

	router := gin.Default()
	store := cookie.NewStore()
	router.Use(session.Sessions("mysession", store))
	router.Use(middleware.CORSMiddleware())
	v1 := router.Group("/v1")
	v1.GET("/auth", checkAuth)
	AddRoutes(v1)
	fmt.Println("Server is running at http://localhost:8080")
	router.Run()
}

func AddRoutes(superRoute *gin.RouterGroup) {
	SpotifyRoutes(superRoute)
	StravaRoutes(superRoute)
	LastFMRoutes(superRoute)
}

func checkAuth(c *gin.Context) {
	authenticated := clientManager.GetAccessToken() != ""
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"isAuthenticated": authenticated})
}
