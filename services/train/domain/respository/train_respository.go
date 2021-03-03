package respository

import (
	"github.com/mamachengcheng/12306/services/train/domain/model"
	"gorm.io/gorm"
)

type ITrainRepository interface {
	InitTable() error
}

func NewTrainRepository(db *gorm.DB) ITrainRepository {
	return &TrainRepository{mysqlDB: db}
}

type TrainRepository struct {
	mysqlDB *gorm.DB
}

func (u *TrainRepository) InitTable() error {
	return u.mysqlDB.AutoMigrate(&model.Station{}, &model.Schedule{}, &model.Seat{}, &model.Stop{}, &model.Train{})
}
