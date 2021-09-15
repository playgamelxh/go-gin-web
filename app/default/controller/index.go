package controller

import "fmt"

type IndexController struct {
	Module string
	Controller string
}

var Index IndexController

func init() {
	Index.Module = "default"
	Index.Controller = "index"
}


func (index *IndexController)IndexAction() {
	fmt.Println("I'm index controller index action")
}

func (index *IndexController)ListAction() {
	fmt.Println("I'm index controller list action")
}