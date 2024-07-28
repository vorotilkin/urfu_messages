package services

import (
	"context"
	"messages/domain/models"
)

type FetchMessageByIDRepository interface {
	MessageByID(ctx context.Context, messageID int32) (models.Message, error)
}

type FetchMessageByID struct {
	repo FetchMessageByIDRepository
}

func (m *FetchMessageByID) FetchMessageByID(ctx context.Context, messageID int32) (models.Message, error) {
	return m.repo.MessageByID(ctx, messageID)
}

func NewFetchMessageByID(repo FetchMessageByIDRepository) *FetchMessageByID {
	return &FetchMessageByID{
		repo: repo,
	}
}
