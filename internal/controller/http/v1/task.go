package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shareU/internal/service"
)

type taskRoutes struct {
	taskService service.Task
}

func newTaskRoutes(g *echo.Group, taskService service.Task) {
	r := &taskRoutes{
		taskService: taskService,
	}

	g.PUT("/create", r.create)
	g.POST("/allocate", r.allocate)
	g.POST("/submit", r.submit)
	//g.DELETE("/submit", r.delete)
	//g.GET("/get", r.getById)
	//g.GET("/getAll", r.getAll)
}

type taskCreateInput struct {
	ProjectId int    `json:"project-id" validate:"required"`
	Tag       string `json:"tag" validate:"required"`
	Command   string `json:"command" validate:"required"`
	Condition string `db:"condition"`
	After     string `db:"after"`
	Result    string `db:"result"`
	Priority  int    `json:"priority" validate:"required"`
}

func (r *taskRoutes) create(c echo.Context) error {
	var input taskCreateInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}
	err := r.taskService.Create(c.Request().Context(), service.TaskCreateInput{
		Project:   input.ProjectId,
		Tag:       input.Tag,
		Command:   input.Command,
		Condition: input.Condition,
		After:     input.After,
		Result:    input.Result,
		Priority:  input.Priority,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "task created",
	})
}

type taskAllocateInput struct {
	Ip string `json:"ip" validate:"required"`
}

func (r *taskRoutes) allocate(c echo.Context) error {
	var input taskAllocateInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	type taskAllocateResponse struct {
		TaskId    int    `json:"task-id"`
		Command   string `json:"command"`
		Condition string `db:"condition,omitempty"`
		After     string `db:"after,omitempty"`
		//DockerLink string `json:"docker-link,omitempty"`
	}

	out, err := r.taskService.Allocate(c.Request().Context(), service.TaskAllocateInput{
		AgentIP: input.Ip,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, taskAllocateResponse{
		TaskId:    out.TaskId,
		Command:   out.Command,
		Condition: out.Condition,
		After:     out.After,
		//DockerLink: out.DockerLink,
	})
}

type taskSubmitInput struct {
	Id     int    `json:"id" validate:"required"`
	Ip     string `json:"ip" validate:"required"`
	Result string `json:"result" validate:"required"` //TODO: should be file
}

func (r *taskRoutes) submit(c echo.Context) error {
	var input taskSubmitInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	err := r.taskService.Submit(c.Request().Context(), service.TaskSubmitInput{
		TaskId:      input.Id,
		AgentIP:     input.Ip,
		TerminalLog: input.Result,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
	})

}
