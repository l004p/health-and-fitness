CREATE TABLE IF NOT EXISTS memberships (
    membership_id SERIAL PRIMARY KEY,
    member_id INTEGER NOT NULL,
    title VARCHAR(25) NOT NULL,
    membership_status VARCHAR(25) NOT NULL,
    membership_start DATE DEFAULT CURRENT_DATE,
    membership_end DATE DEFAULT CURRENT_DATE + interval '1 year',
    UNIQUE(member_id, title),
    CONSTRAINT fk_member_id
        FOREIGN KEY (member_id)
            REFERENCES users(user_id)
);