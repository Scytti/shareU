package service

import (
	"context"
	"mime/multipart"
	"shareU/internal/repo"
)

type TaskCreateInput struct {
	project  string
	command  string
	priority int
}

type TaskAllocateInput struct {
	agentIP int
}

type TaskAllocateOutput struct {
	command    string
	dockerLink string
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

type ProjectCreateInput struct {
	project  string
	command  string
	priority int
}

type Project interface {
	Create(ctx context.Context, input ProjectCreateInput) error
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
