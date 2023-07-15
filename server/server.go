package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shani34/Recon-Database/connection"
	"gorm.io/gorm"
)

var db *gorm.DB


//models
type IP struct{
	gorm.Model
	Id string `json:"id"`
	Program string `json:"program"`
	Subdomains []*Subdomain `json:"subdomains" gorm:"many2many:subdomain_ips;"`
}

type Platform struct {
	gorm.Model
	ID       string    `json:"id" gorm:"PrimaryKey"`
	URL      string    `json:"url"`
	Programs []Program `json:"programs"`
}

type Program struct {
	gorm.Model
	ID          string       `json:"id" gorm:"PrimaryKey"`
	PlatformID  string       `json:"platform"`
	Subdomains  []Subdomain  `json:"subdomains"`
	RootDomains []RootDomain `json:"rootdomains"`
	IPs         []IP         `json:"ips"`
}

type RootDomain struct {
	gorm.Model
	ID         string      `json:"id" gorm:"PrimaryKey"`
	ProgramID  string      `json:"program"`
	Subdomains []Subdomain `json:"subdomains"`
}

type Subdomain struct {
	gorm.Model
	ID           string `json:"id" gorm:"PrimaryKey"`
	ProgramID    string `json:"program"`
	RootDomainID string `json:"rootdomain"`
	CNAME        string `json:"cname"`
	Nameservers  string `json:"nameservers"`
	IPs          []*IP  `json:"ips" gorm:"many2many:subdomain_ips;"`
}

//get IP
func Get(c *fiber.Ctx)error{
     id:=c.Params("id")

	 var Ip IP

	 db=connection.DbConnection()
	 db.Where("ID=?",id).Find(&Ip)
	 err:=c.JSON(Ip)

	return err
}

//create IP

func Post(c *fiber.Ctx)error{
     var  ips IP

	 db=connection.DbConnection()

	 if err:=c.BodyParser(ips); err!=nil{
		c.Status(503)
	 }

	 db.Create(&ips)
	
	return c.JSON(ips)
}
