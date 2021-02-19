package main

import (
	router "SoftwareGoDay2/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/home", router.World)
	r.GET("/repeat-my-query", router.QueryRequest)
	r.GET("/repeat-my-param:message", router.Param)
	r.POST("/repeat-my-body", router.BodyRequest)
	r.GET("/repeat-my-header", router.HeaderRequest)
	r.GET("/repeat-my-cookie", router.CookieRequest)
	r.GET("/hello", router.HelloWorld)
	r.GET("/repeat-all-my-queries", router.GetAllQueries)
}

func main() {
	r := gin.Default()
	godotenv.Load()
	ApplyRoutes(r)
	r.Run()
}