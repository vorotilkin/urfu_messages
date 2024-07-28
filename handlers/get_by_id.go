package handlers

import (
	"context"
	"messages/domain/services"
	"messages/pkg/http"
	"messages/usecases/get_by_id/models"
	usecases "messages/usecases/models"
	stdHttp "net/http"
)

func GetByID(repo *services.FetchMessageByID) func(c http.Context) error {
	return func(c http.Context) error {
		request := models.GetByIDRequest{}

		err := c.Bind(&request)
		if err != nil {
			return err
		}

		err = c.Validate(&request)
		if err != nil {
			c.JSON(stdHttp.StatusUnprocessableEntity, err.Error())
			return nil
		}

		message, err := repo.FetchMessageByID(context.Background(), request.ID)
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
