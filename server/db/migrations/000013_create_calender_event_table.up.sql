CREATE TABLE IF NOT EXISTS calender_events (
    event_id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL,
    event_type VARCHAR(25) NOT NULL,
    max_attendees INTEGER NOT NULL,
    event_status VARCHAR(25) NOT NULL,
    event_title VARCHAR(25) NOT NULL,
    start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP + interval '2 hours'
);