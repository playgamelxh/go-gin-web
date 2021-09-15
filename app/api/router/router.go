package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	r := e.Group("/api")
	{
		r.GET("/index", indexHandler)
		r.GET("/list", listHandler)
	}
}

func indexHandler(c *gin.Context)  {
	fmt.Println("Api index")
}

func listHandler(c *gin.Context)  {
	fmt.Println("Api list")
}