package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type taskRoutes struct {
	//taskService service.Task
}

func newTaskRoutes(g *echo.Group /*, taskService service.Task*/) {
	r := &taskRoutes{
		//taskService: taskService,
	}

	g.PUT("/create", r.create)
	g.POST("/allocate", r.allocate)
	//g.POST("/submit", r.submit)
	//g.DELETE("/submit", r.delete)
	//g.GET("/get", r.getById)
	//g.GET("/getAll", r.getAll)
}

type taskCreateInput struct {
	ProjectId int    `json:"project-id" validate:"required"`
	Tag       string `json:"tag" validate:"required"`
	Command   string `json:"command" validate:"required"`
	Priority  int    `json:"priority" validate:"required"`
}

func (r *taskRoutes) create(c echo.Context) error {
	/*var input taskCreateInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}
	err := r.taskService.Create(c.Request().Context(), service.TaskCreateInput{
		Project:  input.ProjectId,
		Tag:      input.Tag,
		Command:  input.Command,
		Priority: input.Priority,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}*/

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
	})
}

type taskAllocateInput struct {
	Ip int `json:"ip" validate:"required"`
}

func (r *taskRoutes) allocate(c echo.Context) error {
	/*var input taskAllocateInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	type taskAllocateResponse struct {
		Command    string `json:"command"`
		DockerLink string `json:"docker-link,omitempty"`
	}

	out, err := r.taskService.Allocate(c.Request().Context(), service.TaskAllocateInput{
		AgentIP: input.Ip,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, taskAllocateResponse{
		Command:    out.Command,
		DockerLink: out.DockerLink,
	})*/

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
	})
}
