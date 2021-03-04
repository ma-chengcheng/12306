package handler

import (
	"context"
	"github.com/mamachengcheng/12306/services/train/domain/service"
	train "github.com/mamachengcheng/12306/srv/train/proto"
)

type Train struct{
	TrainDataService service.ITrainDataService
}


func (t *Train) GetStationList(ctx context.Context, in *train.GetStationListRequest, out *train.GetStationListReply) error {
	out = &train.GetStationListReply{StationList: nil}
	var err error = nil
	out.StationList, err = t.TrainDataService.GetStationList(in.InitialName)
	if err != nil {
		return err
	}
	return nil
}

func (t *Train) SearchStation(ctx context.Context, in *train.SearchStationRequest, out *train.SearchStationReply) error {
	out = &train.SearchStationReply{StationList: nil}
	var err error = nil
	out.StationList, err = t.TrainDataService.SearchStation(in.Key)
	if err != nil {
		return err
	}
	return nil
}


// TODO: Request param should not be scheduleID but startDate, startStationID and endStationID

func (t *Train) GetScheduleList(ctx context.Context, in *train.GetScheduleListRequest, out *train.GetScheduleListReply) error {
	out = &train.GetScheduleListReply{ScheduleList: nil}
	var err error = nil
	out.ScheduleList, err = t.TrainDataService.GetScheduleList(in.ScheduleID)
	if err != nil {
		return err
	}
	return nil
}



func (t *Train) GetStops(ctx context.Context, in *train.GetStopsRequest, out *train.GetStopsReply) error {
	out = &train.GetStopsReply{Stops: nil}
	var err error = nil
	out.Stops, err = t.TrainDataService.GetStop(in.ScheduleID)
	if err != nil {
		return err
	}
	return nil
}


