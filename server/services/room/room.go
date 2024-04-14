package roomservice

import (
	"context"
	"server/core"
)

func GetAllRooms(r core.RoomRepository, ctx context.Context) ([]core.Room, error) {
	rooms, err := r.GetAllRooms(ctx)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}