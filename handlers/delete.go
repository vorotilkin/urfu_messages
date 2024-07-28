package handlers

import (
	"context"
	"messages/domain/services"
	"messages/pkg/http"
	"messages/usecases/delete/models"
	stdHttp "net/http"
)

func Delete(repo *services.DeleteMessageByID) func(c http.Context) error {
	return func(c http.Context) error {
		request := models.DeleteRequest{}

		err := c.Bind(&request)
		if err != nil {
			return err
		}

		err = c.Validate(&request)
		if err != nil {
			c.JSON(stdHttp.StatusUnprocessableEntity, err.Error())
			return nil
		}

		ok, err := repo.DeleteMessageByID(context.Background(), request.ID)
		if err != nil {
			return err
		}

		if !ok {
			c.NoContent(stdHttp.StatusNotFound)
		}

		c.NoContent(stdHttp.StatusNoContent)

		return nil
	}
}
