package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Dados de conexão (exemplo: modifique para o seu banco)
	dsn := "host=localhost user=docker password=docker dbname=fnc port=5435 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	fmt.Println("Banco de dados conectado com sucesso!")
	DB = db
}
