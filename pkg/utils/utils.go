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
			Unit: &graphModel.Unit{
				ID:           strconv.FormatUint(uint64(opt.Unit.ID), 10),
				Name:         opt.Unit.Name,
				Abbreviation: opt.Unit.Abbreviation,
			},
		}
	}
	return graphWeightOptions
}

func ConvertToGraphCategories(internalCategories []internalModel.Category) []*graphModel.Category {
	graphCategories := make([]*graphModel.Category, len(internalCategories))
	for i, opt := range internalCategories {
		graphCategories[i] = &graphModel.Category{
			ID:          strconv.FormatUint(uint64(opt.ID), 10),
			Name:        opt.Name,
			Description: opt.Description,
			Products:    ConvertToGraphProducts(opt.Products),
		}
	}
	return graphCategories
}

func ConvertToGraphProducts(internalProducts []internalModel.Product) []*graphModel.Product {
	graphProducts := make([]*graphModel.Product, len(internalProducts))
	for i, product := range internalProducts {
		graphProducts[i] = ConvertToGraphProduct(&product)
	}
	return graphProducts
}

func ConvertToGraphProduct(product *internalModel.Product) *graphModel.Product {
	return &graphModel.Product{
		ID:            strconv.FormatUint(uint64(product.ID), 10),
		Name:          product.Name,
		Description:   product.Description,
		Sku:           product.Sku,
		Status:        product.Status,
		Price:         product.Price,
		Discount:      product.Discount,
		Image:         product.Image,
		Rating:        int(product.Rating),
		Quantity:      int(product.Quantity),
		WeightOptions: ConvertToGraphWeightOptions(product.WeightOptions),
	}
}
