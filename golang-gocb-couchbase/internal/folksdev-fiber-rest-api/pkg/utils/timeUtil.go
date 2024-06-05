package utils

import "time"

//go:generate mockery --name=ITimeUtil
type ITimeUtil interface {
	EpochNow() int64
	ParseDateTime(datetime string, datetimeFormat string) int64
}

type timeUtil struct {
}

func NewTimeUtil() ITimeUtil {
	return &timeUtil{}
}

func (t timeUtil) EpochNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (t timeUtil) ParseDateTime(datetime string, datetimeFormat string) int64 {
	localTime, _ := time.ParseInLocation(datetimeFormat, datetime, time.Local)
	return localTime.UnixNano() / int64(time.Millisecond)
}
