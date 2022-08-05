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
	ChatterParams
	CreatedAt time.Time `json:"createdAt"`
}

type VerifyChatterParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Service interface {
	CreateRoom(context.Context, RoomParams) error
	ListRooms(context.Context) ([]Room, error)
	RegisterChatter(context.Context, ChatterParams) error
	IsPasswordValid(context.Context, VerifyChatterParams) error
}

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("chatter not found")
	ErrChatterConflict     = errors.New("chatter already exists")
	ErrRoomConflict        = errors.New("room alredy exists")
	ErrBadParamInput       = errors.New("given param is not valid")
)
