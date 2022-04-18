package tools

import (
	"github.com/Manusiabodoh4/go-sql/src/entity"
	"github.com/labstack/echo/v4"
)

func SenderResponseJSON(c echo.Context, code int, msg string, data interface{}) error {
	res := entity.TemplateResponse{}
	res.Status = code
	res.Message = msg
	res.Data = data
	return c.JSON(code, res)
}
