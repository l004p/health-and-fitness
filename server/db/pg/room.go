package pg

import(
	"context"
	"server/core"
	"strconv"
)

func (r *repoService) GetAllRooms(ctx context.Context) ([]core.Room, error){
	var result []core.Room
	rooms, err := r.getAllRooms(ctx)
	if err != nil {
		return nil, err

	}
	for _, room := range rooms {
		//fmt.Println(piece.EquipmentID)
		current := core.Room{
			RoomID: strconv.Itoa(int(room.RoomID)),
			Capacity: room.Capacity,
			Description: room.RoomDescription,
		}
		//fmt.Println("here")
		result = append(result, current)
	}
	return result, nil
}

//TODO: implement room availability checking for event registration
