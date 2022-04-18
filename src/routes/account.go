package routes

import (
	"github.com/Manusiabodoh4/go-sql/src/controller"
	"github.com/labstack/echo/v4"
)

func NewAccountRoutes(router *echo.Group) {

	accountController := controller.NewAccountController()

	router.GET("/login", accountController.Login)
	router.GET("/register", accountController.Register)
	router.GET("/all", accountController.GetAll)
	router.GET("/filter/email/:email", accountController.GetByEmail)

}
