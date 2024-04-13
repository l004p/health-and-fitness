// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: equipment.sql

package pg

import (
	"context"
)

const addEquipment = `-- name: addEquipment :exec
INSERT INTO equipment (equipment_type, equipment_status)
VALUES ($1, 'ACTIVE')
`

func (q *Queries) addEquipment(ctx context.Context, equipmentType string) error {
	_, err := q.db.Exec(ctx, addEquipment, equipmentType)
	return err
}

const getEquipment = `-- name: getEquipment :many
SELECT equipment_id, equipment_type, equipment_status FROM equipment
`

func (q *Queries) getEquipment(ctx context.Context) ([]Equipment, error) {
	rows, err := q.db.Query(ctx, getEquipment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Equipment
	for rows.Next() {
		var i Equipment
		if err := rows.Scan(&i.EquipmentID, &i.EquipmentType, &i.EquipmentStatus); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEquipmentByID = `-- name: getEquipmentByID :one
SELECT equipment_id, equipment_type, equipment_status FROM equipment
WHERE equipment_id=$1
`

func (q *Queries) getEquipmentByID(ctx context.Context, equipmentID int32) (Equipment, error) {
	row := q.db.QueryRow(ctx, getEquipmentByID, equipmentID)
	var i Equipment
	err := row.Scan(&i.EquipmentID, &i.EquipmentType, &i.EquipmentStatus)
	return i, err
}

const getEquipmentByStatus = `-- name: getEquipmentByStatus :many
SELECT equipment_id, equipment_type, equipment_status FROM equipment
WHERE equipment_status=$1
`

func (q *Queries) getEquipmentByStatus(ctx context.Context, equipmentStatus string) ([]Equipment, error) {
	rows, err := q.db.Query(ctx, getEquipmentByStatus, equipmentStatus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Equipment
	for rows.Next() {
		var i Equipment
		if err := rows.Scan(&i.EquipmentID, &i.EquipmentType, &i.EquipmentStatus); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEquipmentByType = `-- name: getEquipmentByType :many
SELECT equipment_id, equipment_type, equipment_status FROM equipment
WHERE equipment_type=$1
`

func (q *Queries) getEquipmentByType(ctx context.Context, equipmentType string) ([]Equipment, error) {
	rows, err := q.db.Query(ctx, getEquipmentByType, equipmentType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Equipment
	for rows.Next() {
		var i Equipment
		if err := rows.Scan(&i.EquipmentID, &i.EquipmentType, &i.EquipmentStatus); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStatus = `-- name: updateStatus :exec
UPDATE equipment
SET equipment_status=$2
WHERE equipment_id=$1
`

type updateStatusParams struct {
	EquipmentID     int32
	EquipmentStatus string
}

func (q *Queries) updateStatus(ctx context.Context, arg updateStatusParams) error {
	_, err := q.db.Exec(ctx, updateStatus, arg.EquipmentID, arg.EquipmentStatus)
	return err
}
