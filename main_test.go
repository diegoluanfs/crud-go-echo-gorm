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
	fmt.Println("Iniciando testes...")
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
	fmt.Println("Testes concluídos.")
	os.Exit(code)
}

func TestUsersEndpoints(t *testing.T) {
	fmt.Println("Iniciando teste de endpoints de usuários...")

	// Cria uma instância do Echo
	e := echo.New()

	// Configuração do fuso horário (opcional)
	if tz := os.Getenv("TZ"); tz != "" {
		os.Setenv("TZ", tz)
	}

	// Configuração das rotas
	routes.SetupRoutes(e, db)

	// Teste do endpoint POST /api/v1/users
	fmt.Println("Iniciando teste do endpoint POST /api/v1/users")
	reqCreate := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/", nil)
	recCreate := httptest.NewRecorder()
	e.ServeHTTP(recCreate, reqCreate)
	assert.Equal(t, http.StatusCreated, recCreate.Code)
	fmt.Println("Teste do endpoint POST /api/v1/users concluído")

	// Teste do endpoint GET /api/v1/users
	fmt.Println("Iniciando teste do endpoint GET /api/v1/users")
	reqFindAll := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users/", nil)
	recFindAll := httptest.NewRecorder()
	e.ServeHTTP(recFindAll, reqFindAll)
	assert.Equal(t, http.StatusOK, recFindAll.Code)
	fmt.Println("Teste do endpoint GET /api/v1/users concluído")

	// Teste do endpoint GET /api/v1/users/:id
	fmt.Println("Iniciando teste do endpoint GET /api/v1/users/:id")
	reqFindOne := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users/1", nil)
	recFindOne := httptest.NewRecorder()
	e.ServeHTTP(recFindOne, reqFindOne)
	assert.Equal(t, http.StatusOK, recFindOne.Code)
	fmt.Println("Teste do endpoint GET /api/v1/users/:id concluído")

	// Teste do endpoint PUT /api/v1/users/:id
	fmt.Println("Iniciando teste do endpoint PUT /api/v1/users/:id")
	reqUpdate := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/users/1", nil)
	recUpdate := httptest.NewRecorder()
	e.ServeHTTP(recUpdate, reqUpdate)
	assert.Equal(t, http.StatusOK, recUpdate.Code)
	fmt.Println("Teste do endpoint PUT /api/v1/users/:id concluído")

	// Teste do endpoint DELETE /api/v1/users/:id
	fmt.Println("Iniciando teste do endpoint DELETE /api/v1/users/:id")
	reqDelete := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/users/1", nil)
	recDelete := httptest.NewRecorder()
	e.ServeHTTP(recDelete, reqDelete)
	assert.Equal(t, http.StatusNoContent, recDelete.Code)
	fmt.Println("Teste do endpoint DELETE /api/v1/users/:id concluído")

	fmt.Println("Teste de endpoints de usuários concluído.")
}

func TestOrdersEndpoints(t *testing.T) {
	fmt.Println("Iniciando teste de endpoints de pedidos...")

	// Cria uma instância do Echo
	e := echo.New()

	// Configuração do fuso horário (opcional)
	if tz := os.Getenv("TZ"); tz != "" {
		os.Setenv("TZ", tz)
	}

	// Configuração das rotas
	routes.SetupRoutes(e, db)

	// Teste do endpoint POST /api/v1/orders
	fmt.Println("Iniciando teste do endpoint POST /api/v1/orders")
	reqCreate := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/orders/", nil)
	recCreate := httptest.NewRecorder()
	e.ServeHTTP(recCreate, reqCreate)
	assert.Equal(t, http.StatusCreated, recCreate.Code)
	fmt.Println("Teste do endpoint POST /api/v1/orders concluído")

	// Teste do endpoint GET /api/v1/orders
	fmt.Println("Iniciando teste do endpoint GET /api/v1/orders")
	reqFindAll := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/orders/", nil)
	recFindAll := httptest.NewRecorder()
	e.ServeHTTP(recFindAll, reqFindAll)
	assert.Equal(t, http.StatusOK, recFindAll.Code)
	fmt.Println("Teste do endpoint GET /api/v1/orders concluído")

	// Teste do endpoint GET /api/v1/orders/:id
	fmt.Println("Iniciando teste do endpoint GET /api/v1/orders/:id")
	reqFindOne := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/orders/1", nil)
	recFindOne := httptest.NewRecorder()
	e.ServeHTTP(recFindOne, reqFindOne)
	assert.Equal(t, http.StatusOK, recFindOne.Code)
	fmt.Println("Teste do endpoint GET /api/v1/orders/:id concluído")

	// Teste do endpoint PUT /api/v1/orders/:id
	fmt.Println("Iniciando teste do endpoint PUT /api/v1/orders/:id")
	reqUpdate := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/orders/1", nil)
	recUpdate := httptest.NewRecorder()
	e.ServeHTTP(recUpdate, reqUpdate)
	assert.Equal(t, http.StatusOK, recUpdate.Code)
	fmt.Println("Teste do endpoint PUT /api/v1/orders/:id concluído")

	// Teste do endpoint DELETE /api/v1/orders/:id
	fmt.Println("Iniciando teste do endpoint DELETE /api/v1/orders/:id")
	reqDelete := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/orders/1", nil)
	recDelete := httptest.NewRecorder()
	e.ServeHTTP(recDelete, reqDelete)
	assert.Equal(t, http.StatusNoContent, recDelete.Code)
	fmt.Println("Teste do endpoint DELETE /api/v1/orders/:id concluído")

	fmt.Println("Teste de endpoints de pedidos concluído.")
}
