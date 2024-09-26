package postgres

import (
	"database/sql"
	"fmt"

	"github.com/axadjonovsardorbek/MiniTwitter/twit_service/config"
	"github.com/axadjonovsardorbek/MiniTwitter/twit_service/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db      *sql.DB
	FollowS storage.FollowI
	LikeS   storage.LikeI
	TwitS   storage.TwitI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	follow := NewFollowRepo(db)
	like := NewLikeRepo(db)
	twit := NewTwitRepo(db)

	return &Storage{
		Db:      db,
		FollowS: follow,
		LikeS:   like,
		TwitS:   twit,
	}, nil
}
