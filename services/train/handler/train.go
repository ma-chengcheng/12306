package handler

import (
	"context"
	"github.com/mamachengcheng/12306/services/train/domain/service"
	train "github.com/mamachengcheng/12306/services/train/proto"
)

type Train struct{
	TrainDataService service.ITrainDataService
}


func (t *Train) GetStationList(ctx context.Context, in *train.GetStationListRequest, out *train.GetStationListReply) error {
	out = &train.GetStationListReply{StationList: nil}

	return nil
}

func (t *Train) SearchStation(ctx context.Context, in *train.SearchStationRequest, out *train.SearchStationReply) error {
	out = &train.SearchStationReply{StationList: nil}

	return nil
}



func (t *Train) GetScheduleList(ctx context.Context, in *train.GetScheduleListRequest, out *train.GetScheduleListReply) error {
	out = &train.GetScheduleListReply{ScheduleList: nil}

	return nil
}



func (t *Train) GetStops(ctx context.Context, in *train.GetStopsRequest, out *train.GetStopsReply) error {
	out = &train.GetStopsReply{Stops: nil}

	return nil
}


