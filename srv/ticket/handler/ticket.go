package handler

import (
	"context"
	"github.com/mamachengcheng/12306/srv/ticket/domain/service"
	"github.com/mamachengcheng/12306/srv/ticket/proto/ticket"
)

type Ticket struct{
	TicketDataService service.ITicketDataService
}

func (t *Ticket) BookTickets(ctx context.Context, in *ticket.BookTicketsRequest, out *ticket.BookTicketsReply) error {
	out = &ticket.BookTicketsReply{IsSuccess: false}

	return nil
}

func (t *Ticket) RefundTicket(ctx context.Context, in *ticket.RefundTicketRequest, out *ticket.RefundTicketReply) error {
	out = &ticket.RefundTicketReply{IsSuccess: false}

	return nil
}