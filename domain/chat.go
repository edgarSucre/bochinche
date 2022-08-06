package domain

import (
	"context"
	"errors"
	"time"
)

//TODO: add validation
type Room struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type RoomParams struct {
	Name string `json:"name"`
}

type ChatterParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Chatter struct {
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
}

type VerifyChatterParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type ChatParam struct {
	Room    string `json:"room"`
	Author  string `json:"author"`
	Message string `json:"message"`
}

type Chat struct {
	ID        int64     `json:"id"`
	Room      string    `json:"room"`
	Author    string    `json:"author"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

type ChatRepository interface {
	CreateRoom(context.Context, string) (Room, error)
	ListRooms(context.Context) ([]Room, error)
	RegisterChatter(context.Context, ChatterParams) error
	AreCredentialsValid(context.Context, VerifyChatterParams) error
	GetChatter(context.Context, string) (Chatter, error)
	SaveChat(context.Context, ChatParam) error
	ListChats(context.Context, string) ([]Chat, error)
}

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("chatter not found")
	ErrChatterConflict     = errors.New("chatter already exists")
	ErrRoomConflict        = errors.New("room alredy exists")
	ErrBadParamInput       = errors.New("given param is not valid")
)
