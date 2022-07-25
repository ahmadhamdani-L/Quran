package main

import (
	"golang-api1/config"
	"net/http"
	"golang-api1/route"
	"github.com/gin-gonic/gin"
	
)

func main () {
	r := gin.Default () ;

	db := config.ConnectDb()
	r.Use (func (c *gin.Context){
		c.Set ("db", db)
		c.Next ()
	})


    router := r.Group("/api/v1")
    route.AddRoutes(router)


	r.GET("/", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H {"data" : "selamat anda dapat terhubung ke API"})
	})

	r.Run()
}