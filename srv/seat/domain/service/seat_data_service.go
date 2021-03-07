package service

import (
	"github.com/mamachengcheng/12306/srv/seat/domain/respository"
)

type ISeatDataService interface {
	CountRemainingSeats(seatType uint32, trainID, scheduleStatus uint64) uint32
	LockSeats(seatType, number uint32, trainID, scheduleStatus uint64) ([]uint64, error)
	ReleaseSeat(seatID, scheduleStatus uint64) error
}

func NewSeatDataService(userRepository respository.ISeatRepository) ISeatDataService {
	return &SeatDataService{SeatRepository: userRepository}
}

type SeatDataService struct {
	SeatRepository respository.ISeatRepository
}

func (s *SeatDataService) CountRemainingSeats(seatType uint32, trainID, scheduleStatus uint64) (number uint32) {
	seats, _ := s.SeatRepository.FindSeats(seatType, trainID, scheduleStatus)
	return uint32(len(seats))
}

func (s *SeatDataService) LockSeats(seatType, number uint32, trainID, scheduleStatus uint64) (seatsID []uint64, err error) {
	seats, err := s.SeatRepository.FindSeats(seatType, trainID, scheduleStatus)
	if err != nil {
		return seatsID, err
	}
	return s.SeatRepository.UpdateSeats(number, seats, scheduleStatus)
}

func (s *SeatDataService) ReleaseSeat(seatID, scheduleStatus uint64) error {
	err := s.SeatRepository.UpdateSeat(seatID, scheduleStatus)
	return err
}
