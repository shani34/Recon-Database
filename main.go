package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shani34/Recon-Database/server"
)
func main()  {
     app:=fiber.New()

	 //ips routes
	 app.Get("/api/ips", server.Get)

	 app.Listen(":8080")
}