// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: rooms.sql

package pg

import (
	"context"
)

const addEquipmentToRoom = `-- name: addEquipmentToRoom :exec
INSERT INTO equipment_rooms (room_id, equipment_id)
VALUES ($1, $2)
`

type addEquipmentToRoomParams struct {
	RoomID      int32
	EquipmentID int32
}

func (q *Queries) addEquipmentToRoom(ctx context.Context, arg addEquipmentToRoomParams) error {
	_, err := q.db.Exec(ctx, addEquipmentToRoom, arg.RoomID, arg.EquipmentID)
	return err
}

const getAllRooms = `-- name: getAllRooms :many
SELECT room_id, capacity, room_description FROM rooms
`

func (q *Queries) getAllRooms(ctx context.Context) ([]Room, error) {
	rows, err := q.db.Query(ctx, getAllRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Room
	for rows.Next() {
		var i Room
		if err := rows.Scan(&i.RoomID, &i.Capacity, &i.RoomDescription); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEquipmentInRoom = `-- name: getEquipmentInRoom :many
SELECT e.equipment_id, e.equipment_type, e.equipment_status
FROM equipment e 
INNER JOIN equipment_rooms er 
ON e.equipment_id=er.equipment_id
WHERE er.room_id=$1
`

func (q *Queries) getEquipmentInRoom(ctx context.Context, roomID int32) ([]Equipment, error) {
	rows, err := q.db.Query(ctx, getEquipmentInRoom, roomID)
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

const getRoomsByMinimumCapacity = `-- name: getRoomsByMinimumCapacity :many
SELECT room_id, capacity, room_description 
FROM rooms
WHERE capacity >= $1
`

func (q *Queries) getRoomsByMinimumCapacity(ctx context.Context, capacity int32) ([]Room, error) {
	rows, err := q.db.Query(ctx, getRoomsByMinimumCapacity, capacity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Room
	for rows.Next() {
		var i Room
		if err := rows.Scan(&i.RoomID, &i.Capacity, &i.RoomDescription); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeEquipmentFromRoom = `-- name: removeEquipmentFromRoom :exec
DELETE FROM equipment_rooms
WHERE equipment_id=$1
`

func (q *Queries) removeEquipmentFromRoom(ctx context.Context, equipmentID int32) error {
	_, err := q.db.Exec(ctx, removeEquipmentFromRoom, equipmentID)
	return err
}
