package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shani34/Recon-Database/connection"
	"gorm.io/gorm"
)

var db *gorm.DB



type Ips struct{
	gorm.Model
	Id string `json:"id"`
	Program string `json:"Program"`
}

func Get(c *fiber.Ctx)error{
     id:=c.Params("id")

	 var Ip Ips

	 db=connection.DbConnection()
	 db.Find(&Ip, id)
	 err:=c.JSON(Ip)

	return err
}

