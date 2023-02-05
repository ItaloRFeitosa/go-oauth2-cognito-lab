package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/italorfeitosa/go-oauth2-cognito-lab/order"
	"github.com/italorfeitosa/go-oauth2-cognito-lab/payment"
)

var (
	isPayment bool
	isOrder   bool
)

func init() {
	flag.BoolVar(&isPayment, "payment", false, "run payment")
	flag.BoolVar(&isOrder, "order", false, "run order")
}

func main() {
	flag.Parse()

	startServer()
}

func startServer() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println(err.Error())
			return fiber.DefaultErrorHandler(c, err)
		},
	})

	app.Use(recover.New())

	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"isPayment": isPayment,
			"isOrder":   isOrder,
		})
	})

	if isPayment {
		payment.SetRoutes(app)
	}

	if isOrder {
		order.SetRoutes(app)
	}

	err := app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))

	if err != nil {
		log.Fatal(err)
	}

	GracefulShutdown(func(ctx context.Context) {
		if err := app.Shutdown(); err != nil {
			log.Fatal(err)
		}
	})
}

func GracefulShutdown(shutdownCallback func(context.Context)) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit

	log.Println("gracefully shutdown process...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer signal.Stop(quit)

	shutdownCallback(ctx)

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("process exiting")
}
