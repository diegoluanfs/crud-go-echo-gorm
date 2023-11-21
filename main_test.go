package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"crud-go-echo-gorm/database"
	"crud-go-echo-gorm/routes"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	fmt.Println("Iniciando testes...")

	if err := godotenv.Load(); err != nil {
		panic("Erro carregando o arquivo .env")
	}

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

	code := m.Run()

	fmt.Println("Testes concluídos.")
	os.Exit(code)
}

func TestUsersEndpoints(t *testing.T) {
	e := echo.New()

	if tz := os.Getenv("TZ"); tz != "" {
		os.Setenv("TZ", tz)
	}

	routes.SetupRoutes(e, db)

	testEndpoint("Get All", e, t, GetAll)
	testEndpoint("Create", e, t, Create)
	testEndpoint("Get All", e, t, GetAll)

	lastUser, err := GetLastUser(e, t)
	if err != nil {
		t.Fatalf("Erro ao obter o último usuário: %v", err)
	}
	testEndpoint("Get By ID", e, t, func(e *echo.Echo, t *testing.T) (string, error) {
		return GetById(e, t, lastUser)
	})

	firstUser, err := GetFirstUser(e, t)
	if err != nil {
		t.Fatalf("Erro ao obter o primeiro usuário: %v", err)
	}
	testEndpoint("Delete User", e, t, func(e *echo.Echo, t *testing.T) (string, error) {
		return DeleteById(e, t, firstUser)
	})
}

func testEndpoint(name string, e *echo.Echo, t *testing.T, fn func(e *echo.Echo, t *testing.T) (string, error)) {
	// fmt.Printf("----------------------------------------\n")
	// fmt.Printf("-------------- %s -----------------\n", name)

	result, err := fn(e, t)
	if err != nil {
		t.Fatalf("Erro ao %s: %v", name, err)
	}

	fmt.Println(result)
	// fmt.Printf("----------------------------------------\n\n")
}

func GetAll(e *echo.Echo, t *testing.T) (string, error) {
	reqFindAll := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users/", nil)
	recFindAll := httptest.NewRecorder()
	e.ServeHTTP(recFindAll, reqFindAll)

	if recFindAll.Code != http.StatusOK {
		return "", fmt.Errorf("Erro no código de status: %d", recFindAll.Code)
	}

	var result interface{}
	err := json.Unmarshal(recFindAll.Body.Bytes(), &result)
	if err != nil {
		return "", fmt.Errorf("Erro ao analisar o JSON retornado: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Erro ao formatar o JSON: %v", err)
	}

	return string(prettyJSON), nil
}

func GetLastUser(e *echo.Echo, t *testing.T) (map[string]interface{}, error) {
	reqFindAll := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users/", nil)
	recFindAll := httptest.NewRecorder()
	e.ServeHTTP(recFindAll, reqFindAll)

	if recFindAll.Code != http.StatusOK {
		return nil, fmt.Errorf("Erro ao obter usuários. Código de status: %d", recFindAll.Code)
	}

	var userList []map[string]interface{}
	err := json.Unmarshal(recFindAll.Body.Bytes(), &userList)
	if err != nil {
		return nil, fmt.Errorf("Erro ao analisar o JSON retornado: %v", err)
	}

	if len(userList) > 0 {
		lastUser := userList[len(userList)-1]
		return lastUser, nil
	}

	return nil, fmt.Errorf("Nenhum usuário encontrado.")
}

func GetFirstUser(e *echo.Echo, t *testing.T) (map[string]interface{}, error) {
	reqFindAll := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users/", nil)
	recFindAll := httptest.NewRecorder()
	e.ServeHTTP(recFindAll, reqFindAll)

	if recFindAll.Code != http.StatusOK {
		return nil, fmt.Errorf("Erro ao obter usuários. Código de status: %d", recFindAll.Code)
	}

	var userList []map[string]interface{}
	err := json.Unmarshal(recFindAll.Body.Bytes(), &userList)
	if err != nil {
		return nil, fmt.Errorf("Erro ao analisar o JSON retornado: %v", err)
	}

	if len(userList) > 0 {
		firstUser := userList[0]
		return firstUser, nil
	}

	return nil, fmt.Errorf("Nenhum usuário encontrado.")
}

func Create(e *echo.Echo, t *testing.T) (string, error) {
	const (
		baseURL     = "http://localhost:8080/api/v1/users/"
		letters     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		emailDomain = "@teste.com"
		length      = 5
		phoneNumber = "1234567899"
	)

	user := generateRandomUser(length, letters, emailDomain)

	requestBodyJSON, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("Erro ao criar o corpo JSON: %v", err)
	}

	reqCreate := httptest.NewRequest(
		http.MethodPost,
		baseURL,
		bytes.NewBuffer(requestBodyJSON),
	)
	reqCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recCreate := httptest.NewRecorder()
	e.ServeHTTP(recCreate, reqCreate)

	res := recCreate.Result()
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Erro ao ler o corpo da resposta HTTP: %v", err)
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Response: %s\n", body)

	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("Erro no código de status: %d", res.StatusCode)
	}

	return string(body), nil
}

func generateRandomUser(length int, letters, emailDomain string) map[string]interface{} {
	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	email := string(result) + emailDomain

	return map[string]interface{}{
		"name":         string(result),
		"email":        email,
		"phone_number": "1234567899",
	}
}

func GetById(e *echo.Echo, t *testing.T, userID map[string]interface{}) (string, error) {
	// Tentar obter o ID com a chave 'ID' ou 'id'
	var userIDValue interface{}
	var ok bool
	if userIDValue, ok = userID["ID"]; !ok {
		if userIDValue, ok = userID["id"]; !ok {
			return "", fmt.Errorf("Campo 'ID' ou 'id' não encontrado no mapa")
		}
	}

	// Verificar se o ID é uma string
	userIDString, ok := userIDValue.(string)
	if !ok {
		return "", fmt.Errorf("Campo 'ID' não é uma string")
	}

	// Verificar se o ID é um UUID válido
	_, err := uuid.Parse(userIDString)
	if err != nil {
		return "", fmt.Errorf("Erro ao converter o ID para UUID: %v", err)
	}

	url := fmt.Sprintf("http://localhost:8080/api/v1/users/%s", userIDString)

	reqFindOne := httptest.NewRequest(http.MethodGet, url, nil)
	recFindOne := httptest.NewRecorder()

	e.ServeHTTP(recFindOne, reqFindOne)

	if recFindOne.Code != http.StatusOK {
		return "", fmt.Errorf("Erro no código de status: %d", recFindOne.Code)
	}

	return recFindOne.Body.String(), nil
}

func DeleteById(e *echo.Echo, t *testing.T, userID map[string]interface{}) (string, error) {
	// Tentar obter o ID com a chave 'ID' ou 'id'
	var userIDValue interface{}
	var ok bool
	if userIDValue, ok = userID["ID"]; !ok {
		if userIDValue, ok = userID["id"]; !ok {
			return "", fmt.Errorf("Campo 'ID' ou 'id' não encontrado no mapa")
		}
	}

	// Verificar se o ID é uma string
	userIDString, ok := userIDValue.(string)
	if !ok {
		return "", fmt.Errorf("Campo 'ID' não é uma string")
	}

	// Verificar se o ID é um UUID válido
	_, err := uuid.Parse(userIDString)
	if err != nil {
		return "", fmt.Errorf("Erro ao converter o ID para UUID: %v", err)
	}

	url := fmt.Sprintf("http://localhost:8080/api/v1/users/%s", userIDString)

	reqDelete := httptest.NewRequest(http.MethodDelete, url, nil)
	recDelete := httptest.NewRecorder()

	e.ServeHTTP(recDelete, reqDelete)

	if recDelete.Code != http.StatusOK {
		return "", fmt.Errorf("Erro no código de status: %d", recDelete.Code)
	}

	return recDelete.Body.String(), nil
}
