package repo

import (
	"context"
	"log/slog"
	"shareU/internal/entity"
	"shareU/internal/repo/pgdb"
	"shareU/pkg/postgres"
)

type Project interface {
	CreateProject(ctx context.Context, name string) (int, error)
	DeleteProjectById(ctx context.Context, id int) error
	//GetProjectById(ctx context.Context, id int) (entity.Project, error)
	GetProject(ctx context.Context) ([]entity.Project, error)
}

type Task interface {
	ChangeTaskStatus(ctx context.Context, id int, status int) error
	CreateTask(ctx context.Context, task entity.Task) error
	DeleteTaskById(ctx context.Context, id int) error
	GetAllTasks(ctx context.Context) ([]entity.Task, error)
	GetPriorityTask(ctx context.Context) (entity.Task, error)
	GetTaskById(ctx context.Context, id int) (entity.Task, error)
	AddToLogTask(ctx context.Context, task_id int, ip string, status int, result string) error
}

type Repositories struct {
	Task    Task
	Project Project
}

func NewRepositories(pg *postgres.Postgres, log *slog.Logger) *Repositories {
	return &Repositories{
		Task:    pgdb.NewTaskRepo(pg),
		Project: pgdb.NewProjectRepo(pg, log),
	}
}
