package pgdb

import (
	"context"
	"shareU/internal/entity"
	"shareU/pkg/postgres"
)

type TaskRepo struct {
	*postgres.Postgres
}

func (t TaskRepo) ChangeTaskStatus(ctx context.Context, id int, status int) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) CreateTask(ctx context.Context, task entity.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) DeleteTaskById(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) GetAllTasks(ctx context.Context) ([]entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) GetPriorityTask(ctx context.Context) (entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepo) GetTaskById(ctx context.Context, id int) (entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func NewTaskRepo(pg *postgres.Postgres) *TaskRepo {
	return &TaskRepo{pg}
}
