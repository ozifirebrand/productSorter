package strategy

import (
	"productSorter/internals/dto"
	"sort"
)

type SortStrategy interface {
	Sort([]dto.Product) []dto.Product
}

type SortByPrice struct {
}

type SortBySalesPerView struct {
}

func (s SortByPrice) Sort(input []dto.Product) []dto.Product {
	products := append([]dto.Product(nil), input...)

	sort.Slice(products, func(firstComparator, secondComparator int) bool {
		return products[firstComparator].Price < products[secondComparator].Price
	})

	return products
}

func (s SortBySalesPerView) Sort(input []dto.Product) []dto.Product {
	products := append([]dto.Product(nil), input...)

	sort.Slice(products, func(firstComparator, secondComparator int) bool {

		firstRatio := float64(products[firstComparator].Sales) / float64(products[firstComparator].Views)
		secondRatio := float64(products[secondComparator].Sales) / float64(products[secondComparator].Views)

		return firstRatio < secondRatio
	})

	return products
}
