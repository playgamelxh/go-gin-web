package controller

import "fmt"

type LoginController struct {
	Module string
	Controller string
}

var Login LoginController

func init() {
	Index.Module = "default"
	Index.Controller = "index"
}


func (login *LoginController)IndexAction() {
	fmt.Println("I'm login controller index action")

}

func (login *LoginController)ListAction() {
	fmt.Println("I'm login controller list action")
}