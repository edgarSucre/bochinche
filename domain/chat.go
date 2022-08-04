package domain

import (
	"context"
	"time"
)

type Room struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type RoomParams struct {
	Name string
}

type ChatterParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Chatter struct {
	ChatterParams
	CreatedAt time.Time `json:"createdAt"`
}

type Service interface {
	CreateRoom(context.Context, RoomParams) error
	ListRooms(context.Context) ([]Room, error)
	RegisterChatter(context.Context, ChatterParams) error
	CreateSession(context.Context) error
}
