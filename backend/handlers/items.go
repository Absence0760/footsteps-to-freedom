package handlers

import (
	"app/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

// GetItems returns all items
func GetItems(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not initialized",
		})
	}

	var items []models.Item
	result := DB.Find(&items)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(items)
}

// GetItem returns a single item by ID
func GetItem(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not initialized",
		})
	}

	id := c.Params("id")
	var item models.Item

	result := DB.First(&item, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "Item not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(item)
}

// CreateItem creates a new item
func CreateItem(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not initialized",
		})
	}

	item := new(models.Item)

	if err := c.BodyParser(item); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	result := DB.Create(&item)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.Status(201).JSON(item)
}

// UpdateItem updates an existing item
func UpdateItem(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not initialized",
		})
	}

	id := c.Params("id")
	var item models.Item

	result := DB.First(&item, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "Item not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	DB.Save(&item)

	return c.JSON(item)
}

// DeleteItem deletes an item
func DeleteItem(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not initialized",
		})
	}

	id := c.Params("id")
	var item models.Item

	result := DB.First(&item, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "Item not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	DB.Delete(&item)

	return c.JSON(fiber.Map{
		"message": "Item deleted successfully",
	})
}
