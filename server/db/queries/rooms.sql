-- name: addEquipmentToRoom :exec
INSERT INTO equipment_rooms (room_id, equipment_id)
VALUES ($1, $2);

-- name: removeEquipmentFromRoom :exec
DELETE FROM equipment_rooms
WHERE equipment_id=$1;

-- name: getEquipmentInRoom :many
SELECT e.equipment_id, e.equipment_type, e.equipment_status
FROM equipment e 
INNER JOIN equipment_rooms er 
ON e.equipment_id=er.equipment_id
WHERE er.room_id=$1;

-- name: getRoomsByMinimumCapacity :many
SELECT * 
FROM rooms
WHERE capacity >= $1;

-- name: getAllRooms :many
SELECT * FROM rooms;

-- name: getRoomAvailability :many

