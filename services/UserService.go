package services

import (
	"crud-go-echo-gorm/models"
	"crud-go-echo-gorm/repositories"
	"fmt"

	myerrors "crud-go-echo-gorm/errors"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user models.User) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(userID string) (*models.User, error)
	UpdateUser(userID string, updatedUser models.User) (*models.User, error)
	DeleteUser(userID string) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(user models.User) (*models.User, error) {

	// Gere um novo UUID
	user.ID = uuid.New()

	return s.userRepository.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAllUsers()
}

func (s *userService) GetUserByID(userID string) (*models.User, error) {
	return s.userRepository.GetUserByID(userID)
}

func (s *userService) UpdateUser(userID string, updatedUser models.User) (*models.User, error) {
	// Tentar converter a string para UUID
	id, err := uuid.Parse(userID)
	if err != nil {
		fmt.Println("Erro ao converter a string para UUID:", err)
		return nil, err
	}

	if id == uuid.Nil {
		return nil, myerrors.ErrInvalidUserID
	}

	// Atualizar o ID do usuário com o uint convertido
	updatedUser.ID = id

	// Chamar a função de atualização no repositório
	_, err = s.userRepository.UpdateUser(updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (s *userService) DeleteUser(userID string) error {
	// Verificar se o ID é vazio ou inválido
	id, err := uuid.Parse(userID)
	if err != nil || id == uuid.Nil {
		return myerrors.ErrInvalidUserID
	}

	// Verificar se o usuário existe antes de tentar deletar
	existingUser, err := s.GetUserByID(userID)
	if err != nil {
		if err == myerrors.ErrUserNotFound {
			return myerrors.ErrUserNotFound
		}
		// Outro erro ao buscar o usuário
		return err
	}

	// O usuário existe, então podemos prosseguir com a exclusão
	if existingUser != nil {
		return s.userRepository.DeleteUser(userID)
	}

	// Se o existingUser for nil, significa que o usuário não foi encontrado
	return myerrors.ErrUserNotFound
}
