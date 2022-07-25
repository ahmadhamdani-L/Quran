package route

import (
	"golang-api1/controller"

	"github.com/gin-gonic/gin"
)

func surahRoutes(superRoute *gin.RouterGroup) {
	surahRouter := superRoute.Group("/quran")
	{
		surahRouter.GET("/surah", controller.SurahTampil)

		surahRouter.POST("/surahAdd", controller.SurahAdd)

		surahRouter.PUT("/surahChange/:id", controller.SurahChange)

		surahRouter.DELETE("/surahDelete/:juz", controller.SurahDelete)

		surahRouter.GET("/surahOne/:id", controller.FindByIdsurah)



	}
}
