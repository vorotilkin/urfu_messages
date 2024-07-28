package services

import (
	"context"
	"messages/domain/models"
)

type UpdateMessageByIDRepository interface {
	UpdateByID(ctx context.Context, messageID int32, message string) (models.Message, error)
}

type UpdateMessageByID struct {
	repo UpdateMessageByIDRepository
}

func (m *UpdateMessageByID) UpdateMessageByID(ctx context.Context, messageID int32, message string) (models.Message, error) {
	return m.repo.UpdateByID(ctx, messageID, message)
}

func NewUpdateMessageByID(repo UpdateMessageByIDRepository) *UpdateMessageByID {
	return &UpdateMessageByID{
		repo: repo,
	}
}
