package service

import (
	"fmt"
	"productSorter/internals/dto"
	"productSorter/internals/pkg/enums"
	"productSorter/internals/strategy"
	"strings"
)

type SorterService interface {
	SortProduct(input *dto.SortProductRequest) ([]dto.Product, error)
}

func NewSortService() SorterService {
	sorterMap := make(map[string]strategy.SortStrategy)
	sorterMap[string(enums.SorterTypePriceAsc)] = strategy.SortByPriceAsc{}
	sorterMap[string(enums.SorterTypeSVRatioAsc)] = strategy.SortBySalesPerViewAsc{}
	sorterMap[string(enums.SorterTypePriceDesc)] = strategy.SortByPriceDesc{}
	sorterMap[string(enums.SorterTypeSVRatioDesc)] = strategy.SortBySalesPerViewDesc{}

	return &sortService{
		sorter: sorterMap,
	}
}

type sortService struct {
	sorter map[string]strategy.SortStrategy
}

func (s *sortService) SortProduct(input *dto.SortProductRequest) ([]dto.Product, error) {
	sortBy := strings.TrimSpace(input.SortBy)

	if !enums.IsValidSorterType(sortBy) {
		return nil, fmt.Errorf("bad request: %v is not a valid sorting value", sortBy)
	}

	if len(input.Products) == 0 {
		return nil, fmt.Errorf("products cannot be empty. Add a product and try again")
	}

	sorter, ok := s.sorter[input.SortBy]
	if !ok {
		return nil, fmt.Errorf("sorter does not exist")
	}

	return sorter.Sort(input.Products), nil
}
