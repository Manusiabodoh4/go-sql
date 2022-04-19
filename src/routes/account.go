package routes

import (
	"github.com/Manusiabodoh4/go-sql/src/controller"
	"github.com/labstack/echo/v4"
)

type RoutesAccount struct {
	Router *echo.Group
}

func NewRoutesAccount(router *echo.Group) Routes {
	return &RoutesAccount{Router: router}
}

func (st *RoutesAccount) NewCreateRoutes() {

	accountController := controller.NewAccountController()

	st.Router.GET("/all", accountController.GetAll)
	st.Router.GET("/filter/email/:email", accountController.GetByEmail)

	st.Router.POST("/login", accountController.Login)
	st.Router.POST("/register", accountController.Register)

}
