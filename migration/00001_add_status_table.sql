-- +goose Up
-- +goose StatementBegin
CREATE TABLE status (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR
);

INSERT INTO status(name)
    VALUES (to_do),
           (doing),
           (done),
           (error),
           (expired);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE(status)
-- +goose StatementEnd
