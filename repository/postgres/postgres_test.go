package postgres_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/edgarSucre/bochinche/domain"
	"github.com/stretchr/testify/require"
)

func TestCreateRoom(t *testing.T) {
	name := gofakeit.StreetName()
	room, err := repo.CreateRoom(context.Background(), name)

	require.NoError(t, err)
	require.Equal(t, name, room.Name)
	require.NotEmpty(t, room.CreatedAt)
}

func TestListRooms(t *testing.T) {
	name := gofakeit.StreetName()
	repo.CreateRoom(context.Background(), name)

	rooms, err := repo.ListRooms(context.Background())
	require.NoError(t, err)

	last := rooms[len(rooms)-1]
	require.Equal(t, name, last.Name)
	require.NotEmpty(t, last.CreatedAt)
}

func TestRegisterChatter(t *testing.T) {
	params := domain.ChatterParams{
		UserName: gofakeit.Username(),
		Password: gofakeit.ProgrammingLanguage(),
		Email:    gofakeit.Email(),
	}

	err := repo.RegisterChatter(context.Background(), params)
	require.NoError(t, err)

	// try to duplicate chatter
	err = repo.RegisterChatter(context.Background(), params)
	require.ErrorIs(t, err, domain.ErrChatterConflict)

	err = repo.AreCredentialsValid(context.Background(), domain.VerifyChatterParams{
		UserName: params.UserName,
		Password: params.Password,
	})

	require.NoError(t, err)

	chatter, err := repo.GetChatter(context.Background(), params.UserName)
	require.NoError(t, err)
	require.Equal(t, chatter.UserName, params.UserName)
}

func TestSaveChat(t *testing.T) {
	chatter := domain.ChatterParams{
		UserName: gofakeit.Username(),
		Password: gofakeit.ProgrammingLanguage(),
		Email:    gofakeit.Email(),
	}

	repo.RegisterChatter(context.Background(), chatter)
	room, _ := repo.CreateRoom(context.Background(), gofakeit.Name())

	params := domain.ChatParam{
		Room:    room.Name,
		Author:  chatter.UserName,
		Message: gofakeit.BeerName(),
	}

	err := repo.SaveChat(context.Background(), params)
	require.NoError(t, err)

	chats, err := repo.ListChats(context.Background(), room.Name)
	require.NoError(t, err)

	last := chats[len(chats)-1]
	require.Equal(t, last.Author, chatter.UserName)
	require.Equal(t, last.Message, params.Message)
	require.Equal(t, last.Room, room.Name)
}
