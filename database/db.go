package database

import (
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=" + os.Getenv("HOST") + " user=" + os.Getenv("USER") + " password=" + os.Getenv("PASSWORD") + " dbname=" + os.Getenv("DBNAME") + " port=" + os.Getenv("DBPORT")
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
