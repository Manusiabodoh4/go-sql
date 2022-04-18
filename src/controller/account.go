package controller

import (
	"net/http"
	"sync"

	"github.com/Manusiabodoh4/go-sql/src/entity"
	"github.com/Manusiabodoh4/go-sql/src/repository"
	"github.com/Manusiabodoh4/go-sql/src/tools"
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

	defer tools.Recover(c)

	body := entity.RequestAccountLogin{}

	err := c.Bind(body)

	if err != nil {
		tools.SenderResponseJSON(c, http.StatusInternalServerError, "Server Error", nil)
		panic(err)
	}

	group := &sync.WaitGroup{}

	channel := make(chan entity.TemplateChannelResponse)

	defer close(channel)

	db := repository.GetConnectionPostgres()

	defer db.Close()

	repoAccount := repository.NewAccountRepo(db)

	go repoAccount.FindByEmail(c.Request().Context(), body.Email, group, channel)

	data := <-channel

	if data.Error != nil {
		tools.SenderResponseJSON(c, http.StatusInternalServerError, "Server Error", nil)
		panic(data.Error)
	}

	group.Wait()

	if data.Data == nil {
		return tools.SenderResponseJSON(c, http.StatusNotFound, "Email Belum Terdaftar", nil)
	}

	return tools.SenderResponseJSON(c, http.StatusOK, "Berhasil Login", data.Data)
}

func (st *AccountControllerImpl) Register(c echo.Context) error {
	return nil
}

func (st *AccountControllerImpl) GetAll(c echo.Context) error {

	defer tools.Recover(c)

	group := &sync.WaitGroup{}

	channel := make(chan entity.TemplateChannelResponse)

	defer close(channel)

	db := repository.GetConnectionPostgres()

	defer db.Close()

	repoAccount := repository.NewAccountRepo(db)

	go repoAccount.FindAll(c.Request().Context(), group, channel)

	data := <-channel

	group.Wait()

	if data.Error != nil {
		tools.SenderResponseJSON(c, http.StatusInternalServerError, "Server Error", nil)
		panic(data.Error)
	}

	if data.Data == nil {
		return tools.SenderResponseJSON(c, http.StatusOK, "Data Account Empty", nil)
	}

	return tools.SenderResponseJSON(c, http.StatusOK, "Berhasil mengembalikan data Account", data.Data)

}

func (st *AccountControllerImpl) GetByEmail(c echo.Context) error {

	defer tools.Recover(c)

	group := &sync.WaitGroup{}

	channel := make(chan entity.TemplateChannelResponse)

	defer close(channel)

	db := repository.GetConnectionPostgres()

	defer db.Close()

	repoAccount := repository.NewAccountRepo(db)

	go repoAccount.FindByEmail(c.Request().Context(), c.Param("email"), group, channel)

	data := <-channel

	group.Wait()

	if data.Error != nil {
		tools.SenderResponseJSON(c, http.StatusInternalServerError, "Server Error", nil)
		panic(data.Error)
	}

	if data.Data == nil {
		return tools.SenderResponseJSON(c, http.StatusNotFound, "Email not found", nil)
	}

	return tools.SenderResponseJSON(c, http.StatusOK, "Berhasil mengembalikan data email", data.Data)

}
