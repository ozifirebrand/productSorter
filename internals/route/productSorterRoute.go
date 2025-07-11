package route

import (
	"github.com/gin-gonic/gin"
	"productSorter/internals/handler"
	"productSorter/internals/service"
)

func SetUpSortRouter(router *gin.Engine, sortService service.SortService) {
	sortHandler := handler.NewSortHandler(sortService)

	routes := router.Group("/sorter")
	{
		routes.POST("/price_asc", sortHandler.SortProducts)
	}

}
