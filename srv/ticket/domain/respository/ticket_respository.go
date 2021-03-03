package respository

import (
	"github.com/mamachengcheng/12306/services/ticket/domain/model"
	"gorm.io/gorm"
)

type ITicketRepository interface {
	InitTable() error
	CreateTicket(ticket *model.Ticket) (int64, error)
}

func NewTicketRepository(db *gorm.DB) ITicketRepository {
	return &TicketRepository{mysqlDB: db}
}

type TicketRepository struct {
	mysqlDB *gorm.DB
}

func (t *TicketRepository) InitTable() error {
	return t.mysqlDB.AutoMigrate(&model.Ticket{})
}

func (t *TicketRepository) CreateTicket(ticket *model.Ticket) (int64, error) {
	return ticket.ID, t.mysqlDB.Create(ticket).Error
}