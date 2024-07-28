package services

import "context"

type CreateMessageRepository interface {
	CreateMessage(ctx context.Context, userID int32, message string) (bool, error)
}

type CreateMessage struct {
	repo CreateMessageRepository
}

func (m *CreateMessage) CreateMessage(ctx context.Context, userID int32, message string) (bool, error) {
	return m.repo.CreateMessage(ctx, userID, message)
}

func NewCreateMessage(repo CreateMessageRepository) *CreateMessage {
	return &CreateMessage{
		repo: repo,
	}
}
