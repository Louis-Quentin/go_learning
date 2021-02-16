package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func World(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong"})
}

func CookieRequest(c *gin.Context) {
	cookie, _ := c.Request.Cookie("message")
	c.JSON(http.StatusOK, gin.H {
		"message": cookie.String()	})

}

func BodyRequest(c *gin.Context) {
	body, _ := c.Request.Cookie("message")
	c.JSON(http.StatusOK, gin.H {
		"message": body.String()	})
}

func HeaderRequest(c *gin.Context) {
	header, _ := c.Request.Cookie("message")
	c.JSON(http.StatusOK, gin.H {
		"message": header.String()	})
}

func QueryRequest(c *gin.Context) {
	query, _ := c.Request.Cookie("message")
	c.JSON(http.StatusOK, gin.H {
		"message": query.String()})
}

func Param(c *gin.Context) {
	param, _ := c.Request.Cookie("message")
	c.JSON(http.StatusOK, gin.H {
		"message": param.String()	})
}