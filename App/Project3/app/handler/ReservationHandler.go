package handler

import (
	"Project3/app/configs"
	"Project3/app/model"
	"Project3/app/responses"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var reservationCollection *mongo.Collection = configs.GetCollection(configs.DB, "Reservations")

//var bookCollection *mongo.Collection = configs.GetCollection(configs.DB, "Books")

func GetAReservation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Id := c.Params("_id")
	objId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ReservationResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": "Invalid reservation ID"},
		})
	}
	var reservation model.Reservation
	if err := reservationCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&reservation); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ReservationResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	return c.Status(http.StatusOK).JSON(responses.ReservationResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": reservation},
	})
}

func GetAllReservations(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var reservations []model.Reservation
	results, err := reservationCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ReservationResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleReservation model.Reservation
		if err := results.Decode(&singleReservation); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ReservationResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
		}
		reservations = append(reservations, singleReservation)
	}

	return c.Status(http.StatusOK).JSON(responses.ReservationResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": reservations},
	})
}

func EditAReservation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Id := c.Params("_id")
	var Reservation model.Reservation
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(Id)

	if err := c.BodyParser(&Reservation); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ReservationResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&Reservation); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ReservationResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{
		"book_id":              Reservation.BookId,
		"user_id":              Reservation.UserId,
		"is_completed":         Reservation.IsCompleted,
		"end_time":             Reservation.EndTime,
		"start_time":           Reservation.StartTime,
		"reservation_duration": Reservation.ReservationDuration,
		"status":               Reservation.Status,
	}

	result, err := reservationCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ReservationResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	var updatedReservation model.Book
	if result.MatchedCount == 1 {
		err := reservationCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedReservation)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ReservationResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.ReservationResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedReservation}})
}

func DeleteAReservation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(c.Params("_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid reservation ID"})
	}

	result, err := reservationCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BookResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(responses.BookResponse{
			Status:  http.StatusNotFound,
			Message: "error",
			Data:    &fiber.Map{"data": "Reservation with specified ID not found!"},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.BookResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": "Reservation successfully deleted!"},
	})
}

func CreateAReservation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var reservation model.Reservation
	defer cancel()

	if err := c.BodyParser(&reservation); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ReservationResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&reservation); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ReservationResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newReservation := model.Reservation{
		Id:                  primitive.NewObjectID(),
		BookId:              reservation.BookId,
		UserId:              reservation.UserId,
		IsCompleted:         reservation.IsCompleted,
		EndTime:             reservation.EndTime,
		StartTime:           reservation.StartTime,
		ReservationDuration: reservation.ReservationDuration,
		Status:              reservation.Status,
	}

	result, err := reservationCollection.InsertOne(ctx, newReservation)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ReservationResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.ReservationResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}
