package route

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(superRoute *gin.RouterGroup) {
	juzRoutes(superRoute)
	surahRoutes(superRoute)
}
