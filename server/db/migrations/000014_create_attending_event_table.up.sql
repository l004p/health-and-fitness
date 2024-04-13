CREATE TABLE IF NOT EXISTS attending_event (
    member_id INTEGER NOT NULL,
    event_id INTEGER NOT NULL,
    PRIMARY KEY (member_id, event_id),
    CONSTRAINT fk_member_id
        FOREIGN KEY (member_id)
            REFERENCES users(user_id),
    CONSTRAINT fk_event_id
        FOREIGN KEY (event_id)
            REFERENCES calender_events(event_id)
);