package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hcsouza/go-rest-api-template/internal/core/domain"
)

type ticketHandler struct {
	interactor domain.TicketUseCase
}

func NewTicketHandler(gRouter *gin.RouterGroup, interactor domain.TicketUseCase) {
	handler := &ticketHandler{
		interactor: interactor,
	}

	gRouter.GET("/tickets", handler.GetTicketsHandler)

}

func (handler *ticketHandler) GetTicketsHandler(c *gin.Context) {

	logCtx := context.Background()

	actions, err := handler.interactor.FindAll(logCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, actions)
}
