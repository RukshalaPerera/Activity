package handler

import (
	"Project3/app/configs"
	model "Project3/app/model"
	"Project3/app/responses"
	"context"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookCollection *mongo.Collection = configs.GetCollection(configs.DB, "books")

var validate = validator.New()

func CreateBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var book model.Book
	defer cancel()

	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&book); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newBook := model.Book{
		ID:              primitive.NewObjectID(),
		Title:           book.Title,
		Author:          book.Author,
		BookName:        book.BookName,
		IsBookAvailable: book.IsBookAvailable,
		RegisterDate:    book.RegisterDate,
	}

	result, err := bookCollection.InsertOne(ctx, newBook)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.BookResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

// getUser by id
func GetABook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Id := c.Params("_id")
	var book model.Book
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(Id)

	err := bookCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&book)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.BookResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": book}})
}

//updateUser

func EditABook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Id := c.Params("_id")
	var book model.Book
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(Id)

	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&book); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{
		"title":           book.Title,
		"author":          book.Author,
		"book_name":       book.BookName,
		"isBookAvailable": book.IsBookAvailable,
		"registerDate":    book.RegisterDate,
	}

	result, err := bookCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	var updatedBook model.Book
	if result.MatchedCount == 1 {
		err := bookCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedBook)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.BookResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedBook}})
}

//deleteUser

func DeleteABook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Id := c.Params("_id")

	objId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": "Invalid book ID format"}})
	}

	result, err := bookCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.BookResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Book with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.BookResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Book successfully deleted!"}},
	)
}

func GetAllBooks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var books []model.Book
	results, err := bookCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BookResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleBook model.Book
		if err = results.Decode(&singleBook); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.BookResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
		}
		books = append(books, singleBook)
	}

	return c.Status(http.StatusOK).JSON(responses.BookResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": books},
	})
}

// this is book-list
func GetBookList(c *fiber.Ctx) error {
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex", "0"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))
	sortActive := c.Query("sortActive", "title")
	sortDirection := c.Query("sortDirection", "asc")

	sortOrder := 1
	if sortDirection == "desc" {
		sortOrder = -1
	}

	options := options.Find()
	options.SetSkip(int64(pageIndex * pageSize))
	options.SetLimit(int64(pageSize))
	options.SetSort(bson.D{{sortActive, sortOrder}})

	cursor, err := bookCollection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer cursor.Close(context.TODO())

	var books []model.Book
	if err = cursor.All(context.TODO(), &books); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{})
}
