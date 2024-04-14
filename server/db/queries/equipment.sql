-- name: addEquipment :exec
INSERT INTO equipment (equipment_type, equipment_status)
VALUES ($1, 'ACTIVE');

-- name: getEquipmentByStatus :many
SELECT * FROM equipment
WHERE equipment_status=$1;

-- name: getEquipmentByType :many
SELECT * FROM equipment
WHERE equipment_type=$1;

-- name: getEquipmentByTypeAndStatus :many
SELECT * FROM equipment
WHERE equipment_status=$1 AND equipment_type=$2;

-- name: getEquipment :many
SELECT * FROM equipment;

-- name: getEquipmentByID :one
SELECT * FROM equipment
WHERE equipment_id=$1;

-- name: updateStatus :one
UPDATE equipment
SET equipment_status=$2
WHERE equipment_id=$1
RETURNING equipment_id, equipment_status, equipment_type;

-- name: getRoomEquipmentIn :one
SELECT r.room_id, r.capacity, r.room_description 
FROM rooms r 
INNER JOIN equipment_rooms er 
ON r.room_id=er.room_id
WHERE er.equipment_id=$1;


