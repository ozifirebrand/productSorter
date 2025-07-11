package route

import (
	"github.com/gin-gonic/gin"
	"productSorter/internals/handler"
	"productSorter/internals/service"
)

func SetUpSortRouter(router *gin.Engine, sortService service.SorterService) {
	sorterHandler := handler.NewSorterHandler(sortService)

	routes := router.Group("/sorter")
	{
		routes.POST("/", sorterHandler.SortProducts)
	}

}
