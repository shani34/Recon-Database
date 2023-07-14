package connection

import (
	"log"
	"os"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
	"gopkg.in/yaml.v2"
)

type Config struct{
	Database struct {
	Host     string `yaml:"host" envconfig:"DATABASE_HOST"`
	User     string `yaml:"user" envconfig:"DATABASE_USER"`
	Password string `yaml:"password" envconfig:"DATABASE_PASSWORD"`
	DBName   string `yaml:"dbname" envconfig:"DATABASE_NAME"`
	Port     string `yaml:"port" envconfig:"DATABASE_PORT"`
	SSLmode  string `yaml:"sslmode" envconfig:"SSLMODE"`
} `yaml:"database"`
}
var config *Config
func DbConnection()*gorm.DB{
	f, err := os.Open("")
	if err != nil {
		log.Fatal("Error opening config file, does it exist?:", err)
	}
	defer f.Close()

	// parse the config file
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error decoding config.yml", err)

	}

	dsn := "host=" + config.Database.Host + " user=" + config.Database.User + " password=" + config.Database.Password + " dbname=" + config.Database.DBName + " port=" + config.Database.Port + " sslmode=" + config.Database.SSLmode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return db
}