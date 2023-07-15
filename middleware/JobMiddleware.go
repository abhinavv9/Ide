package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func JobMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract the necessary data from the request
		// code := c.FormValue("code")
		// userToken := c.FormValue("userToken")
		// language := c.FormValue("language")

		// Get the image name from the language
		// image := ""
		// switch language {
		// case "py":
		// 	image = "chat-py"
		// case "cpp":
		// 	image = "chat-cpp"
		// case "go":
		// 	image = "chat-go"
		// }

		// Create a new Job instance
		// job := Job{
		// 	UserID: userToken,
		// 	Code:   code,
		// 	Lang:   language,
		// 	Image:  image,
		// }

		// Pass the job to the channel for the worker to execute
		// jobCh <- job

		// Proceed to the next middleware/handler
		return c.Next()
	}
}
