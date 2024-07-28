package message

import (
	"context"
	"github.com/go-jet/jet/v2/postgres"
	"messages/domain/models"
	"messages/pkg/database"
	"messages/schema/gen/model"
	"messages/schema/gen/table"
	"time"
)

type Repository struct {
	conn *database.Database
}

func (r *Repository) CreateMessage(ctx context.Context, userID int32, message string) (bool, error) {
	query, args := table.Message.
		INSERT(table.Message.UserID, table.Message.Message).
		MODEL(model.Message{UserID: userID, Message: message}).
		Sql()

	tag, err := r.conn.Exec(ctx, query, args...)
	if err != nil {
		return false, err
	}

	return tag.RowsAffected() > 0, nil
}

func (r *Repository) MessageByID(ctx context.Context, id int32) (models.Message, error) {
	query, args := table.Message.
		SELECT(table.Message.AllColumns).
		WHERE(table.Message.ID.
			EQ(postgres.Int32(id))).
		Sql()

	row := r.conn.QueryRow(ctx, query, args...)
	message := model.Message{}

	err := row.Scan(
		&message.ID,
		&message.UserID,
		&message.Message,
		&message.CreatedAt,
		&message.UpdatedAt,
	)
	if err != nil {

		return models.Message{}, err
	}

	return toDomain(message), nil
}

func (r *Repository) DeleteByID(ctx context.Context, id int32) (bool, error) {
	query, args := table.Message.
		DELETE().
		WHERE(table.Message.ID.EQ(postgres.Int32(id))).
		Sql()

	tag, err := r.conn.Exec(ctx, query, args...)
	if err != nil {
		return false, err
	}

	return tag.RowsAffected() > 0, nil
}

func (r *Repository) MessagesByUserID(ctx context.Context, userID int32) ([]models.Message, error) {
	query, args := table.Message.
		SELECT(table.Message.AllColumns).
		WHERE(table.Message.UserID.
			EQ(postgres.Int32(userID))).
		Sql()

	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := make([]models.Message, 0)
	for rows.Next() {
		message := model.Message{}

		err := rows.Scan(
			&message.ID,
			&message.UserID,
			&message.Message,
			&message.CreatedAt,
			&message.UpdatedAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, toDomain(message))
	}

	return messages, nil
}

func (r *Repository) UpdateByID(ctx context.Context, id int32, message string) (models.Message, error) {
	query, args := table.Message.
		UPDATE(table.Message.Message, table.Message.UpdatedAt).
		MODEL(model.Message{Message: message, UpdatedAt: time.Now()}).
		WHERE(table.Message.ID.
			EQ(postgres.Int32(id))).
		RETURNING(table.Message.AllColumns).
		Sql()

	row := r.conn.QueryRow(ctx, query, args...)
	updatedMessage := model.Message{}

	err := row.Scan(
		&updatedMessage.ID,
		&updatedMessage.UserID,
		&updatedMessage.Message,
		&updatedMessage.CreatedAt,
		&updatedMessage.UpdatedAt,
	)
	if err != nil {
		return models.Message{}, err
	}

	return toDomain(updatedMessage), nil
}

func NewRepository(conn *database.Database) *Repository {
	return &Repository{conn: conn}
}
