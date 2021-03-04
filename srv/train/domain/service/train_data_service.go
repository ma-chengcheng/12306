package service

import (
	"github.com/mamachengcheng/12306/services/train/domain/respository"
	train "github.com/mamachengcheng/12306/services/train/proto"
)

type ITrainDataService interface {
	GetStationList(string) ([]*train.Station, error)
	SearchStation(string) ([]*train.Station, error)
	GetScheduleList(int64) ([]*train.Schedule, error)
	GetStop(int64) ([]*train.Stop, error)
}

func NewTrainDataService(trainRepository respository.ITrainRepository) ITrainDataService {
	return &TrainDataService{TrainRepository: trainRepository}
}

type TrainDataService struct {
	TrainRepository respository.ITrainRepository
}

func (t *TrainDataService) GetStationList(initialName string) ([]*train.Station, error)  {
	stations, err := t.TrainRepository.GetStationList(initialName)
	var stationList []*train.Station
	for i := range stations {
		stationList = append(stationList, &train.Station{
			ID: int64(stations[i].ID),
			StationName: stations[i].StationName,
		})
	}
	return stationList, err
}


func (t *TrainDataService) SearchStation(key string) ([]*train.Station, error) {
	stations, err := t.TrainRepository.SearchStation(key)
	var stationList []*train.Station
	for i := range stations {
		stationList = append(stationList, &train.Station{
			ID: int64(stations[i].ID),
			StationName: stations[i].StationName,
		})
	}
	return stationList, err
}


// TODO: Request param should not be scheduleID but startDate, startStationID and endStationID
func (t *TrainDataService) GetScheduleList(scheduleID int64) ([]*train.Schedule, error) {
	schedules, err := t.TrainRepository.GetScheduleList(scheduleID)
	var scheduleList []*train.Schedule
	for i := range schedules {
		scheduleList = append(scheduleList, &train.Schedule{
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


func (t *TrainDataService) GetStop(scheduleID int64) ([]*train.Stop, error){
	stops, err := t.TrainRepository.GetStop(scheduleID)
	var stopList []*train.Stop
	for i := range stops {
		stopList = append(stopList, &train.Stop{
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