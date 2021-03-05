package service

import (
	"github.com/mamachengcheng/12306/srv/train/domain/respository"
	train "github.com/mamachengcheng/12306/srv/train/proto"
)

type ITrainDataService interface {
	GetStationList(string) ([]*train.Station, error)
	SearchStation(string) ([]*train.Station, error)
	GetScheduleList(string, int64, int64) ([]*train.Schedule, error)
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
			ID: stations[i].ID,
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
			ID: stations[i].ID,
			StationName: stations[i].StationName,
		})
	}
	return stationList, err
}



func (t *TrainDataService) GetScheduleList(startDate string, startStationID, endStationID int64) ([]*train.Schedule, error) {
	schedules, err := t.TrainRepository.GetScheduleList(startDate, startStationID, endStationID)
	var scheduleList []*train.Schedule
	for i := range schedules {
		scheduleList = append(scheduleList, &train.Schedule{
			ID: schedules[i].ID,
			TrainNo: schedules[i].TrainNo,
			TrainType: schedules[i].TrainType,
			StartTime: schedules[i].StartTime.Format("2006-01-02"),
			EndTime: schedules[i].EndTime.Format("2006-01-02"),
			Duration: schedules[i].Duration,
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
			NO: stops[i].No,
			StartTime: stops[i].StartTime.Format("2006-01-02"),
			EndTime: stops[i].EndTime.Format("2006-01-02"),
			Duration: stops[i].Duration,
			StartStationName: stops[i].StartStation.StationName,
			// TODO: there is not exist EndStationName field
			//EndStationName: ,
		})
	}
	return stopList, err
}