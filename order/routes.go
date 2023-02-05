package order

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-oauth2-cognito-lab/payment"
)

func SetRoutes(app *fiber.App) {
	paymentClient := payment.NewClient()

	app.Post("/orders", func(c *fiber.Ctx) error {
		var order Order

		ctx := c.UserContext()

		if err := c.BodyParser(&order); err != nil {
			return err
		}

		if err := paymentClient.CreatePayment(ctx, payment.Payment{CorrelationID: order.ID}); err != nil {
			return err
		}

		return c.SendStatus(http.StatusCreated)
	})

}
