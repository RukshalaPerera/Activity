package dashboard

import (
	"Project3/app/configs"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

func GenerateMonthlyBookChart(ctx *fiber.Ctx) error {
	collection := configs.GetCollection(configs.DB, "books")
	mongoCtx := context.Background()

	count := ctx.Query("count")
	monthID := ctx.Query("_id")

	pipeline := bson.A{
		// Convert the register_date string to a date
		bson.D{
			{"$addFields", bson.D{
				{"register_date", bson.D{
					{"$dateFromString", bson.D{
						{"dateString", "$register_date"},
						{"onError", nil},
						{"onNull", nil},
					}},
				}},
			}},
		},
		// Group by year and month
		bson.D{
			{"$group", bson.D{
				{"_id", bson.D{
					{"$dateToString", bson.D{
						{"format", "%Y-%m"},
						{"date", "$register_date"},
					}},
				}},
				{"count", bson.D{
					{"$sum", 1},
				}},
			}},
		},
	}

	if monthID != "" {
		pipeline = append(pipeline, bson.D{
			{"$match", bson.D{
				{"_id", monthID},
			}},
		})
	}

	if count != "" {
		limit, err := strconv.Atoi(count)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Invalid count parameter")
		}
		pipeline = append(pipeline, bson.D{
			{"$limit", limit},
		})
	}

	cursor, err := collection.Aggregate(mongoCtx, pipeline)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error aggregating data: %v", err))
	}
	defer cursor.Close(mongoCtx)

	var results []bson.M
	if err := cursor.All(mongoCtx, &results); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error decoding results: %v", err))
	}
	return ctx.JSON(results)
}
