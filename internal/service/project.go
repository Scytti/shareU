package service

import (
	"context"
	"shareU/internal/repo"
)

type ProjectService struct {
	projectRepo repo.Project
}

func (p *ProjectService) Create(ctx context.Context, name string) error {
	// Implement the logic to create a new project
	println("Зашли в репо")
	_, err := p.projectRepo.CreateProject(ctx, name)
	return err
}

func (p *ProjectService) Delete(ctx context.Context, id int) error {
	// Implement the logic to delete a project
	return p.projectRepo.DeleteProjectById(ctx, id)
}

func (p *ProjectService) Get(ctx context.Context) ([]ProjectGetOutput, error) {
	// Implement the logic to get all projects
	projects, err := p.projectRepo.GetProject(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]ProjectGetOutput, 0, len(projects))
	for _, project := range projects {
		response = append(response, ProjectGetOutput{
			Id:   project.ID,
			Name: project.Name,
		})
	}

	return response, nil
}

func NewProjectService(projectRepo repo.Project) *ProjectService {
	return &ProjectService{projectRepo: projectRepo}
}
