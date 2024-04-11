CREATE TABLE IF NOT EXISTS trainer_interests (
    trainer_id INTEGER NOT NULL,
    interest VARCHAR(25) NOT NULL UNIQUE,
    CONSTRAINT fk_trainer_id
        FOREIGN KEY (trainer_id)
            REFERENCES users(user_id)
);