package postgres

import (
	"context"

	"github.com/edgarSucre/bochinche/domain"
)

type ChatRepository interface {
	CreateRoom(context.Context, domain.Room) error
	ListRooms(context.Context) ([]domain.Room, error)
	// RegisterChatter(context.Context, domain.Chatter) error
	// CreateSession(context.Context) error
}
