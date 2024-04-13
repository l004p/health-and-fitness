CREATE TABLE IF NOT EXISTS goals (
    goal_id SERIAL PRIMARY KEY,
    member_id INTEGER NOT NULL,
    goal_date DATE DEFAULT CURRENT_DATE + interval '6 months',
    title VARCHAR(25) NOT NULL,
    unit VARCHAR(10) NOT NULL,
    goal_value INTEGER NOT NULL,
    UNIQUE(member_id, goal_date, title, unit, goal_value),
    CONSTRAINT fk_member_id
        FOREIGN KEY (member_id)
            REFERENCES users(user_id)
);