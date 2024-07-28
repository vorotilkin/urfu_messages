package services

import (
	"context"
	"messages/domain/models"
)

type FetchMessagesByUserIDRepository interface {
	MessagesByUserID(ctx context.Context, messageID int32) ([]models.Message, error)
}

type FetchMessagesByUserID struct {
	repo FetchMessagesByUserIDRepository
}

func (m *FetchMessagesByUserID) FetchMessagesByUserID(ctx context.Context, userID int32) ([]models.Message, error) {
	return m.repo.MessagesByUserID(ctx, userID)
}

func NewFetchMessagesByUserID(repo FetchMessagesByUserIDRepository) *FetchMessagesByUserID {
	return &FetchMessagesByUserID{
		repo: repo,
	}
}
