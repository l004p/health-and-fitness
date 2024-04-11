CREATE TABLE IF NOT EXISTS member_trained (
    trainer_id INTEGER NOT NULL,
    member_id INTEGER NOT NULL,
    PRIMARY KEY (trainer_id, member_id),
    CONSTRAINT fk_trainer_id
        FOREIGN KEY (trainer_id)
            REFERENCES users(user_id),
    CONSTRAINT fk_member_id
        FOREIGN KEY (member_id)
            REFERENCES users(user_id)
);