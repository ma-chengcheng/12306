package service

import (
	"github.com/mamachengcheng/12306/services/train/domain/respository"
)

type ITrainDataService interface {
}

func NewTrainDataService(trainRepository respository.ITrainRepository) ITrainDataService {
	return &TrainDataService{TrainRepository: trainRepository}
}

type TrainDataService struct {
	TrainRepository respository.ITrainRepository
}
