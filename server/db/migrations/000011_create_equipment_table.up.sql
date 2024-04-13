CREATE TABLE IF NOT EXISTS equipment (
    equipment_id SERIAL PRIMARY KEY,
    equipment_type VARCHAR(25) NOT NULL,
    equipment_status VARCHAR(25) NOT NULL
);