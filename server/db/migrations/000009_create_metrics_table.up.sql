CREATE TABLE IF NOT EXISTS metrics (
    metric_id SERIAL PRIMARY KEY,
    member_id INTEGER NOT NULL,
    date_recorded DATE DEFAULT CURRENT_DATE,
    unit VARCHAR(10) NOT NULL,
    metric_value INTEGER NOT NULL,
    metric_type VARCHAR(25) NOT NULL,
    CONSTRAINT fk_member_id
        FOREIGN KEY (member_id)
            REFERENCES users(user_id)
);