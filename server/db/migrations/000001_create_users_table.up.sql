CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(177) NOT NULL UNIQUE,
    user_password VARCHAR(177) NOT NULL,
    first_name VARCHAR(177) NOT NULL,
    last_name VARCHAR(177) NOT NULL
);