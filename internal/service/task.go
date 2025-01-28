package service

import (
	"context"
	"shareU/internal/entity"
	"shareU/internal/repo"
)

type TaskService struct {
	taskRepo repo.Task
}

func (t TaskService) Allocate(ctx context.Context, input TaskAllocateInput) (TaskAllocateOutput, error) {
	task, err := t.taskRepo.GetPriorityTask(ctx)
	if err != nil {
		return TaskAllocateOutput{}, err
	}

	// Assuming the agent IP is used to allocate the task
	// You can add more logic here to handle the allocation
	return TaskAllocateOutput{
		Command:    task.Command,
		DockerLink: "", // Assuming DockerLink is part of the task entity
	}, nil
}

func (t TaskService) Submit(ctx context.Context, input TaskCreateInput) error {
	return t.taskRepo.ChangeTaskStatus(ctx, input.Project, 1)
}

func NewTaskService(taskRepo repo.Task) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (t *TaskService) Create(ctx context.Context, input TaskCreateInput) error {
	task := entity.Task{
		ProjectID: input.Project,
		Tag:       input.Tag,
		Command:   input.Command,
		Priority:  input.Priority,
	}
	return t.taskRepo.CreateTask(ctx, task)
}
