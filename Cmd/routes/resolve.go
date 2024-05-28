package routes

import (
	"Go-URL-shortener-Redis/api/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(ctx *fiber.Ctx) error {

	url := ctx.Params("url")

	client := database.CreateClient(0)
	defer client.Close()

	value, err := client.Get(database.Context, url).Result()
	if err != redis.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": " short not found on database",
		})
	} else if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot connect DB",
		})
	}

	createClient := database.CreateClient(1)

	defer createClient.Close()
	_ = createClient.Incr(database.Context, "counter")
	return ctx.Redirect(value, 301)

}
