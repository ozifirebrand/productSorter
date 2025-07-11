package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"productSorter/internals/route"
	"productSorter/internals/service"
)

func main() {

	router := gin.Default()

	sortService := service.NewSortService()

	route.SetUpSortRouter(router, sortService)

	err := router.Run(":4000")
	if err != nil {
		log.Fatalf("Failed to start sorter: %v", err)
		return
	}
}
