package api

import (
	"messages/domain/services"
	"messages/handlers"
	"messages/pkg/http"
)

func Registry(
	server *http.Server,
	createSvc *services.CreateMessage,
	deleteSvc *services.DeleteMessageByID,
	fetchByIDSvc *services.FetchMessageByID,
	fetchByUserSvc *services.FetchMessagesByUserID,
	updateSvc *services.UpdateMessageByID,
) {
	server.POST("/message/new", handlers.Post(createSvc))
	server.DELETE("/message/:id", handlers.Delete(deleteSvc))
	server.GET("/message/:id", handlers.GetByID(fetchByIDSvc))
	server.GET("/message/user/:id", handlers.GetByUserID(fetchByUserSvc))
	server.PUT("/message/update", handlers.Put(updateSvc))
}
