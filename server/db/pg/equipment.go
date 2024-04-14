package pg

import (
	"context"
	"fmt"
	"server/core"
	"strconv"
)

func (r *repoService) AllEquipment(ctx context.Context) ([]core.Equipment, error) {
	var equipment []core.Equipment

	e, err := r.getEquipment(ctx)
	if err != nil {
		return nil, err
	}
	if e == nil {
		fmt.Println("EMPTY")
	}
	//fmt.Println("Here")
	for _, piece := range e {
		//fmt.Println(piece.EquipmentID)
		if piece.EquipmentStatus == string(core.ACTIVE) {
			r, err := r.getRoomEquipmentIn(ctx, piece.EquipmentID)
			if err != nil {
				return nil, err
			}
			current := core.Equipment{
				EquipmentID: strconv.Itoa(int(piece.EquipmentID)),
				EquipmentType: core.EquipmentType(piece.EquipmentType),
				EquipmentStatus: core.EquipmentStatus(piece.EquipmentStatus),
				Room: core.Room{
					RoomID: strconv.Itoa(int(r.RoomID)),
					Capacity: r.Capacity,
					Description: r.RoomDescription,
				},
			}
			equipment = append(equipment, current)
			continue;
		}
		current := core.Equipment{
			EquipmentID: strconv.Itoa(int(piece.EquipmentID)),
			EquipmentType: core.EquipmentType(piece.EquipmentType),
			EquipmentStatus: core.EquipmentStatus(piece.EquipmentStatus),
			Room: core.Room{},
		}
		//fmt.Println("here")
		equipment = append(equipment, current)
	}
	return equipment, nil
}

func (r *repoService) EquipmentByID(ctx context.Context, id string) (core.Equipment, error) {
	intID, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		return core.Equipment{}, err
	}

	equipment, err := r.getEquipmentByID(ctx, int32(intID))
	if err != nil {
		return core.Equipment{}, err
	}
	
	return core.Equipment{
		EquipmentID: id,
		EquipmentType: core.EquipmentType(equipment.EquipmentType),
		EquipmentStatus: core.EquipmentStatus(equipment.EquipmentStatus),
	}, nil
}

func (r *repoService) EquipmentByStatus(ctx context.Context, equipmentStatus core.EquipmentStatus) ([]core.Equipment, error) {
	//panic("not implemented")
	var equipment []core.Equipment

	e, err := r.getEquipmentByStatus(ctx, string(equipmentStatus))
	if err != nil {
		return nil, err
	}
	if e == nil {
		fmt.Println("EMPTY")
	}
	//fmt.Println("Here")
	for _, piece := range e {
		//fmt.Println(piece.EquipmentID)
		if piece.EquipmentStatus == string(core.ACTIVE) {
			r, err := r.getRoomEquipmentIn(ctx, piece.EquipmentID)
			if err != nil {
				return nil, err
			}
			current := core.Equipment{
				EquipmentID: strconv.Itoa(int(piece.EquipmentID)),
				EquipmentType: core.EquipmentType(piece.EquipmentType),
				EquipmentStatus: core.EquipmentStatus(piece.EquipmentStatus),
				Room: core.Room{
					RoomID: strconv.Itoa(int(r.RoomID)),
					Capacity: r.Capacity,
					Description: r.RoomDescription,
				},
			}
			equipment = append(equipment, current)
			continue;
		}
		current := core.Equipment{
			EquipmentID: strconv.Itoa(int(piece.EquipmentID)),
			EquipmentType: core.EquipmentType(piece.EquipmentType),
			EquipmentStatus: core.EquipmentStatus(piece.EquipmentStatus),
			Room: core.Room{},
		}
		//fmt.Println("here")
		equipment = append(equipment, current)
	}
	return equipment, nil
}

func (r *repoService) EquipmentByType(ctx context.Context, equipmentType core.EquipmentType) ([]core.Equipment, error) {
	//panic("not implemented")
	var equipment []core.Equipment

	e, err := r.getEquipmentByType(ctx, string(equipmentType))
	if err != nil {
		return nil, err
	}
	if e == nil {
		fmt.Println("EMPTY")
	}
	//fmt.Println("Here")
	for _, piece := range e {
		//fmt.Println(piece.EquipmentID)
		if piece.EquipmentStatus == string(core.ACTIVE) {
			r, err := r.getRoomEquipmentIn(ctx, piece.EquipmentID)
			if err != nil {
				return nil, err
			}
			current := core.Equipment{
				EquipmentID: strconv.Itoa(int(piece.EquipmentID)),
				EquipmentType: core.EquipmentType(piece.EquipmentType),
				EquipmentStatus: core.EquipmentStatus(piece.EquipmentStatus),
				Room: core.Room{
					RoomID: strconv.Itoa(int(r.RoomID)),
					Capacity: r.Capacity,
					Description: r.RoomDescription,
				},
			}
			equipment = append(equipment, current)
			continue;
		}
		current := core.Equipment{
			EquipmentID: strconv.Itoa(int(piece.EquipmentID)),
			EquipmentType: core.EquipmentType(piece.EquipmentType),
			EquipmentStatus: core.EquipmentStatus(piece.EquipmentStatus),
			Room: core.Room{},
		}
		//fmt.Println("here")
		equipment = append(equipment, current)
	}
	return equipment, nil
}

func (r *repoService) EquipmentByStatusAndType(ctx context.Context, equipmentType core.EquipmentType, equipmentStatus core.EquipmentStatus) ([]core.Equipment, error) {
	//panic("not implemented")
	var equipment []core.Equipment

	params := getEquipmentByTypeAndStatusParams{
		EquipmentStatus: string(equipmentStatus),
		EquipmentType: string(equipmentType),
	}
	e, err := r.getEquipmentByTypeAndStatus(ctx, params)
	if err != nil {
		return nil, err
	}
	if e == nil {
		fmt.Println("EMPTY")
	}
	//fmt.Println("Here")
	for _, piece := range e {
		//fmt.Println(piece.EquipmentID)
		if piece.EquipmentStatus == string(core.ACTIVE) {
			r, err := r.getRoomEquipmentIn(ctx, piece.EquipmentID)
			if err != nil {
				return nil, err
			}
			current := core.Equipment{
				EquipmentID: strconv.Itoa(int(piece.EquipmentID)),
				EquipmentType: core.EquipmentType(piece.EquipmentType),
				EquipmentStatus: core.EquipmentStatus(piece.EquipmentStatus),
				Room: core.Room{
					RoomID: strconv.Itoa(int(r.RoomID)),
					Capacity: r.Capacity,
					Description: r.RoomDescription,
				},
			}
			equipment = append(equipment, current)
			continue;
		}
		current := core.Equipment{
			EquipmentID: strconv.Itoa(int(piece.EquipmentID)),
			EquipmentType: core.EquipmentType(piece.EquipmentType),
			EquipmentStatus: core.EquipmentStatus(piece.EquipmentStatus),
			Room: core.Room{},
		}
		//fmt.Println("here")
		equipment = append(equipment, current)
	}
	return equipment, nil
}

func (r *repoService) ActivateStatus(ctx context.Context, equipmentID string, roomID string) (core.Equipment, error) {
	//panic("not implemented")
	rID, err := strconv.ParseInt(roomID, 10, 32)
	if err != nil {
		return core.Equipment{}, err
	}
	eID, err := strconv.ParseInt(equipmentID, 10, 32)
	if err != nil {
		return core.Equipment{}, err
	}
	statusParams := updateStatusParams{
		EquipmentID: int32(eID),
		EquipmentStatus: "ACTIVE",
	}
	roomParams := addEquipmentToRoomParams{
		RoomID: int32(rID),
		EquipmentID: int32(eID),
	}
	//this should be wrapped in a transaction.... but shhhhhhhhh don't tell anyone
	err = r.addEquipmentToRoom(ctx, roomParams)
	if err != nil {
		return core.Equipment{}, err
	}
	
	e, err := r.updateStatus(ctx, statusParams)
	if err != nil {
		return core.Equipment{}, err
	}
	returnValue := core.Equipment{
		EquipmentID: strconv.Itoa(int(e.EquipmentID)),
		EquipmentType: core.EquipmentType(e.EquipmentType),
		EquipmentStatus: core.EquipmentStatus(e.EquipmentStatus),
	}
	return returnValue, nil
}

func (r *repoService) DeactivateStatus(ctx context.Context, equipmentID string, equipmentStatus core.EquipmentStatus) (core.Equipment, error) {
	//panic("not implemented")
	eID, err := strconv.ParseInt(equipmentID, 10, 32)
	if err != nil {
		return core.Equipment{}, err
	}
	statusParams := updateStatusParams{
		EquipmentID: int32(eID),
		EquipmentStatus: string(equipmentStatus),
	}
	//this should be wrapped in a transaction.... but shhhhhhhhh don't tell anyone
	err = r.removeEquipmentFromRoom(ctx, int32(eID))
	if err != nil {
		return core.Equipment{}, err
	}
	
	e, err := r.updateStatus(ctx, statusParams)
	if err != nil {
		return core.Equipment{}, err
	}
	returnValue := core.Equipment{
		EquipmentID: strconv.Itoa(int(e.EquipmentID)),
		EquipmentType: core.EquipmentType(e.EquipmentType),
		EquipmentStatus: core.EquipmentStatus(e.EquipmentStatus),
	}
	return returnValue, nil

}

func (r *repoService) MoveEquipment(ctx context.Context, equipmentID string, roomID string) (error){
	//panic("not implemented")
	eID, err := strconv.ParseInt(equipmentID, 10, 32)
	if err != nil {
		return err
	}
	rID, err := strconv.ParseInt(roomID, 10, 32)
	if err != nil {
		return err
	}
	e, err := r.getEquipmentByID(ctx, int32(eID))
	if err != nil {
		return err
	}
	if e.EquipmentStatus != "ACTIVE" {
		return fmt.Errorf("Equipment not available")
	}
	err = r.removeEquipmentFromRoom(ctx, int32(eID))
	if err != nil {
		return err
	}
	addParams := addEquipmentToRoomParams{
		RoomID: int32(rID),
		EquipmentID: int32(eID),
	}
	err = r.addEquipmentToRoom(ctx, addParams)
	if err != nil {
		return err
	}
	return nil
}