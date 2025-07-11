package service

import (
	"fmt"
	"productSorter/internals/dto"
	"productSorter/internals/strategy"
)

type SortService interface {
	SortProduct(input *dto.SortProductRequest) ([]dto.Product, error)
}

func NewSortService() SortService {
	strategyMap := make(map[string]strategy.SortStrategy)
	strategyMap["price"] = strategy.SortByPrice{}
	strategyMap["ratio"] = strategy.SortBySalesPerView{}
	return &sortService{}
}

type sortService struct {
	strategy map[string]strategy.SortStrategy
}

func (s *sortService) SortProduct(input *dto.SortProductRequest) ([]dto.Product, error) {
	sorter, ok := s.strategy[input.Strategy]
	if !ok {
		return nil, fmt.Errorf("invalid sorter")
	}
	return sorter.Sort(input.Products), nil
}
