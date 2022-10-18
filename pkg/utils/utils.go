package utils

import (
	"github.com/fa9566509/POSBackEnd/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func RegisterAllRoutes(app *fiber.App) {
	routes.RegisterCategoryRoutes(app)
	routes.RegisterCustomerRoutes(app)
	routes.RegisterOrderItemRoutes(app)
	routes.RegisterOrderRoutes(app)
	routes.RegisterProductRoutes(app)
	routes.RegisterPurchaseItemRoutes(app)
	routes.RegisterPurchaseRoutes(app)
	routes.RegisterRoleRoutes(app)
	routes.RegisterStockRoutes(app)
	routes.RegisterStoreRoutes(app)
	routes.RegisterSupplierRoutes(app)
	routes.RegisterUserRoutes(app)
}