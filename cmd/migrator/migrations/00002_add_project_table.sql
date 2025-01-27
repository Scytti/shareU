-- +goose Up
-- +goose StatementBegin
CREATE TABLE project (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL,
    env JSONB,
    expect_res BOOLEAN DEFAULT FALSE,
    use_docker BOOLEAN DEFAULT FALSE,
    link_docker VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE(project)
-- +goose StatementEnd
