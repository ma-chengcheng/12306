package handler

import (
	"context"
	"github.com/mamachengcheng/12306/srv/seat/domain/service"
	seat "github.com/mamachengcheng/12306/srv/seat/proto/seat"
)

type Seat struct {
	SeatDataService service.ISeatDataService
}

func (s *Seat) CountRemainingSeats(ctx context.Context, in *seat.CountRemainingSeatsRequest, out *seat.CountRemainingSeatsReply) error {
	out.Number = s.SeatDataService.CountRemainingSeats(in.SeatType, in.TrainID, in.ScheduleStatus)
	return nil
}

func (s *Seat) GetSeats(ctx context.Context, in *seat.GetSeatsRequest, out *seat.GetSeatsReply) error {
	seatIDs, err := s.SeatDataService.LockSeats(in.SeatType, in.Number, in.TrainID, in.ScheduleStatus)

	if err != nil {
		return err
	}

	if len(seatIDs) == int(in.Number) {
		out.SeatIDs = seatIDs
		out.IsSuccess = true
	}

	return nil
}

func (s *Seat) RollbackSeat(ctx context.Context, in *seat.RollbackSeatRequest, out *seat.RollbackSeatReply) error {
	err := s.SeatDataService.ReleaseSeat(in.SeatID, in.ScheduleStatus)
	if err == nil {
		out.IsSuccess = true
	}
	return err
}
