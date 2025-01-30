-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_logs (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    task_id INT REFERENCES task (id),
    ip TEXT,
    status int REFERENCES status (id),
    result TEXT,
    log_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE (task_logs);
-- +goose StatementEnd
