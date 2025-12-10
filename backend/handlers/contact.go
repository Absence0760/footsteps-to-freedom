package handlers

import (
	"app/models"

	"github.com/gofiber/fiber/v2"
)

// CreateContact handles contact form submissions
func CreateContact(c *fiber.Ctx) error {
	// Prevent caching of contact form submissions
	c.Set("Cache-Control", "no-store, no-cache, must-revalidate, private")
	c.Set("Pragma", "no-cache")
	c.Set("Expires", "0")

	contact := new(models.Contact)

	if err := c.BodyParser(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if contact.Name == "" || contact.Email == "" || contact.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, email, and message are required",
		})
	}

	// Set default status
	if contact.Status == "" {
		contact.Status = "new"
	}

	if DB != nil {
		// Save to database
		if err := DB.Create(&contact).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save contact",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Thank you for contacting us! We'll get back to you soon.",
		"contact": contact,
	})
}

// GetContacts retrieves all contact submissions (admin function)
func GetContacts(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Database not connected",
		})
	}

	var contacts []models.Contact
	if err := DB.Order("created_at desc").Find(&contacts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch contacts",
		})
	}

	return c.JSON(contacts)
}

// GetContact retrieves a single contact by ID
func GetContact(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Database not connected",
		})
	}

	id := c.Params("id")
	var contact models.Contact

	if err := DB.First(&contact, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Contact not found",
		})
	}

	return c.JSON(contact)
}

// UpdateContact updates a contact's status (admin function)
func UpdateContact(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Database not connected",
		})
	}

	id := c.Params("id")
	var contact models.Contact

	if err := DB.First(&contact, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Contact not found",
		})
	}

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := DB.Save(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update contact",
		})
	}

	return c.JSON(contact)
}

// DeleteContact deletes a contact (admin function)
func DeleteContact(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Database not connected",
		})
	}

	id := c.Params("id")
	var contact models.Contact

	if err := DB.First(&contact, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Contact not found",
		})
	}

	if err := DB.Delete(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete contact",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
