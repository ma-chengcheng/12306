package respository

import (
	"errors"
	"github.com/mamachengcheng/12306/srv/seat/domain/model"
	"gorm.io/gorm"
)

type ISeatRepository interface {
	InitTable() error
	FindSeats(seatType uint32, trainID, ScheduleStatus uint64) ([]model.Seat, error)
	UpdateSeats(number uint32, seats []model.Seat, scheduleStatus uint64) ([]uint64, error)
	UpdateSeat(seatID, scheduleStatus uint64) error
}

func NewSeatRepository(db *gorm.DB) ISeatRepository {
	return &SeatRepository{mysqlDB: db}
}

type SeatRepository struct {
	mysqlDB *gorm.DB
}

func (u *SeatRepository) InitTable() error {
	return u.mysqlDB.AutoMigrate(&model.Seat{})
}

// 用于车票的初始化操作
func (u *SeatRepository) FindSeats(seatType uint32, trainID, scheduleStatus uint64) (seats []model.Seat, err error) {
	var _seats []model.Seat
	err = u.mysqlDB.Where("seat_type = ? AND train_id = ?", seatType, trainID).Find(&_seats).Error

	if err == nil {
		for _, seat := range _seats {
			if seat.SeatStatus&scheduleStatus == 0 {
				seats = append(seats, seat)
			}
		}
	}

	return seats, err
}

// 购票时更新座位状态信息
func (u *SeatRepository) UpdateSeats(number uint32, seats []model.Seat, scheduleStatus uint64) (seatIDs []uint64, err error) {

	err = u.mysqlDB.Transaction(func(tx *gorm.DB) error {
		for _, seat := range seats {
			if len(seatIDs) > int(number) {
				break
			}
			t := u.mysqlDB.Find(&model.Seat{}, seat.ID).Where("tag = ?", seat.Tag).Update("seat_status", scheduleStatus|seat.SeatStatus).RowsAffected
			if t > 0 {
				seatIDs = append(seatIDs, seat.ID)
			}
		}

		if len(seatIDs) != int(number) {
			return errors.New("TicketIssuanceFailed")
		}
		return nil
	})

	if err != nil {
		seatIDs = nil
	}

	return seatIDs, err
}

// 退票时更新座位状态信息
func (u *SeatRepository) UpdateSeat(seatID, scheduleStatus uint64) (err error) {
	scheduleStatus = ^scheduleStatus

	err = u.mysqlDB.Transaction(func(tx *gorm.DB) error {
		err := u.mysqlDB.Find(&model.Seat{}, seatID).Update("seat_status", gorm.Expr("seat_status&?", scheduleStatus)).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
