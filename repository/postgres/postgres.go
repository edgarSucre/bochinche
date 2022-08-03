package postgres

import (
	"context"

	"github.com/edgarSucre/bochinche/domain"
)

type PostgresRepository struct {
	q *Queries
}

func (r *PostgresRepository) CreateRoom(ctx context.Context, name string) (domain.Room, error) {

	room, err := r.q.CreateRoom(ctx, name)
	if err != nil {
		return domain.Room{}, err
	}

	return domain.Room{Name: room.Name}, nil
}

func (r *PostgresRepository) ListRooms(ctx context.Context) ([]domain.Room, error) {
	rooms := make([]domain.Room, 0)
	result, err := r.q.ListRooms(ctx)
	if err != nil {
		return rooms, err
	}

	for _, v := range result {
		rooms = append(rooms, domain.Room{Name: v.Name})
	}

	return rooms, nil
}
