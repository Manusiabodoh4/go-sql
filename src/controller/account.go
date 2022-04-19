package controller

import (
	"net/http"

	"github.com/Manusiabodoh4/go-sql/src/entity"
	"github.com/Manusiabodoh4/go-sql/src/repository"
	"github.com/Manusiabodoh4/go-sql/src/repository/connection"
	"github.com/Manusiabodoh4/go-sql/src/tools"
	"github.com/labstack/echo/v4"
)

type AccountController interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	GetAll(c echo.Context) error
	GetByEmail(c echo.Context) error
}

type AccountControllerImpl struct {
	Logger     tools.ToolsLogger
	Response   tools.ToolsResponse
	Repository repository.Repository
}

func NewAccountController() AccountController {
	return &AccountControllerImpl{
		Logger:     tools.NewToolsLogger(),
		Response:   tools.NewToolsReponse(),
		Repository: repository.NewAccountRepo(connection.GetConnectionPostgres()),
	}
}

func (st *AccountControllerImpl) Login(c echo.Context) error {
	defer st.Logger.LoggerError(c)
	st.Repository = repository.NewAccountRepo(connection.GetConnectionPostgres())
	body := entity.RequestAccountLogin{}
	err := c.Bind(&body)
	if err != nil {
		panic(err)
	}
	channel := make(chan entity.TemplateChannelResponse)
	defer close(channel)
	go st.Repository.FindWithParam(c.Request().Context(), channel, "Email = $1 AND Password = $2", body.Email, body.Password)
	data := <-channel
	if data.Error != nil {
		panic(data.Error)
	}
	if data.Data == nil {
		return st.Response.SenderResponseJSON(c, http.StatusNotFound, "Data not found", nil)
	}
	return st.Response.SenderResponseJSON(c, http.StatusOK, "Berhasil Login", data.Data)
}

func (st *AccountControllerImpl) Register(c echo.Context) error {
	defer st.Logger.LoggerError(c)
	st.Repository = repository.NewAccountRepo(connection.GetConnectionPostgres())
	body := entity.RequestAccountRegister{}
	err := c.Bind(&body)
	if err != nil {
		panic(err)
	}
	channel := make(chan entity.TemplateChannelResponse)
	defer close(channel)
	go st.Repository.FindWithParam(c.Request().Context(), channel, "Email = $1", body.Email)
	data := <-channel
	if data.Error != nil {
		panic(data.Error)
	}
	if data.Data != nil {
		return st.Response.SenderResponseJSON(c, http.StatusAlreadyReported, "Email Already Registered", data.Data)
	}
	go st.Repository.InsertOne(c.Request().Context(), channel, body.Nama, body.Email, body.Password, body.Age)
	data = <-channel
	if data.Error != nil {
		panic(data.Error)
	}
	return st.Response.SenderResponseJSON(c, http.StatusOK, "Register Success", body)
}

func (st *AccountControllerImpl) GetAll(c echo.Context) error {
	defer st.Logger.LoggerError(c)
	st.Repository = repository.NewAccountRepo(connection.GetConnectionPostgres())
	channel := make(chan entity.TemplateChannelResponse)
	defer close(channel)
	go st.Repository.Find(c.Request().Context(), channel)
	data := <-channel
	if data.Error != nil {
		panic(data.Error)
	}
	if data.Data == nil {
		return st.Response.SenderResponseJSON(c, http.StatusOK, "Data Account Empty", nil)
	}
	return st.Response.SenderResponseJSON(c, http.StatusOK, "Berhasil mengembalikan data Account", data.Data)
}

func (st *AccountControllerImpl) GetByEmail(c echo.Context) error {
	defer st.Logger.LoggerError(c)
	st.Repository = repository.NewAccountRepo(connection.GetConnectionPostgres())
	channel := make(chan entity.TemplateChannelResponse)
	defer close(channel)
	go st.Repository.FindWithParam(c.Request().Context(), channel, "Email = $1", c.Param("email"))
	data := <-channel
	if data.Error != nil {
		panic(data.Error)
	}
	if data.Data == nil {
		return st.Response.SenderResponseJSON(c, http.StatusNotFound, "Data not found", nil)
	}
	return st.Response.SenderResponseJSON(c, http.StatusOK, "Email Berhasil ditemukan", data.Data)
}
