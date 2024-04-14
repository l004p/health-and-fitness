package core

import (
	"context"
)

type EquipmentStatus string

const (
	EQUIPMENTACTIVE EquipmentStatus= "ACTIVE"
	MAINTENANCE EquipmentStatus = "MAINTENANCE"
	RETIRED EquipmentStatus = "RETIRED"
)

type EquipmentType string

const (
	BIKE EquipmentType= "BIKE"
	TREADMILL EquipmentType = "TREADMILL"
	STAIRMASTER EquipmentType = "STAIRMASTER"
	ROWINGMACHINE EquipmentType = "ROWING MACHINE"
	WEIGHT EquipmentType = "WEIGHT"
)

type Equipment struct {
	EquipmentID string
	EquipmentType EquipmentType
	EquipmentStatus EquipmentStatus
	Room Room
}

type EquipmentRepository interface {
	AllEquipment(ctx context.Context) ([]Equipment, error)
	EquipmentByID(ctx context.Context, id string) (Equipment, error)
	EquipmentByStatus(ctx context.Context, equipmentStatus EquipmentStatus) ([]Equipment, error)
	EquipmentByType(ctx context.Context, equipmentType EquipmentType) ([]Equipment, error)
	EquipmentByStatusAndType(ctx context.Context, equipmentType EquipmentType, equipmentStatus EquipmentStatus) ([]Equipment, error)
	MoveEquipment(ctx context.Context, equipmentID string, roomID string) (error)
	ActivateStatus(ctx context.Context, equipmentID string, roomID string) (Equipment, error)
	DeactivateStatus(ctx context.Context, equipmentID string, equipmentStatus EquipmentStatus) (Equipment, error)
}