package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shareU/internal/service"
)

func NewRouter(handler *echo.Echo, services *service.Services) {

	handler.Use(middleware.Recover())

	handler.GET("/health", func(c echo.Context) error { return c.NoContent(200) })

	v1 := handler.Group("/api/v1")
	{
		newTaskRoutes(v1.Group("/tasks"), services.Task)
		newProjectRoutes(v1.Group("/projects"), services.Project)
	}
}
