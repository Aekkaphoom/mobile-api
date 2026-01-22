package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"grouplease.co.th/mobile_api/config"
	"grouplease.co.th/mobile_api/internal/routes"
	database "grouplease.co.th/mobile_api/pkg"
)

func main() {
	// == Load Config ==
	cfg := config.LoadConfig()
	fmt.Printf("Path store photo: %s\n", config.GetLocalPath())

	// == Connect Database ==
	db, err := database.SQLConnection(cfg.DbHost, cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbPort)
	if err != nil {
		log.Fatal("Db Connection error: ", err)
	}
	fmt.Printf("Database is %v", db)

	// == Run echo Application ==
	// limApi := echo.New()
	// limApi.Use(middleware.Recover())
	// limApi.Use(middleware.CORS())
	// // == Register Route ==
	// routes.RouteUser(limApi, db)
	// == Start Server ==
	// limApi.Logger.Fatal(limApi.Start(":" + cfg.AppPort))

	// == Run fiber Application ==
	fiberApp := fiber.New()
	fiberApp.Use(logger.New())
	// fiberApp.Use(middleware.Logger())
	fiberApp.Use(cors.New())

	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("mobile api is running !!!!")
	})

	// == Register LIMapi Route ==
	limApi := fiberApp.Group("/mobile/")
	routes.FiberRouteUser(limApi, db)
	routes.FiberRouteUserLogin(limApi, db)
	routes.FiberRouteAuthToken(limApi, db)
	// == FollowUp Customer Api ==
	routes.FiberRouteCustomers(limApi, db)
	// == Customer Photo ==
	routes.FiberRouteCustomersPhoto(limApi, db, config.GetLocalPath())

	// == Start Server ==
	log.Fatal(fiberApp.Listen(":" + cfg.AppPort))
}
