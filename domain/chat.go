package domain

import "context"

type Room struct {
	Name string
}

type Chatter struct {
	Name     string
	Password string
}

type Service interface {
	CreateRoom(context.Context, Room) error
	ListRooms(context.Context) error
	RegisterChatter(context.Context, Chatter) error
	CreateSession(context.Context) error
}
