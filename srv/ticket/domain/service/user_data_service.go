package service

import "github.com/mamachengcheng/12306/srv/ticket/domain/respository"

type ITicketDataService interface {
}

func NewTicketDataService(trainRepository respository.ITicketRepository) ITicketDataService {
	return &TicketDataService{TicketRepository: trainRepository}
}

type TicketDataService struct {
	TicketRepository respository.ITicketRepository
}

