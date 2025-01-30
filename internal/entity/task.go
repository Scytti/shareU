package entity

// Task represents the task entity.
type Task struct {
	ID        int    `db:"id"`
	ProjectID int    `db:"project_id"`
	Tag       string `db:"tag"`
	Command   string `db:"command"`
	Condition string `db:"condition"`
	After     string `db:"after"`
	Result    string `db:"result"`
	Status    int    `db:"status"`
	Priority  int    `db:"priority"`
}
