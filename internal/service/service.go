package service

import (
	"context"
	"mime/multipart"
	"shareU/internal/repo"
)

type TaskCreateInput struct {
	Project  int
	Tag      string
	Command  string
	Priority int
}

type TaskAllocateInput struct {
	AgentIP int
}

type TaskAllocateOutput struct {
	Command    string
	DockerLink string
}

type TaskSubmitInput struct {
	agentIP     int
	terminalLog multipart.File
}

type Task interface {
	Allocate(ctx context.Context, input TaskAllocateInput) (TaskAllocateOutput, error)
	Submit(ctx context.Context, input TaskCreateInput) error
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
