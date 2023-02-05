package payment

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-oauth2-cognito-lab/auth"
)

func SetRoutes(app *fiber.App) {
	app.Post("/payments", auth.FiberMiddleware(), auth.FiberEnsureScopes("payment/payment.write"), func(c *fiber.Ctx) error {
		var payment Payment
		payment.Status = StatusPending

		if err := c.BodyParser(&payment); err != nil {
			return err
		}

		return c.SendStatus(http.StatusAccepted)

	})

}
