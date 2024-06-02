package dashboard

import (
	"Project3/app/configs"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var BookCollection *mongo.Collection = configs.GetCollection(configs.DB, "books")

// bookCount
func BookCount(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{"$group", bson.D{{"_id", nil}, {"count", bson.D{{"$sum", 1}}}}}},
	}
	cursor, err := BookCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	count := 0
	if len(result) > 0 {
		count = int(result[0]["count"].(int32))
	}

	return c.JSON(fiber.Map{"count": count})
}
