package repositories

import (
	"crud-go-echo-gorm/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(userID string) (*models.User, error)
	UpdateUser(updatedUser models.User) (*models.User, error)
	DeleteUser(userID string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user models.User) (*models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", userID).Error
	return &user, err
}

func (r *userRepository) UpdateUser(updatedUser models.User) (*models.User, error) {
	err := r.db.Model(&models.User{}).Where("id = ?", updatedUser.ID).Updates(updatedUser).Scan(&updatedUser).Error
	return &updatedUser, err
}

func (r *userRepository) DeleteUser(userID string) error {
	err := r.db.Delete(&models.User{}, "id = ?", userID).Error
	if err != nil {
		fmt.Printf("Erro ao deletar usuário com ID %s: %s\n", userID, err)
	}
	return err
}
