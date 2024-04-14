package equipmentservice

import (
	"server/core"
	"context"
)

func EquipmentByID(r core.EquipmentRepository, ctx context.Context, id string) (core.Equipment, error) {
	e, err := r.EquipmentByID(ctx, id)
	if err != nil {
		return core.Equipment{}, err
	}
	return e, nil
}

func Equipment(r core.EquipmentRepository, ctx context.Context) ([]core.Equipment, error){
	//panic("not implemented")
	e, err := r.AllEquipment(ctx)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func ByType(r core.EquipmentRepository, ctx context.Context, equipmentType core.EquipmentType) ([]core.Equipment, error) {
	//panic("not implemented")
	e, err := r.EquipmentByType(ctx, equipmentType)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func ByStatus(r core.EquipmentRepository, ctx context.Context, equipmentStatus core.EquipmentStatus) ([]core.Equipment, error) {
	//panic("not implemented")
	e, err := r.EquipmentByStatus(ctx, equipmentStatus)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func ByStatusAndType(r core.EquipmentRepository, ctx context.Context, equipmentStatus core.EquipmentStatus, equipmentType core.EquipmentType) ([]core.Equipment, error) {
	//panic("not implemented")
	e, err := r.EquipmentByStatusAndType(ctx, equipmentType, equipmentStatus)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func ActivateStatus(r core.EquipmentRepository, ctx context.Context, equipmentID string, roomID string) (core.Equipment, error) {
	//panic("not implemented")
	e, err := r.ActivateStatus(ctx, equipmentID, roomID)
	if err != nil {
		return core.Equipment{}, err
	}
	return e, nil
}

func DeactivateStatus(r core.EquipmentRepository, ctx context.Context, equipmentID string, equipmentStatus core.EquipmentStatus) (core.Equipment, error) {
	//panic("not implemented")
	e, err := r.DeactivateStatus(ctx, equipmentID, equipmentStatus)
	if err != nil {
		return core.Equipment{}, err
	}
	return e, nil
}


func AssignRoom(r core.EquipmentRepository, ctx context.Context, roomID string, equipmentID string) (error) {
	//panic("not implemented")
	err := r.MoveEquipment(ctx, equipmentID, roomID)
	if err != nil {
		return err
	}
	return nil
}