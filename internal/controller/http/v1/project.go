package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shareU/internal/service"
)

type projectRoutes struct {
	projectService service.Project
}

func newProjectRoutes(g *echo.Group, projectService service.Project) {
	r := &projectRoutes{
		projectService: projectService,
	}

	g.POST("/create", r.create)
	g.POST("/delete", r.delete)
	g.GET("/get", r.getAll)
}

type projectCreateInput struct {
	Name string `json:"name"`
}

func (r *projectRoutes) create(c echo.Context) error {
	println("Пытаемся записать")
	var input projectCreateInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}
	err := r.projectService.Create(c.Request().Context(), input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
	})
}

type projectDeleteInput struct {
	Id int `json:"id"`
}

func (r *projectRoutes) delete(c echo.Context) error {
	var input projectDeleteInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	err := r.projectService.Delete(c.Request().Context(), input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (r *projectRoutes) getAll(c echo.Context) error {

	type projectGetResponse struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	projects, err := r.projectService.Get(c.Request().Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	response := make([]projectGetResponse, 0, len(projects))
	for _, project := range projects {
		response = append(response, projectGetResponse{
			Id:   project.Id,
			Name: project.Name,
		})
	}

	return c.JSON(http.StatusOK, response)
}
