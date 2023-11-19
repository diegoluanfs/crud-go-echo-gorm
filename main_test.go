package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"crud-go-echo-gorm/database"
	"crud-go-echo-gorm/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB // Variável global para armazenar a instância do banco de dados

func TestMain(m *testing.M) {
	// Configurações antes de rodar os testes
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		// Se ocorrer um erro ao carregar o arquivo .env, você pode tratar aqui
		panic("Erro carregando o arquivo .env")
	}

	// Configuração do banco de dados de teste
	var err error
	db, err = database.NewTestDB()
	if err != nil {
		panic("Erro ao configurar o banco de dados de teste")
	}
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// Chamada para rodar os testes
	code := m.Run()

	// Limpeza após rodar os testes (se necessário)
	os.Exit(code)
}

func TestUsersEndpoint(t *testing.T) {
	// Cria uma instância do Echo
	e := echo.New()

	// Configuração do fuso horário (opcional)
	if tz := os.Getenv("TZ"); tz != "" {
		os.Setenv("TZ", tz)
	}

	// Configuração das rotas
	routes.SetupRoutes(e, db)

	// Cria uma solicitação HTTP de teste para o endpoint /api/v1/users
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users/", nil)
	rec := httptest.NewRecorder()

	// Execute a solicitação
	e.ServeHTTP(rec, req)

	// Verifique o código de status da resposta
	assert.Equal(t, http.StatusOK, rec.Code)

	// Log do teste
	fmt.Println("TestUsersEndpoint: Status Code =", rec.Code)
}

func TestOrdersEndpoint(t *testing.T) {
	// Cria uma instância do Echo
	e := echo.New()

	// Configuração do fuso horário (opcional)
	if tz := os.Getenv("TZ"); tz != "" {
		os.Setenv("TZ", tz)
	}

	// Configuração das rotas
	routes.SetupRoutes(e, db)

	// Cria uma solicitação HTTP de teste para o endpoint /api/v1/orders
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/orders/", nil)
	rec := httptest.NewRecorder()

	// Execute a solicitação
	e.ServeHTTP(rec, req)

	// Verifique o código de status da resposta
	assert.Equal(t, http.StatusOK, rec.Code)

	// Log do teste
	fmt.Println("TestOrdersEndpoint: Status Code =", rec.Code)
}
