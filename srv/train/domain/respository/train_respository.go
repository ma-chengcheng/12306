package respository

import (
	"github.com/mamachengcheng/12306/srv/train/domain/model"
	"gorm.io/gorm"
	"time"
)

type ITrainRepository interface {
	InitTable() error
	GetStationList(string) ([]model.Station, error)
	SearchStation(string) ([]model.Station, error)
	GetScheduleList(string, int64, int64) ([]model.Schedule, error)
	GetStop(int64) ([]model.Stop, error)
}

func NewTrainRepository(db *gorm.DB) ITrainRepository {
	return &TrainRepository{mysqlDB: db}
}

type TrainRepository struct {
	mysqlDB *gorm.DB
}



func (u *TrainRepository) SearchStation(key string) ([]model.Station, error) {
	var stations []model.Station
	return stations, u.mysqlDB.Where("station_name LIKE ? OR pinyin LIKE ?", key, key).Find(&stations).Error
}

func (u *TrainRepository) GetStationList(initialName string) ([]model.Station, error) {
	var stations []model.Station
	return stations, u.mysqlDB.Where("initial_name = ?", initialName).Find(&stations).Error
}


func (u *TrainRepository) GetScheduleList(startDate string,startStationID, endStationID  int64) ([]model.Schedule, error) {
	var schedules []model.Schedule
	startTime, _ := time.ParseInLocation("2006-01-02", startDate, time.Local)
	err := u.mysqlDB.Preload("StartStation").Preload("EndStation").Where("start_time >= ? AND end_time < ?", startTime, startTime.Add(time.Hour*24)).Where("start_station_refer = ? AND end_station_refer = ?", startStationID, endStationID).Find(&schedules).Error
	return schedules, err
}

func (u *TrainRepository) GetStop(scheduleID int64) ([]model.Stop, error) {
	var stops []model.Stop
	var schedule model.Schedule
	var train model.Train
	u.mysqlDB.Where("id = ?", scheduleID).First(&schedule)
	err := u.mysqlDB.Preload("Stops").Preload("Stops.StartStation").Where("id = ?", schedule.TrainRefer).Find(&train).Error
	stops = train.Stops
	return stops, err
}

func (u *TrainRepository) InitTable() error {
	return u.mysqlDB.AutoMigrate(&model.Station{}, &model.Schedule{}, &model.Seat{}, &model.Stop{}, &model.Train{})
}
