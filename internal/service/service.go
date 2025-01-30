package service

import (
	"context"
	"shareU/internal/repo"
)

type TaskCreateInput struct {
	Project   int
	Tag       string
	Command   string
	Condition string
	After     string
	Result    string
	Priority  int
}

type TaskAllocateInput struct {
	AgentIP string
}

type TaskAllocateOutput struct {
	TaskId    int
	Command   string
	Condition string
	After     string
	//DockerLink string
}

type TaskSubmitInput struct {
	TaskId      int
	AgentIP     string
	TerminalLog string
	//terminalLog multipart.File
}

type Task interface {
	Allocate(ctx context.Context, input TaskAllocateInput) (TaskAllocateOutput, error)
	Submit(ctx context.Context, input TaskSubmitInput) error
	Create(ctx context.Context, input TaskCreateInput) error
}

type ProjectGetOutput struct {
	Id   int
	Name string
}

type Project interface {
	Create(ctx context.Context, name string) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context) ([]ProjectGetOutput, error)
}

type Services struct {
	Task    Task
	Project Project
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Task:    NewTaskService(deps.Repos.Task),
		Project: NewProjectService(deps.Repos.Project),
	}
}
