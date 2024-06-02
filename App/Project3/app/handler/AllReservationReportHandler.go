package handler

import (
	"Project3/app/model"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"
)

func GenerateAllReservationReportHandler(c *fiber.Ctx) error {
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	var startTime, endTime time.Time
	var err error
	if startTimeStr != "" {
		startTime, err = time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start_time format"})
		}
	}
	if endTimeStr != "" {
		endTime, err = time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end_time format"})
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if !startTime.IsZero() || !endTime.IsZero() {
		filter["start_time"] = bson.M{}
		if !startTime.IsZero() {
			filter["start_time"].(bson.M)["$gte"] = startTime
		}
		if !endTime.IsZero() {
			filter["start_time"].(bson.M)["$lte"] = endTime
		}
	}

	var reservations []model.Reservation
	cursor, err := reservationCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := cursor.All(ctx, &reservations); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "All Reservation Report\n\n")

	if !startTime.IsZero() || !endTime.IsZero() {
		pdf.Cell(0, 10, "Time Range:\n\n")
		if !startTime.IsZero() {
			pdf.Cell(0, 10, "Start: "+startTime.Format(time.RFC3339)+"\n\n")
		}
		if !endTime.IsZero() {
			pdf.Cell(0, 10, "End: "+endTime.Format(time.RFC3339)+"\n\n")
		}
	}

	generatedAt := time.Now().Format("2006-01-02 15:04:05")
	for _, reservation := range reservations {
		pdf.Cell(0, 10, "Reservation ID: "+reservation.Id.Hex()+"\n")
		pdf.Cell(0, 10, "Book ID: "+reservation.BookId.Hex()+"\n")
		pdf.Cell(0, 10, "User ID: "+reservation.UserId.Hex()+"\n")
		pdf.Cell(0, 10, "Is Complete: "+strconv.FormatBool(reservation.IsCompleted)+"\n")
		pdf.Cell(0, 10, "Reservation Start Time: "+reservation.StartTime.Format(time.RFC3339)+"\n")
		pdf.Cell(0, 10, "Reservation End Time: "+reservation.EndTime.Format(time.RFC3339)+"\n")
		pdf.Cell(0, 10, "Status: "+reservation.Status+"\n\n")
	}
	pdf.Cell(0, 10, "Generated At: "+generatedAt+"\n\n")

	filename := "AllReservationReport.pdf"
	if err := pdf.OutputFileAndClose(filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendFile(filename)
}
