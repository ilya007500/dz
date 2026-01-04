package main

import (
	"log"
	"project/dz/config"
	"project/dz/internal/pages"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Init()
	dbConf := config.NewDataBaseConfig()
	log.Println(dbConf)

	app := fiber.New()
	pages.NewHandler(app)
	app.Listen(":3000")
}