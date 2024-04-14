package core

import "context"

type Room struct {
	RoomID string
	Capacity int32
	Description string
}

type RoomRepository interface {
	GetAllRooms(ctx context.Context) ([]Room, error)
	
}