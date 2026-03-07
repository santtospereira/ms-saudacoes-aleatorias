package database

import (
	"log"
	"os"

	"github.com/avanti-dvp/ms-saudacoes-aleatorias/models"

	"gorm.io/driver/postgres" // Trocamos sqlite por postgres
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Pega a URL que configuramos no painel do Render
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("A variável de ambiente DATABASE_URL não foi definida!")
	}

	// Conecta ao Postgres do Render
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados PostgreSQL!")
	}

	err = db.AutoMigrate(&models.Greeting{})
	if err != nil {
		log.Fatal("Falha ao migrar o banco de dados!")
	}

	DB = db
	SeedDatabase()
}

func SeedDatabase() {
	var count int64
	DB.Model(&models.Greeting{}).Count(&count)

	if count == 0 {
		log.Println("Banco de dados vazio. Inserindo saudações iniciais...")
		greetings := []models.Greeting{
			{Text: "Olá"},
			{Text: "Bem-vindo"},
			{Text: "Que a Força esteja com você"},
			{Text: "E aí, tudo certo"},
			{Text: "Live long and prosper"},
			{Text: "Opa, bão"},
			{Text: "Saudações"},
			{Text: "Keep calm and code on"},
			{Text: "Alô, alô"},
		}

		if err := DB.Create(&greetings).Error; err != nil {
			log.Fatalf("Falha ao inserir carga inicial: %v", err)
		}
	}
}