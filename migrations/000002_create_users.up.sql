CREATE TABLE
    IF NOT EXISTS users (
        id serial PRIMARY KEY,
        tg_id INTEGER NULL UNIQUE,
        user_role VARCHAR(255) NULL
    );