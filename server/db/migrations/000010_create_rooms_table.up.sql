CREATE TABLE IF NOT EXISTS rooms (
    room_id SERIAL PRIMARY KEY,
    capacity INTEGER NOT NULL,
    room_description VARCHAR(25) NOT NULL UNIQUE
);