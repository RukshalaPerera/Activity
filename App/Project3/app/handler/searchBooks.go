package handler

import (
	"Project3/app/configs"
	"Project3/app/model"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

var BookCollection *mongo.Collection = configs.GetCollection(configs.DB, "books")

func SearchBooks(c *fiber.Ctx) error {
	query := c.Query("query")
	field := c.Query("field")

	filter := bson.M{}

	switch field {
	case "id":
		id, err := primitive.ObjectIDFromHex(query)
		if err == nil {
			filter["_id"] = id
		}
	case "title":
		filter["title"] = bson.M{"$regex": query, "$options": "i"}
	case "author":
		filter["author"] = bson.M{"$regex": query, "$options": "i"}
	case "book_name":
		filter["book_name"] = bson.M{"$regex": query, "$options": "i"}
	case "is_book_available":
		isAvailable, err := strconv.ParseBool(query)
		if err == nil {
			filter["is_book_available"] = isAvailable
		}
	}

	var books []model.Book
	cursor, err := BookCollection.Find(context.TODO(), filter)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var book model.Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(books)
}
