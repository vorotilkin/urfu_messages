package handlers

import (
	"context"
	"messages/domain/services"
	"messages/pkg/http"
	"messages/usecases/post/models"
	stdHttp "net/http"
)

func Post(repo *services.CreateMessage) func(c http.Context) error {
	return func(c http.Context) error {
		request := models.PostRequest{}

		err := c.Bind(&request)
		if err != nil {
			return err
		}

		err = c.Validate(&request)
		if err != nil {
			c.JSON(stdHttp.StatusUnprocessableEntity, err.Error())
			return nil
		}

		ok, err := repo.CreateMessage(context.Background(), request.UserID, request.Message)
		if err != nil {
			return err
		}

		if !ok {
			c.NoContent(stdHttp.StatusNotFound)
		}

		c.NoContent(stdHttp.StatusCreated)

		return nil
	}
}
