package service

import (
	"github.com/mamachengcheng/12306/srv/train/domain/respository"
	train2 "github.com/mamachengcheng/12306/srv/train/proto/train"
)

type ITrainDataService interface {
	GetStationList(string) ([]*train2.Station, error)
	SearchStation(string) ([]*train2.Station, error)
	GetScheduleList(int64) ([]*train2.Schedule, error)
	GetStop(int64) ([]*train2.Stop, error)
}

func NewTrainDataService(trainRepository respository.ITrainRepository) ITrainDataService {
	return &TrainDataService{TrainRepository: trainRepository}
}

type TrainDataService struct {
	TrainRepository respository.ITrainRepository
}

func (t *TrainDataService) GetStationList(initialName string) ([]*train2.Station, error)  {
	stations, err := t.TrainRepository.GetStationList(initialName)
	var stationList []*train2.Station
	for i := range stations {
		stationList = append(stationList, &train2.Station{
			ID: int64(stations[i].ID),
			StationName: stations[i].StationName,
		})
	}
	return stationList, err
}


func (t *TrainDataService) SearchStation(key string) ([]*train2.Station, error) {
	stations, err := t.TrainRepository.SearchStation(key)
	var stationList []*train2.Station
	for i := range stations {
		stationList = append(stationList, &train2.Station{
			ID: int64(stations[i].ID),
			StationName: stations[i].StationName,
		})
	}
	return stationList, err
}


// TODO: Request param should not be scheduleID but startDate, startStationID and endStationID
func (t *TrainDataService) GetScheduleList(scheduleID int64) ([]*train2.Schedule, error) {
	schedules, err := t.TrainRepository.GetScheduleList(scheduleID)
	var scheduleList []*train2.Schedule
	for i := range schedules {
		scheduleList = append(scheduleList, &train2.Schedule{
			// TODO: uint to int64
			//ID: int64(schedules[i].ID),
			// TODO: in model.Station both TrainNo and TrainType are string while train.Station are int64
			//TrainNo: schedules[i].TrainNo,
			//TrainType: schedules[i].TrainType,
			// TODO: time.Time to string
			//StartTime: schedules[i].StartTime,
			//EndTime: schedules[i].EndTime,
			// TODO: uint to string
			//Duration: schedules[i].Duration,
			StartStationName: schedules[i].StartStation.StationName,
			EndStationName: schedules[i].EndStation.StationName,
		})
	}
	return scheduleList, err
}


func (t *TrainDataService) GetStop(scheduleID int64) ([]*train2.Stop, error){
	stops, err := t.TrainRepository.GetStop(scheduleID)
	var stopList []*train2.Stop
	for i := range stops {
		stopList = append(stopList, &train2.Stop{
			//TODO: uint to int64
			//NO: stops[i].No,
			//StartTime: stops[i].StartTime,
			//EndTime: stops[i].EndTime,
			//Duration: stops[i].Duration,
			StartStationName: stops[i].StartStation.StationName,
			// TODO: there is not exist EndStationName field
			//EndStationName: ,
		})
	}
	return stopList, err
}