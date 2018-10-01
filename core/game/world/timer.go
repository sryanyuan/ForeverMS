package world

import "time"

type MapTimer struct {
	duration          int64
	predictedStopTime int64
	startTime         int64

	MapToWarpTo    int
	MinLevelToWarp int
	MaxLevelToWarp int
}

func NewMapTimer(newDuration int64, mapToWarpToP int, minLevelToWrapP int, maxLevelToWrapP int) *MapTimer {
	return &MapTimer{
		duration:          newDuration,
		startTime:         time.Now().UnixNano() / 1e6,
		predictedStopTime: time.Now().UnixNano()/1e6 + newDuration,
		MapToWarpTo:       mapToWarpToP,
		MinLevelToWarp:    minLevelToWrapP,
		MaxLevelToWarp:    maxLevelToWrapP,
	}
}

func (t *MapTimer) GetTimeLeft() int64 {
	tn := time.Now().UnixNano() / 1e6
	return (t.predictedStopTime - tn) / 1000
}
