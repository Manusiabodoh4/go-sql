package main

import (
	routes "github.com/Manusiabodoh4/go-sql/src/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	routes.NewAccountRoutes(router.Group("/v1/account"))

	router.Run(":5678")

}
