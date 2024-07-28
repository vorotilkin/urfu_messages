package handlers

import (
	"context"
	"github.com/samber/lo"
	domain "messages/domain/models"
	"messages/domain/services"
	"messages/pkg/http"
	"messages/usecases/get_by_user_id/models"
	usecases "messages/usecases/models"
	stdHttp "net/http"
)

func GetByUserID(repo *services.FetchMessagesByUserID) func(c http.Context) error {
	return func(c http.Context) error {
		request := models.GetByUserIDRequest{}

		err := c.Bind(&request)
		if err != nil {
			return err
		}

		err = c.Validate(&request)
		if err != nil {
			c.JSON(stdHttp.StatusUnprocessableEntity, err.Error())
			return nil
		}

		messages, err := repo.FetchMessagesByUserID(context.Background(), request.ID)
		if err != nil {
			return err
		}

		viewMessages := lo.Map(messages, func(message domain.Message, _ int) usecases.Message {
			return usecases.Message{
				ID:        message.ID,
				UserID:    message.UserID,
				Message:   message.Message,
				CreatedAt: message.CreatedAt,
				UpdatedAt: message.UpdatedAt,
			}
		})

		c.JSON(stdHttp.StatusOK, viewMessages)

		return nil
	}
}
