package main

import (
	"log"
	"os"

	"github.com/avanti-dvp/ms-saudacoes-aleatorias/database"
	"github.com/avanti-dvp/ms-saudacoes-aleatorias/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa a conexão com o banco de dados PostgreSQL
	database.ConnectDatabase()

	// Cria um router Gin com as configurações padrão
	router := gin.Default()

	// Configuração de CORS simplificada
	router.Use(cors.Default())

	// --- NOVA ROTA RAIZ ---
	// Resolve o erro "404 page not found" ao acessar o link puro
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "online",
			"message": "Bem-vindo à API de Saudações Aleatórias, Coimbra!",
			"endpoints": []string{
				"/api/saudacoes/aleatorio",
				"/api/saudacoes",
			},
		})
	})

	// Grupo de rotas da API
	api := router.Group("/api")
	{
		// Rota para cadastrar um novo cumprimento
		api.POST("/saudacoes", handlers.CreateGreeting)

		// Rota para obter um cumprimento aleatório
		api.GET("/saudacoes/aleatorio", handlers.GetRandomGreeting)
	}

	// Lógica para detectar a porta do ambiente (Render)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback local
	}

	log.Printf("Iniciando servidor na porta %s...", port)

	// Inicia o servidor
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Falha ao rodar o servidor: %v", err)
	}
}