package routes

import (
	"crud-go-echo-gorm/controllers"
	"crud-go-echo-gorm/repositories"
	"crud-go-echo-gorm/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	// Crie instâncias dos seus controllers e passe o DB e os repositórios conforme necessário
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	orderRepository := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)

	// Configure as rotas para usuários
	userGroup := e.Group("/api/v1/users")
	userRoutes(userGroup, *userController)

	// Configure as rotas para ordens
	orderGroup := e.Group("/api/v1/orders")
	orderRoutes(orderGroup, orderController)
}

func userRoutes(g *echo.Group, uc controllers.UserController) {
	g.POST("/", uc.Create)
	g.GET("/", uc.FindAll)
	g.GET("/:id", uc.FindByID)
	g.PUT("/:id", uc.Update)
	g.DELETE("/:id", uc.Delete)
}

func orderRoutes(g *echo.Group, oc controllers.OrderController) {
	g.POST("/", oc.Create)
	g.GET("/", oc.FindAll)
	g.GET("/:id", oc.FindByID)
	g.PUT("/:id", oc.Update)
	g.DELETE("/:id", oc.Delete)
	// Adicione outras rotas conforme necessário
}
