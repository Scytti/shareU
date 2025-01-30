-- +goose Up
-- +goose StatementBegin
CREATE TABLE task (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    project_id int REFERENCES project (id) ON DELETE CASCADE,
    tag TEXT,
    command VARCHAR,
    condition TEXT,
    after TEXT,
    result TEXT,
    status int REFERENCES status (id) ON DELETE CASCADE DEFAULT 1,
    priority int
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE (task);
-- +goose StatementEnd
