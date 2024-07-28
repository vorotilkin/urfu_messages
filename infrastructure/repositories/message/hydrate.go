package message

import (
	"messages/domain/models"
	"messages/schema/gen/model"
)

func toDomain(message model.Message) models.Message {
	return models.Message{
		ID:        message.ID,
		UserID:    message.UserID,
		Message:   message.Message,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}
