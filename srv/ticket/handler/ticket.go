package handler

import (
	"context"
	"github.com/mamachengcheng/12306/services/ticket/domain/service"

	ticket "github.com/mamachengcheng/12306/srv/ticket/proto"
)

type Ticket struct{
	TicketDataService service.ITicketDataService
}

func (t *Ticket) BookTickets(ctx context.Context, in *ticket.BookTicketsRequest, out *ticket.BookTicketsReply) error {
	out = &ticket.BookTicketsReply{IsSuccess: false}

	return nil
}

func (t *Ticket) RefundTicket(ctx context.Context, in *ticket.RefundTicketsRequest, out *ticket.RefundTicketsReply) error {
	out = &ticket.RefundTicketsReply{IsSuccess: false}

	return nil
}