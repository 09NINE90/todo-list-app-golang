CREATE SCHEMA todoapp;

CREATE TABLE todoapp.users
(
    id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name   VARCHAR(50) NOT NULL CHECK (char_length(first_name) BETWEEN 3 AND 50),
    last_name    VARCHAR(50) NOT NULL CHECK (char_length(last_name) BETWEEN 3 AND 50),
    phone_number VARCHAR(15) NOT NULL CHECK (
        phone_number ~ '^\+[0-9]+$'
            AND
        char_length(phone_number) BETWEEN 10 AND 15
        )
);

CREATE TABLE todoapp.tasks
(
    id           uuid PRIMARY KEY,
    title        VARCHAR(50) NOT NULL CHECK (char_length(title) BETWEEN 3 AND 50),
    description  VARCHAR(1000) CHECK (char_length(title) BETWEEN 5 AND 1000),
    completed    BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at   timestamptz NOT NULL DEFAULT NOW(),
    completed_at timestamptz,

    CHECK (
        (completed = FALSE AND completed_at IS NULL)
            OR
        (completed = TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
        ),

    author_id    uuid        NOT NULL REFERENCES todoapp.users (id)
);