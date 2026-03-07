package main

import (
	"log"
	"os" // Essencial para ler variáveis de ambiente como PORT e DATABASE_URL

	"github.com/avanti-dvp/ms-saudacoes-aleatorias/database"
	"github.com/avanti-dvp/ms-saudacoes-aleatorias/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa a conexão com o banco de dados. 
	// Lembre-se: sua database.ConnectDatabase() agora deve usar postgres.Open(os.Getenv("DATABASE_URL"))
	database.ConnectDatabase()

	// Cria um router Gin com as configurações padrão
	router := gin.Default()

	// Configuração de CORS simplificada para permitir integração com o frontend
	router.Use(cors.Default())

	// Define as rotas da API
	api := router.Group("/api")
	{
		// Rota para cadastrar um novo cumprimento
		api.POST("/saudacoes", handlers.CreateGreeting)

		// Rota para obter um cumprimento aleatório
		api.GET("/saudacoes/aleatorio", handlers.GetRandomGreeting)
	}

	// Lógica para detectar a porta do ambiente (Render usa 8080 ou 10000)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor padrão para rodar localmente
	}

	log.Printf("Iniciando servidor na porta %s...", port)

	// Inicia o servidor na porta correta injetada pelo Render
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Falha ao rodar o servidor: %v", err)
	}
}