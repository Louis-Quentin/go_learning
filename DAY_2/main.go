package main

import (
	router "SoftwareGoDay2/routes"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/home", router.World)
	r.GET("/repeat-my-query", router.QueryRequest)
	r.GET("/repeat-my-param:message", router.Param)
	r.POST("/repeat-my-body", router.BodyRequest)
	r.GET("/repeat-my-header", router.HeaderRequest)
	r.GET("/repeat-my-cookie", router.CookieRequest)
}

func main() {
	r := gin.Default()
	ApplyRoutes(r)
	r.Run()

}