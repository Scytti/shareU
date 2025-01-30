package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"shareU/internal/entity"
	"shareU/internal/repo"
)

type TaskService struct {
	taskRepo repo.Task
}

func (t TaskService) Allocate(ctx context.Context, input TaskAllocateInput) (TaskAllocateOutput, error) {
	println("Получаем таску")
	task, err := t.taskRepo.GetPriorityTask(ctx)

	println("Меняем ее статус")

	if err != nil {
		if err := t.taskRepo.ChangeTaskStatus(ctx, task.ID, 2); err != nil {
			return TaskAllocateOutput{}, err
		}
	}

	println("Добавялем в лог")

	if err := t.taskRepo.AddToLogTask(ctx, task.ID, input.AgentIP, 2, ""); err != nil {
		return TaskAllocateOutput{}, err
	}

	if err != nil {
		return TaskAllocateOutput{}, err
	}

	return TaskAllocateOutput{
		TaskId:    task.ID,
		Command:   task.Command,
		Condition: task.Condition,
		After:     task.After,
		//DockerLink: "",
	}, nil
}

func (t TaskService) Submit(ctx context.Context, input TaskSubmitInput) error {
	if err := t.taskRepo.ChangeTaskStatus(ctx, input.TaskId, 3); err != nil {
		return err
	}

	if err := t.taskRepo.AddToLogTask(ctx, input.TaskId, input.AgentIP, 3, input.TerminalLog); err != nil {
		return err
	}

	return nil
}

func NewTaskService(taskRepo repo.Task) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (t *TaskService) Create(ctx context.Context, input TaskCreateInput) error {
	log.Info("Зашли в создание таски не сервисе")

	task := entity.Task{
		ProjectID: input.Project,
		Tag:       input.Tag,
		Command:   input.Command,
		Condition: input.Condition,
		After:     input.After,
		Result:    input.Result,
		Priority:  input.Priority,
	}
	return t.taskRepo.CreateTask(ctx, task)
}
