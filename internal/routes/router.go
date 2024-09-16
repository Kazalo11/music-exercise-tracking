package routes

import (
	session "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	middleware "music-exercise-tracking/middleware"

	"github.com/gin-gonic/gin"
)

func Start() {

	router := gin.Default()
	store := cookie.NewStore()
	router.Use(session.Sessions("mysession", store))
	router.Use(middleware.CORSMiddleware())
	v1 := router.Group("/v1")
	AddRoutes(v1)
	router.Run()
}

func AddRoutes(superRoute *gin.RouterGroup) {
	SpotifyRoutes(superRoute)
	StravaRoutes(superRoute)
	LastFMRoutes(superRoute)
}
