package main

import (
	"crud-go-echo-gorm/database"
	"crud-go-echo-gorm/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		// Se ocorrer um erro ao carregar o arquivo .env, você pode tratar aqui
		panic("Erro carregando o arquivo .env")
	}

	// Obtém a porta do ambiente ou usa 8080 como padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Cria uma instância do Echo
	e := echo.New()

	// Configuração do fuso horário (opcional)
	if tz := os.Getenv("TZ"); tz != "" {
		os.Setenv("TZ", tz)
	}

	dsn := database.ComposeDsn(database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})

	// Cria uma instância do gorm.DB
	db, err := database.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Certifique-se de fechar a conexão do banco de dados ao sair
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// Configuração das rotas
	routes.SetupRoutes(e, db)

	// Inicia o servidor na porta configurada
	e.Start(":" + port)
}
