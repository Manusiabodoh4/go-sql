package routes

import (
	"github.com/Manusiabodoh4/go-sql/src/controller"
	"github.com/gin-gonic/gin"
)

func NewAccountRoutes(router *gin.RouterGroup) {

	accountController := controller.NewAccountController()

	router.GET("/login", accountController.Login)
	router.GET("/register", accountController.Register)
	router.GET("/all", accountController.GetAll)
	router.GET("/filter/:email", accountController.GetByEmail)

}
