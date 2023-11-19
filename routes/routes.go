package routes

import (
	"crud-go-echo-gorm/controllers"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	// Crie instâncias dos seus controllers e passe o DB conforme necessário
	userController := controllers.NewUserController(db)
	orderController := controllers.NewOrderController(db)

	// Configure as rotas para usuários
	userGroup := e.Group("/api/v1/users")
	userRoutes(userGroup, userController)

	// Configure as rotas para ordens
	orderGroup := e.Group("/api/v1/orders")
	orderRoutes(orderGroup, orderController)
}

func userRoutes(g *echo.Group, uc *controllers.UserController) {
	g.POST("/", uc.Create)
	g.GET("/", uc.FindAll)
	g.GET("/:id", uc.FindByID)
	g.PUT("/:id", uc.Update)
	g.DELETE("/:id", uc.Delete)
}

func orderRoutes(g *echo.Group, oc *controllers.OrderController) {
	g.POST("/", oc.Create)
	g.GET("/", oc.FindAll)
	g.GET("/:id", oc.FindByID)
	g.PUT("/:id", oc.Update)
	g.DELETE("/:id", oc.Delete)
	// Adicione outras rotas conforme necessário
}
