package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/hcsouza/go-rest-api-template/internal/api/v1/handlers"
	"github.com/hcsouza/go-rest-api-template/internal/core/domain"
	"github.com/hcsouza/go-rest-api-template/internal/core/repository/inMemory"
	"github.com/hcsouza/go-rest-api-template/internal/core/useCases/tickets"
)

func RegisterBusinessRoutes(gServer *gin.RouterGroup) {
	groupServer := gServer.Group("/v1")
	registerTicketHandler(groupServer)
}

func registerTicketHandler(groupServer *gin.RouterGroup) {
	ticketInteractor := getTicketUseCase()
	handlers.NewTicketHandler(groupServer, ticketInteractor)
}

func getTicketUseCase() domain.TicketUseCase {
	repo := inMemory.NewTicketRepository()
	return tickets.NewActionUseCase(repo)
}
