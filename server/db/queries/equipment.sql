-- name: addEquipment :exec
INSERT INTO equipment (equipment_type, equipment_status)
VALUES ($1, 'ACTIVE');

-- name: getEquipmentByStatus :many
SELECT * FROM equipment
WHERE equipment_status=$1;

-- name: getEquipmentByType :many
SELECT * FROM equipment
WHERE equipment_type=$1;

-- name: getEquipmentByID :one
SELECT * FROM equipment
WHERE equipment_id=$1;

-- name: updateStatus :exec
UPDATE equipment
SET equipment_status=$2
WHERE equipment_id=$1;

