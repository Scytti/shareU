package pgdb

import (
	"context"
	"shareU/internal/entity"
	"shareU/pkg/postgres"
)

type ProjectRepo struct {
	*postgres.Postgres
}

func (p ProjectRepo) CreateProject(ctx context.Context, name string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectRepo) DeleteProjectById(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (p ProjectRepo) GetProjectById(ctx context.Context, id int) (entity.Project, error) {
	//TODO implement me
	panic("implement me")
}

func NewProjectRepo(pg *postgres.Postgres) *ProjectRepo {
	return &ProjectRepo{pg}
}
