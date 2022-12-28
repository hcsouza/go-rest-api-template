package inMemory

import (
	"context"

	"github.com/hcsouza/go-rest-api-template/internal/core/domain"
)

type TicketRepository struct {
	tickets []domain.Ticket
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{
		tickets: []domain.Ticket{{Id: "123"}, {Id: "456"}},
	}
}

func (repo TicketRepository) FindAll(logCtx context.Context) ([]domain.Ticket, error) {
	return repo.tickets, nil
}
