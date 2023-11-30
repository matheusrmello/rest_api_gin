package database

import (
	"log"
	"os"

	"github.com/matheusrmello/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() {
	stringConexao := "host="+os.Getenv("HOST")+" user="+os.Getenv("USER")+" password="+os.Getenv("PASSWORD")+" dbname="+os.Getenv("DBNAME")+" port="+os.Getenv("DBPORT")+" sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("Erro ao se conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.AlunosDados{})
}
