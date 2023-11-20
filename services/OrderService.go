package services

import (
	"crud-go-echo-gorm/models"
	"crud-go-echo-gorm/repositories"
)

// OrderService representa a interface para a camada de serviço de pedidos
type OrderService interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(orderID string) (*models.Order, error)
	UpdateOrder(orderID string, updatedOrder *models.Order) (*models.Order, error)
	DeleteOrder(orderID string) error
}

// OrderServiceImpl é a implementação da interface OrderService
type OrderServiceImpl struct {
	OrderRepository repositories.OrderRepository
}

// NewOrderService cria uma nova instância do OrderServiceImpl
func NewOrderService(orderRepository repositories.OrderRepository) OrderService {
	return &OrderServiceImpl{OrderRepository: orderRepository}
}

// Implementações de métodos da interface OrderService
func (os *OrderServiceImpl) CreateOrder(order *models.Order) (*models.Order, error) {
	return os.OrderRepository.CreateOrder(*order)
}

func (os *OrderServiceImpl) GetAllOrders() ([]models.Order, error) {
	return os.OrderRepository.GetAllOrders()
}

func (os *OrderServiceImpl) GetOrderByID(orderID string) (*models.Order, error) {
	return os.OrderRepository.GetOrderByID(orderID)
}

func (os *OrderServiceImpl) UpdateOrder(orderID string, updatedOrder *models.Order) (*models.Order, error) {
	return os.OrderRepository.UpdateOrder(orderID, *updatedOrder)
}

func (os *OrderServiceImpl) DeleteOrder(orderID string) error {
	return os.OrderRepository.DeleteOrder(orderID)
}
