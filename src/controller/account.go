package controller

import (
	"net/http"

	"github.com/Manusiabodoh4/go-sql/src/entity"
	"github.com/Manusiabodoh4/go-sql/src/repository"
	"github.com/labstack/echo/v4"
)

type AccountController interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	GetAll(c echo.Context) error
	GetByEmail(c echo.Context) error
}

type AccountControllerImpl struct{}

func NewAccountController() AccountController {
	return &AccountControllerImpl{}
}

func (st *AccountControllerImpl) Login(c echo.Context) error {
	return nil
}

func (st *AccountControllerImpl) Register(c echo.Context) error {
	return nil
}

func (st *AccountControllerImpl) GetAll(c echo.Context) error {

	db := repository.GetConnectionPostgres()

	defer db.Close()

	repoAccount := repository.NewAccountRepo(db)

	res := entity.TemplateResponse{
		Status:  http.StatusOK,
		Message: "",
		Data:    nil,
	}

	data, error := repoAccount.FindAll(c.Request().Context())

	if error != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Terjadi kesalahan ketika menarik data"
		res.Data = error
	}

	res.Status = http.StatusOK
	res.Message = "Berhasil mengembalikan data"
	res.Data = data

	return c.JSON(http.StatusOK, res)

}

func (st *AccountControllerImpl) GetByEmail(c echo.Context) error {
	return nil
}
