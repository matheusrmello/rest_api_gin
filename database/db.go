package database

import (
	"log"

	"github.com/matheusrmello/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() {
	stringConexao := "host=localhost user=root password=root dbname=postgres port=5431 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("Erro ao se conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.AlunosDados{})
}
