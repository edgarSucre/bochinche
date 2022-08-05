package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/edgarSucre/bochinche/domain"
)

type PostgresRepository struct {
	q *Queries
}

func NewRepository(conn *sql.DB) PostgresRepository {

	return PostgresRepository{New(conn)}
}

func (r *PostgresRepository) CreateRoom(ctx context.Context, name string) (domain.Room, error) {

	room, err := r.q.CreateRoom(ctx, name)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Room{}, domain.ErrRoomConflict
		}
		return domain.Room{}, domain.ErrInternalServerError
	}

	return domain.Room{Name: room.Name}, nil
}

func (r *PostgresRepository) ListRooms(ctx context.Context) ([]domain.Room, error) {
	rooms := make([]domain.Room, 0)
	result, err := r.q.ListRooms(ctx)
	if err != nil {
		return rooms, domain.ErrInternalServerError
	}

	for _, v := range result {
		rooms = append(rooms, domain.Room{Name: v.Name})
	}

	return rooms, nil
}

func (r *PostgresRepository) RegisterChatter(ctx context.Context, params domain.ChatterParams) error {
	pass, err := hashPassword(params.Password)
	if err != nil {
		return domain.ErrInternalServerError
	}

	err = r.q.RegisterChatter(ctx, RegisterChatterParams{
		Username: params.UserName,
		Password: pass,
		Email:    params.Email,
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.ErrChatterConflict
		}
		return domain.ErrInternalServerError
	}

	return nil
}

func (r *PostgresRepository) AreCredentialsValid(ctx context.Context, params domain.VerifyChatterParams) error {
	chatter, err := r.q.VerifyChatter(ctx, params.UserName)

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return domain.ErrNotFound
		}
		return domain.ErrBadParamInput
	}

	if !isPasswordValid(params.Password, chatter.Password) {
		return fmt.Errorf("%s", "Invalid UnserName or Password")
	}

	return nil
}
