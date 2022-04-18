package tools

import "github.com/labstack/echo/v4"

func Recover(c echo.Context) {
	message := recover()
	if message != nil {
		c.Logger().Error(message)
	}
}
