package handler

import (
	"JSONAPI/app/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sort"
)

func GetAllProducts(c *fiber.Ctx) error {
	sortedProducts := sortProductsByID(model.Products)
	if len(sortedProducts) == 0 {
		fmt.Println("No products found")
	} else {
		fmt.Println("Products found:", sortedProducts)
	}
	return c.JSON(sortedProducts)
}

func sortProductsByID(products []model.Product) []model.Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Id < products[j].Id
	})
	return products
}
