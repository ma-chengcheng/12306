package respository

import (
	"github.com/mamachengcheng/12306/services/train/domain/model"
	"gorm.io/gorm"
)

type ITrainRepository interface {
	InitTable() error
	GetStationList(string) ([]model.Station, error)
	SearchStation(string) ([]model.Station, error)
	GetScheduleList(int64) ([]model.Schedule, error)
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


func (u *TrainRepository) GetScheduleList(scheduleID int64) ([]model.Schedule, error) {
	var schedules []model.Schedule
	return schedules, u.mysqlDB.Where("schedule_id = ?", scheduleID).Find(&schedules).Error
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
