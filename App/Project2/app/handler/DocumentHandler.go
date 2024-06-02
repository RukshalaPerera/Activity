package handler

import (
	"Project2/app/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"path/filepath"
)

var documents = []models.Document{}

func UploadDocument(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return fiber.ErrBadRequest
	}

	title := c.FormValue("title")
	if title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Title is required",
		})
	}

	id := uuid.New().String()
	filename := fmt.Sprintf("%s%s", id, filepath.Ext(file.Filename))

	err = c.SaveFile(file, fmt.Sprintf("./uploads/%v", filename))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	document := models.Document{
		Filename: filename,
		Title:    title,
	}
	documents = append(documents, document)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"document": document,
	})
}

func ListDocuments(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":   true,
		"documents": documents,
	})
}

func DownloadDocument(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid ID format",
		})
	}
	var document *models.Document
	for _, doc := range documents {
		if doc.ID == id {
			document = &doc
			break
		}
	}
	if document == nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Document not found",
		})
	}
	return c.Download(fmt.Sprintf("./uploads/%s", document.Filename))
}
