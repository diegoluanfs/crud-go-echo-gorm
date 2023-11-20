package repositories

import (
	"crud-go-echo-gorm/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order models.Order) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(orderID string) (*models.Order, error)
	UpdateOrder(orderID string, updatedOrder models.Order) (*models.Order, error)
	DeleteOrder(orderID string) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order models.Order) (*models.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetOrderByID(orderID string) (*models.Order, error) {
	var order models.Order
	err := r.db.First(&order, "id = ?", orderID).Error
	return &order, err
}

func (r *orderRepository) UpdateOrder(orderID string, updatedOrder models.Order) (*models.Order, error) {
	var order models.Order
	err := r.db.First(&order, "id = ?", orderID).Updates(updatedOrder).Error
	return &order, err
}

func (r *orderRepository) DeleteOrder(orderID string) error {
	return r.db.Delete(&models.Order{}, "id = ?", orderID).Error
}
