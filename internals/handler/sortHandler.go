package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"productSorter/internals/dto"
	"productSorter/internals/pkg/responses"
	"productSorter/internals/service"
)

type SortHandler struct {
	sortService service.SortService
}

func NewSortHandler(sortService service.SortService) *SortHandler {
	return &SortHandler{sortService: sortService}

}

func (h *SortHandler) SortProducts(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic %v", r)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return

		}
	}()

	var input *dto.SortProductRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response := responses.ErrorResponse("invalid input", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := h.sortService.SortProduct(input)

	if err != nil {
		response := responses.ErrorResponse("invalid input", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := responses.SuccessResponse("Successful", http.StatusOK, data)
	ctx.JSON(http.StatusOK, response)
}
