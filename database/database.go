// database.go

package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config contém as configurações do banco de dados
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// ComposeDsn cria a string de conexão do banco de dados
func ComposeDsn(config Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
}

// NewDB cria uma nova instância do Gorm para interagir com o banco de dados
func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// NewTestDB cria uma nova instância do Gorm para o banco de dados de teste
func NewTestDB() (*gorm.DB, error) {
	// Configurações específicas para o banco de dados de teste
	testConfig := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME") + "_test", // Adiciona um sufixo para o banco de dados de teste, caso queira
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	dsn := ComposeDsn(testConfig)
	return NewDB(dsn)
}
