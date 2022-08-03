package domain

import "context"

type Room struct {
	Name string
}

type Chatter struct {
	UserName string
	Password string
	Email    string
}

type Service interface {
	CreateRoom(context.Context, Room) error
	ListRooms(context.Context) error
	RegisterChatter(context.Context, Chatter) error
	CreateSession(context.Context) error
}
