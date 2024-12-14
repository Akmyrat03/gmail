package repository

import (
	"context"
	"mail-service/models"

	"github.com/jmoiron/sqlx"
)

type MessageRepository struct {
	DB *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (r *MessageRepository) SaveMessage(ctx context.Context, message models.ContactMessage) error {
	query := `INSERT INTO contact_messages (name, email, message) VALUES ($1, $2, $3)`
	_, err := r.DB.ExecContext(ctx, query, message.Name, message.Email, message.Message)
	return err
}
