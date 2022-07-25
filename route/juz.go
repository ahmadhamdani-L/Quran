package route

import (
	"golang-api1/controller"

	"github.com/gin-gonic/gin"
)

func juzRoutes(superRoute *gin.RouterGroup) {
	juzRouter := superRoute.Group("/quran")
	{
		juzRouter.GET("/juz", controller.JuzTampil)

		juzRouter.POST("/juzAdd", controller.JuzAdd)

		juzRouter.PUT("/juzChange/:id", controller.JuzChange)

		juzRouter.DELETE("/juzDelete/:juz", controller.JuzDelete)

	}
}
