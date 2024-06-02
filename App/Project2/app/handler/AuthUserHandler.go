package handler

import (
	"Project2/app/configs"
	"Project2/app/models"
	"Project2/app/responses"
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var AuthUserCollection *mongo.Collection = configs.GetCollection(configs.DB, "AuthUser")
var validateAuthUser = validator.New()

// CreateAuthUser
func CreateAuthUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var authUser models.AuthUser

	if err := c.BodyParser(&authUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.AuthUserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	if validationErr := validateAuthUser.Struct(&authUser); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.AuthUserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": validationErr.Error()},
		})
	}

	newAuthUser := models.AuthUser{
		ID:       primitive.NewObjectID(),
		Username: authUser.Username,
		Password: authUser.Password,
		Email:    authUser.Email,
		Role:     authUser.Role,
		RoleID:   authUser.RoleID,
	}

	result, err := AuthUserCollection.InsertOne(ctx, newAuthUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.AuthUserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	return c.Status(http.StatusCreated).JSON(responses.AuthUserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &fiber.Map{"data": result},
	})
}

// GetAuthUser
func GetAuthUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("_id")
	var authUser models.AuthUser

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.AuthUserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": "Invalid ID format"},
		})
	}

	err = AuthUserCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&authUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.AuthUserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	return c.Status(http.StatusOK).JSON(responses.AuthUserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": authUser},
	})
}

// EditAuthUser
func EditAuthUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("_id")
	var authUser models.AuthUser

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.AuthUserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": "Invalid ID format"},
		})
	}

	if err := c.BodyParser(&authUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.AuthUserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	if validationErr := validateAuthUser.Struct(&authUser); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.AuthUserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": validationErr.Error()},
		})
	}

	update := bson.M{
		"username": authUser.Username,
		"password": authUser.Password,
		"email":    authUser.Email,
		"role_id":  authUser.RoleID,
		"role":     authUser.Role,
	}
	result, err := AuthUserCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.AuthUserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	if result.MatchedCount == 1 {
		err := AuthUserCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&authUser)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.AuthUserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.AuthUserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": authUser},
	})
}

// DeleteAuthUser
func DeleteAuthUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("_id")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.AuthUserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": "Invalid ID format"},
		})
	}

	result, err := AuthUserCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.AuthUserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(responses.AuthUserResponse{
			Status:  http.StatusNotFound,
			Message: "error",
			Data:    &fiber.Map{"data": "User with specified ID not found!"},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.AuthUserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": "User successfully deleted!"},
	})
}

// GetAllAuthUsers
func GetAllAuthUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var authUsers []models.AuthUser

	results, err := AuthUserCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.AuthUserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleAuthUser models.AuthUser
		if err = results.Decode(&singleAuthUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.AuthUserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
		}

		authUsers = append(authUsers, singleAuthUser)
	}

	return c.Status(http.StatusOK).JSON(responses.AuthUserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": authUsers},
	})
}
