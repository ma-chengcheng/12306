package model

import "time"

type Station struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	StationName string `gorm:"not null" json:"station_name"`
	InitialName string `gorm:"not null" json:"initial_name"`
	Pinyin      string `gorm:"not null" json:"pinyin"`
	CityNo      string `gorm:"not null" json:"city_no"`
	CityName    string `gorm:"not null" json:"city_name"`
	ShowName    string `gorm:"not null" json:"show_name"`
	NameType    string `gorm:"not null" json:"name_type"`
}

type Schedule struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	TrainNo   int64    `gorm:"not null" json:"train_no"`
	TrainType int64    `gorm:"not null" json:"train_type"`
	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time `gorm:"not null" json:"end_time"`
	Duration  string      `gorm:"not null" json:"duration"`

	ScheduleStatus string `gorm:"not null" json:"ticket_status"`

	StartStation      Station `gorm:"foreignKey:StartStationRefer;not null" json:"start_station"`
	EndStation        Station `gorm:"foreignKey:EndStationRefer;not null" json:"end_station"`
	StartStationRefer uint    // Belongs to Station
	EndStationRefer   uint    // Belongs to Station

	TrainRefer uint
}

type Stop struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	No int64 `gorm:"not null" json:"no"`

	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time `gorm:"not null" json:"end_time"`
	Duration  int64  `gorm:"not null" json:"duration"`

	StartStation      Station `gorm:"foreignKey:StartStationRefer;not null" json:"start_station"`
	StartStationRefer uint    // Belongs to Station

	TrainRefer uint
}

type Train struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	Schedules []Schedule `gorm:"foreignKey:TrainRefer" json:"schedules"` // Has Many Schedules
	Stops     []Stop     `gorm:"foreignKey:TrainRefer" json:"stops"`     // Has Many Stops
}
