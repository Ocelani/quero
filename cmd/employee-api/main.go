package main

import (
	"flag"
	"fmt"
	"os"
	"quero2pay/internal/database"
	"quero2pay/pkg/business"
	"quero2pay/pkg/employee"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"
)

const (
	DefaultPort = 4000 // DefaultPort is the default port of this API.
)

// Accepts is the application header accepted content type.
func Accepts(c *fiber.Ctx) error {
	c.Accepts("application/json", "text/html", "html", "text", "json")
	return nil
}

// Cors stands for Cross Origins Rources Sharing.
func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, X-Auth",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowOrigins:     "*",
		AllowCredentials: true,
	})
}

// main entrypoint.
func main() {
	port := flag.Int("port", DefaultPort, "exposed port of this API")
	flag.Parse()

	connDB := database.NewConnection(
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_TZ"),
		os.Getenv("DATABASE_SSL"),
	)

	svc := &Service{
		Employee: employee.NewDefaultService(employee.NewDefaultRepository(connDB)),
		Business: business.NewDefaultService(business.NewDefaultRepository(connDB)),
	}

	a := API{
		App:     fiber.New(),
		Service: svc,
	}

	defer func() {
		if err := a.App.Shutdown(); err != nil {
			_log.Panic().Err(err).Send()
		}
	}()

	a.App.Use(Cors())
	a.Router()

	if err := a.App.Listen(fmt.Sprintf(":%d", *port)); err != nil {
		log.Panic().Err(err)
	}
}
