package pgdb

import (
	"context"
	"log/slog"
	"shareU/internal/entity"
	"shareU/pkg/postgres"
)

type ProjectRepo struct {
	*postgres.Postgres
	log *slog.Logger
}

func NewProjectRepo(pg *postgres.Postgres, log *slog.Logger) *ProjectRepo {
	return &ProjectRepo{pg, log}
}

func (p ProjectRepo) GetProject(ctx context.Context) ([]entity.Project, error) {
	rows, err := p.Pool.Query(ctx, "SELECT id, name FROM project")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []entity.Project
	for rows.Next() {
		var project entity.Project
		if err := rows.Scan(&project.ID, &project.Name); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (p ProjectRepo) CreateProject(ctx context.Context, name string) (int, error) {
	var id int
	println("пытаемся выполнить sql")
	err := p.Pool.QueryRow(ctx, "INSERT INTO project (name) VALUES ($1) RETURNING id", name).Scan(&id)
	println(err)
	return id, err
}

func (p ProjectRepo) DeleteProjectById(ctx context.Context, id int) error {
	res, err := p.Pool.Exec(ctx, "DELETE FROM project WHERE id = $1", id)
	if err != nil {
		// Обработка ошибки при выполнении запроса
		p.log.Error("Error executing delete query: %v", err)
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		p.log.Error("No project found with id: %d", id)
		return nil
	}

	p.log.Debug("Project with id %d deleted successfully", id)
	return nil
}

/*func (p ProjectRepo) GetProjectById(ctx context.Context, id int) (entity.Project, error) {
	//TODO implement me
	panic("implement me")
}*/
