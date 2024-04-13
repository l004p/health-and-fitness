CREATE TABLE IF NOT EXISTS bills (
    bill_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    bill_description VARCHAR(25) NOT NULL,
    bill_status VARCHAR(25) NOT NULL,
    bill_date DATE DEFAULT CURRENT_DATE,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES users(user_id)
);