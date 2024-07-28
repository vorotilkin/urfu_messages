package handlers

import (
	"context"
	"messages/domain/services"
	"messages/pkg/http"
	usecases "messages/usecases/models"
	"messages/usecases/put/models"
	stdHttp "net/http"
)

func Put(repo *services.UpdateMessageByID) func(c http.Context) error {
	return func(c http.Context) error {
		request := models.PutRequest{}

		err := c.Bind(&request)
		if err != nil {
			return err
		}

		err = c.Validate(&request)
		if err != nil {
			c.JSON(stdHttp.StatusUnprocessableEntity, err.Error())
			return nil
		}

		message, err := repo.UpdateMessageByID(context.Background(), request.ID, request.Message)
		if err != nil {
			return err
		}

		viewMessage := usecases.Message{
			ID:        message.ID,
			UserID:    message.UserID,
			Message:   message.Message,
			CreatedAt: message.CreatedAt,
			UpdatedAt: message.UpdatedAt,
		}

		c.JSON(stdHttp.StatusOK, viewMessage)

		return nil
	}
}
