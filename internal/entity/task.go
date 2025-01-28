package entity

// Task represents the task entity.
type Task struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Tag       string `json:"tag"`
	Command   string `json:"command"`
	Status    int    `json:"status"`
	Priority  int    `json:"priority"`
}
