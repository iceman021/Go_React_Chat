package main

import (  
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/pusher/pusher-http-go"
)

func main() {
    app := fiber.New()

	app.Use(cors.New())

    pusherClient := pusher.Client{
        AppID: "1339575",
        Key: "f27c4e6a127e8ee07398",
        Secret: "108bcfc5e2e47011b080",
        Cluster: "eu",
        Secure: true,
    }

    app.Post("api/messages", func(c *fiber.Ctx) error {
        var data map[string]string

        if err := c.BodyParser(&data); err != nil {
            return err
        }

        pusherClient.Trigger("chat", "message", data)
       // return c.SendString("Hello, World and Care Dusane 👋!")
       return c.JSON([]string{})
    })

    app.Listen(":8000")
}