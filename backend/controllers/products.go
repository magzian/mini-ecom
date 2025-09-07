package controllers

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var (
	products = map[string]models.Product{}
)

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	err := c.BodyParser(&product)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	product.ID = uuid.New().String()

	products[product.ID] = *product

	c.Status(200).JSON(&fiber.Map{
		"product": product,
	})

	return nil
}

// Getting a single product
func ViewProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	if product, ok := products[id]; ok {
		c.Status(200).JSON(&fiber.Map{
			"product": product,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "product not found",
		})
	}
	return nil
}

//Getting all products

func ViewAllProducts(c *fiber.Ctx) error {
	c.Status(200).JSON(&fiber.Map{
		"products": products,
	})
	return nil
}

func UpdateProduct(c *fiber.Ctx) error {
	updateProduct := new(models.Product)

	err := c.BodyParser(updateProduct)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	id := c.Params("id")
	if product, ok := products[id]; ok {
		product.Name = updateProduct.Name
		product.Price = updateProduct.Price
		product.Description = updateProduct.Description
		products[id] = product
		c.Status(200).JSON(&fiber.Map{
			"product": product,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "product not found",
		})
	}
	return nil
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, ok := products[id]; ok {
		delete(products, id)
		c.Status(200).JSON(&fiber.Map{
			"message": "article deleted successfully",
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "article not found",
		})
	}
	return nil
}
