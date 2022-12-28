package tickets

import (
	"context"

	"github.com/hcsouza/go-rest-api-template/internal/core/domain"
)

type ticketUseCase struct {
	repository domain.TicketRepository
}

func NewActionUseCase(repository domain.TicketRepository) domain.TicketUseCase {
	return &ticketUseCase{
		repository: repository,
	}
}

func (interactor *ticketUseCase) FindAll(logContext context.Context) ([]domain.Ticket, error) {
	return interactor.repository.FindAll(context.Background())
}
