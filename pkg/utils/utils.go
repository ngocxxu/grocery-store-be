package utils

import (
	"regexp"
	"strconv"

	graphModel "github.com/ngocxxu/grocery-store-svelte-be/graph/model"
	internalModel "github.com/ngocxxu/grocery-store-svelte-be/internal/model"
)

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func ConvertToGraphWeightOptions(internalWeightOptions []internalModel.WeightOption) []*graphModel.WeightOption {
	graphWeightOptions := make([]*graphModel.WeightOption, len(internalWeightOptions))
	for i, opt := range internalWeightOptions {
		graphWeightOptions[i] = &graphModel.WeightOption{
			Weight: opt.Weight,
			ID:     strconv.FormatUint(uint64(opt.UnitID), 10),
		}
	}
	return graphWeightOptions
}

func ConvertToGraphProduct(product *internalModel.Product) *graphModel.Product {
	return &graphModel.Product{
		ID:            strconv.FormatUint(uint64(product.ID), 10),
		Name:          product.Name,
		Description:   product.Description,
		Type:          product.Type,
		Sku:           product.Sku,
		Status:        product.Status,
		Price:         product.Price,
		Discount:      product.Discount,
		Rating:        int(product.Rating),
		Quantity:      int(product.Quantity),
		WeightOptions: ConvertToGraphWeightOptions(product.WeightOptions),
	}
}
