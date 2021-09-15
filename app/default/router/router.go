package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-web/app/default/controller"
	_ "go-gin-web/app/default/controller"
)

func Routers(e *gin.Engine) {

	// index
	e.GET("/", indexIndexHandler)
	e.GET("/index", indexIndexHandler)
	e.GET("/list", indexListHandler)

	// login
	e.GET("/login/index", loginIndexHandler)
	e.GET("/login/list", loginListHandler)

}

func indexIndexHandler(c *gin.Context)  {
	//fmt.Println("Default index")
	controller.Index.IndexAction()
}

func indexListHandler(c *gin.Context)  {
	//fmt.Println("Default list")
	controller.Index.ListAction()
}

func loginIndexHandler(c *gin.Context)  {
	//fmt.Println("Default index")
	controller.Login.IndexAction()
}

func loginListHandler(c *gin.Context)  {
	//fmt.Println("Default list")
	controller.Login.ListAction()
}