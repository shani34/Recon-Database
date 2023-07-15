package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shani34/Recon-Database/connection"
	"github.com/shani34/Recon-Database/server"
)
func main()  {
     app:=fiber.New()

	 db:=connection.DbConnection()
	 
	 log.Print("successfully connected to db")
	 db.AutoMigrate(&server.IP{})
	 //ip routes
	 app.Get("/api/v1/ip/:id", server.Get)
	 app.Post("/api/v1/ip", server.Post)
	 app.Listen(":8080")
}