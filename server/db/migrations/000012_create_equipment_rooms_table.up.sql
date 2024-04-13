CREATE TABLE IF NOT EXISTS equipment_rooms (
    room_id INTEGER NOT NULL,
    equipment_id INTEGER NOT NULL,
    PRIMARY KEY (room_id, equipment_id),
    UNIQUE(equipment_id),
    CONSTRAINT fk_room_id
        FOREIGN KEY (room_id)
            REFERENCES rooms(room_id),
    CONSTRAINT fk_equipment_id
        FOREIGN KEY (equipment_id)
            REFERENCES equipment(equipment_id)
);