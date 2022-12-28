package domain

import "context"

type Ticket struct {
	Id string `json:"id"`
}

type TicketUseCase interface {
	FindAll(context.Context) ([]Ticket, error)
}

type TicketRepository interface {
	FindAll(logCtx context.Context) ([]Ticket, error)
}
