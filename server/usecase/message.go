package usecase

import (
	"github.com/jmoiron/sqlx"
)

type Message struct {
	db *sqlx.DB
}

func NewMessage(db *sqlx.DB) *Message {
	return &Message{
		db: db,
	}
}

