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
	query := `
		UPDATE task
		SET status = $1
		WHERE id = $2
	`
	_, err := t.Pool.Exec(ctx, query, status, id)
	return err
}

func (t TaskRepo) CreateTask(ctx context.Context, task entity.Task) error {
	// Implement the logic to insert a task into the database
	_, err := t.Pool.Exec(ctx, "INSERT INTO task (project_id, tag, command, priority) VALUES ($1, $2, $3, $4)", task.ProjectID, task.Tag, task.Command, task.Priority)
	return err
}

func (t TaskRepo) DeleteTaskById(ctx context.Context, id int) error {
	query := `
		DELETE FROM task
		WHERE id = $1
	`
	_, err := t.Pool.Exec(ctx, query, id)
	return err
}

func (t TaskRepo) GetAllTasks(ctx context.Context) ([]entity.Task, error) {
	query := `
		SELECT id, project_id, tag, command, priority, status
		FROM task
	`
	rows, err := t.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.ID, &task.ProjectID, &task.Tag, &task.Command, &task.Priority, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t TaskRepo) GetPriorityTask(ctx context.Context) (entity.Task, error) {
	query := `
		SELECT id, project_id, tag, command, priority, status
		FROM task
		WHERE status = 0
		ORDER BY priority DESC
		LIMIT 1
	`
	var task entity.Task
	err := t.Pool.QueryRow(ctx, query).Scan(&task.ID, &task.ProjectID, &task.Tag, &task.Command, &task.Priority, &task.Status)
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (t TaskRepo) GetTaskById(ctx context.Context, id int) (entity.Task, error) {
	query := `
		SELECT id, project_id, tag, command, priority, status
		FROM task
		WHERE id = $1
	`
	var task entity.Task
	err := t.Pool.QueryRow(ctx, query, id).Scan(&task.ID, &task.ProjectID, &task.Tag, &task.Command, &task.Priority, &task.Status)
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func NewTaskRepo(pg *postgres.Postgres) *TaskRepo {
	return &TaskRepo{pg}
}
