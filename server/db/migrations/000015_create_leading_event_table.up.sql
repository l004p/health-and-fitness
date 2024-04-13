CREATE TABLE IF NOT EXISTS leading_event (
    trainer_id INTEGER NOT NULL,
    event_id INTEGER NOT NULL,
    PRIMARY KEY (trainer_id, event_id),
    CONSTRAINT fk_trainer_id
        FOREIGN KEY (trainer_id)
            REFERENCES users(user_id),
    CONSTRAINT fk_event_id
        FOREIGN KEY (event_id)
            REFERENCES calender_events(event_id)
);