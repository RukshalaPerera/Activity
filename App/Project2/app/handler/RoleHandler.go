package handler

import (
	"Project2/app/configs"
	"Project2/app/models"
	"Project2/app/responses"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var roleCollection *mongo.Collection = configs.GetCollection(configs.DB, "roles")
var roleValidate = validator.New()

// create role
func CreateRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var role models.Role
	defer cancel()

	if err := c.BodyParser(&role); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RoleResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&role); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RoleResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newRole := models.Role{
		ID:       primitive.NewObjectID(),
		RoleName: role.RoleName,
	}

	result, err := roleCollection.InsertOne(ctx, newRole)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.RoleResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.RoleResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

// byid
func GetARole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Id := c.Params("_id")
	var role models.Role
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(Id)

	err := roleCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&role)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.RoleResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusOK).JSON(responses.RoleResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": role}})
}

// edit by id
func EditARole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Id := c.Params("_id")
	var role models.Role
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(Id)

	if err := c.BodyParser(&role); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RoleResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&role); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RoleResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	update := bson.M{
		"roleName": role.RoleName,
	}
	result, err := roleCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.RoleResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	var updatedRole models.Role
	if result.MatchedCount == 1 {
		err := roleCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedRole)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.RoleResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.RoleResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedRole}})
}

//deleteUser

func DeleteARole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Id := c.Params("_id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(Id)

	result, err := roleCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.RoleResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.RoleResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.RoleResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
	)
}

// get All users
func GetAllRoles(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var roles []models.Role
	defer cancel()

	results, err := roleCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.RoleResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleRole models.Role
		if err = results.Decode(&singleRole); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.RoleResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		roles = append(roles, singleRole)
	}

	return c.Status(http.StatusOK).JSON(
		responses.RoleResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": roles}},
	)
}
