package repository

import (
	"context"

	"github.com/edgarSucre/bochinche/domain"
)

type ChatRepository interface {
	CreateRoom(context.Context, string) (domain.Room, error)
	ListRooms(context.Context) ([]domain.Room, error)
	RegisterChatter(context.Context, domain.ChatterParams) error
	AreCredentialsValid(context.Context, domain.VerifyChatterParams) error
}
