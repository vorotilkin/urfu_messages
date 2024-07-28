package services

import (
	"context"
)

type DeleteMessageByIDRepository interface {
	DeleteByID(ctx context.Context, messageID int32) (bool, error)
}

type DeleteMessageByID struct {
	repo DeleteMessageByIDRepository
}

func (m *DeleteMessageByID) DeleteMessageByID(ctx context.Context, messageID int32) (bool, error) {
	return m.repo.DeleteByID(ctx, messageID)
}

func NewDeleteMessageByID(repo DeleteMessageByIDRepository) *DeleteMessageByID {
	return &DeleteMessageByID{
		repo: repo,
	}
}
