package application

import (
	"time"
	"strconv"
)

type LogObj struct {
	Lat  float64
	Long float64
	Time time.Time
	LogEntry string
}

func GetNewLogOjb(timeLog, format, cordx, cordy string) (*LogObj, error) {
	timeL, err := time.Parse(format, timeLog)
	if err != nil {
		return nil, err
	}
	cX, err := strconv.ParseFloat(cordx, 64)
	if err != nil {
		return nil, err
	}
	cY, err := strconv.ParseFloat(cordy, 64)
	if err != nil {
		return nil, err
	}
	return &LogObj{
		Lat:  cX,
		Long: cY,
		Time: timeL,
	}, nil
}