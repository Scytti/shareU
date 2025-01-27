package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shareU/internal/service"
)

func NewRouter(handler *echo.Echo, services *service.Services) {

	handler.Use(middleware.Recover())

	handler.GET("/health", func(c echo.Context) error { return c.NoContent(200) })

	auth := handler.Group("/auth")
	{
		newAuthRoutes(auth, services.Auth)
	}

	v1 := handler.Group("/api/v1")
	{
		newAccountRoutes(v1.Group("/accounts"), services.Account)
		newReservationRoutes(v1.Group("/reservations"), services.Reservation)
		newProductRoutes(v1.Group("/products"), services.Product)
		newOperationRoutes(v1.Group("/operations"), services.Operation)
	}
}
