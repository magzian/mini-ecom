package controllers

import (
	"backend/models"
	"backend/utils"
	"errors"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

/* var (
	products = map[string]models.Product{}
) */

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if result := utils.DB.Create(product); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"product": product})

}

// Getting a single product
func ViewProduct(c *fiber.Ctx) error {
	var product models.Product

	id := c.Params("id")

	result := utils.DB.Select("id", "name", "description", "price").Where("id = ?", id).First(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Product was not found",
				"message": "Product with the specified ID does not exist",
			})
		}
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"product": product,
	})

}

//Getting all products

func ViewAllProducts(c *fiber.Ctx) error {
	var products []models.Product

	result := utils.DB.Select("id", "name", "description", "price").Find(&products)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to fetch products",
			"message": result.Error.Error(),
		})
	}

	if len(products) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error":   "No products found",
			"message": "The product is curerently empty",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success":  true,
		"count":    len(products),
		"products": products,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	/* updateProduct := new(models.Product) */
	id := c.Params("id")
	var existingProduct models.Product
	result := utils.DB.Where("id = ?", id).First(&existingProduct)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Product not found",
				"message": "Product with the specified ID does not exist",
			})
		}
	}

	var updatedData models.Product
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	if updatedData.Name != "" {
		existingProduct.Name = updatedData.Name
	}
	if updatedData.Description != "" {
		existingProduct.Description = updatedData.Description
	}
	if updatedData.Price > 0 {
		existingProduct.Price = updatedData.Price
	}

	if err := utils.DB.Save(&existingProduct).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update product",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Product updated successfully",
		"product": existingProduct,
	})
}

func DeleteProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	var product models.Product
	result := utils.DB.Where("id = ?", id).First(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Product not found",
				"message": "Product with the specified ID does not exist",
			})
		}
	}

	if err := utils.DB.Delete(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete product",
			"message": err.Error(),
		})
	}

	// Return success response
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Product deleted successfully",
		"deleted_product": fiber.Map{
			"id":   product.ID,
			"name": product.Name,
		},
	})
}
