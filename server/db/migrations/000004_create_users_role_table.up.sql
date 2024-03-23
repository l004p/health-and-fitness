CREATE TABLE IF NOT EXISTS user_roles (
    user_roles_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES users(user_id),
    CONSTRAINT fk_role_id
        FOREIGN KEY (role_id)
            REFERENCES roles(role_id),
    UNIQUE (user_id, role_id)
);