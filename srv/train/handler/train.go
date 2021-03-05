package handler

import (
	"context"
	"github.com/mamachengcheng/12306/srv/train/domain/service"
	train2 "github.com/mamachengcheng/12306/srv/train/proto/train"
)

type Train struct{
	TrainDataService service.ITrainDataService
}


func (t *Train) GetStationList(ctx context.Context, in *train2.GetStationListRequest, out *train2.GetStationListReply) error {
	out = &train2.GetStationListReply{StationList: nil}
	var err error = nil
	out.StationList, err = t.TrainDataService.GetStationList(in.InitialName)
	if err != nil {
		return err
	}
	return nil
}

func (t *Train) SearchStation(ctx context.Context, in *train2.SearchStationRequest, out *train2.SearchStationReply) error {
	out = &train2.SearchStationReply{StationList: nil}
	var err error = nil
	out.StationList, err = t.TrainDataService.SearchStation(in.Key)
	if err != nil {
		return err
	}
	return nil
}


// TODO: Request param should not be scheduleID but startDate, startStationID and endStationID

func (t *Train) GetScheduleList(ctx context.Context, in *train2.GetScheduleListRequest, out *train2.GetScheduleListReply) error {
	out = &train2.GetScheduleListReply{ScheduleList: nil}
	var err error = nil
	out.ScheduleList, err = t.TrainDataService.GetScheduleList(in.ScheduleID)
	if err != nil {
		return err
	}
	return nil
}



func (t *Train) GetStops(ctx context.Context, in *train2.GetStopsRequest, out *train2.GetStopsReply) error {
	out = &train2.GetStopsReply{Stops: nil}
	var err error = nil
	out.Stops, err = t.TrainDataService.GetStop(in.ScheduleID)
	if err != nil {
		return err
	}
	return nil
}


