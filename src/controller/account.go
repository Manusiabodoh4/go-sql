package controller

import (
	"github.com/Manusiabodoh4/go-sql/src/repository"
	"github.com/gin-gonic/gin"
)

type AccountController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetAll(c *gin.Context)
	GetByEmail(c *gin.Context)
}

type AccontImplController struct{}

func NewAccountController() AccountController {
	return &AccontImplController{}
}

func (st *AccontImplController) Login(c *gin.Context) {

}

func (st *AccontImplController) Register(c *gin.Context) {

}

func (st *AccontImplController) GetAll(c *gin.Context) {

	db := repository.GetConnectionPostgres()

	repoAccount := repository.NewAccountRepo(db)

	res, error := repoAccount.FindAll(c)

	if error != nil {
		c.JSON(500, error)
	}

	defer db.Close()

	c.JSON(200, res)

}

func (st *AccontImplController) GetByEmail(c *gin.Context) {

}
