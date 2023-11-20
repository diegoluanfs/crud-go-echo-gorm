package controllers

import (
	"crud-go-echo-gorm/models"
	"crud-go-echo-gorm/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// OrderController representa a interface para a camada de controle de pedidos
type OrderController interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	FindByID(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

// OrderControllerImpl é a implementação da interface OrderController
type OrderControllerImpl struct {
	OrderService services.OrderService
}

// NewOrderController cria uma nova instância do OrderControllerImpl
func NewOrderController(orderService services.OrderService) OrderController {
	return &OrderControllerImpl{OrderService: orderService}
}

// Implementações de métodos da interface OrderController
func (oc *OrderControllerImpl) Create(c echo.Context) error {
	var order models.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	createdOrder, err := oc.OrderService.CreateOrder(&order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create order"})
	}

	return c.JSON(http.StatusCreated, createdOrder)
}

func (oc *OrderControllerImpl) FindAll(c echo.Context) error {
	orders, err := oc.OrderService.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch orders"})
	}

	return c.JSON(http.StatusOK, orders)
}

func (oc *OrderControllerImpl) FindByID(c echo.Context) error {
	orderID := c.Param("id")
	order, err := oc.OrderService.GetOrderByID(orderID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}

	return c.JSON(http.StatusOK, order)
}

func (oc *OrderControllerImpl) Update(c echo.Context) error {
	orderID := c.Param("id")

	var updatedOrder *models.Order
	if err := c.Bind(&updatedOrder); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	updatedOrder, err := oc.OrderService.UpdateOrder(orderID, updatedOrder)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update order"})
	}

	return c.JSON(http.StatusOK, updatedOrder)
}

func (oc *OrderControllerImpl) Delete(c echo.Context) error {
	orderID := c.Param("id")
	if err := oc.OrderService.DeleteOrder(orderID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete order"})
	}

	return c.NoContent(http.StatusNoContent)
}
