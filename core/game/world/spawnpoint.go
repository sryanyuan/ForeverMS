package world

import (
	"image"
	"time"
)

type SpawnPoint struct {
	immobile bool
	mobTime  int64

	nextPossibleSpawn int64
	spawnedMonsters   int64

	Pos image.Point
}

func NewSpawnPoint(pos image.Point, mobTime int64) *SpawnPoint {
	return &SpawnPoint{
		Pos:               pos,
		mobTime:           mobTime,
		nextPossibleSpawn: time.Now().UnixNano() / 1e6,
	}
}

func (s *SpawnPoint) ShouldSpawnNow() bool {
	return s.ShouldSpawn(time.Now().UnixNano() / 1e6)
}

func (s *SpawnPoint) ShouldSpawn(tn int64) bool {
	if s.mobTime < 0 {
		return false
	}
	if ((s.mobTime != 0 || s.immobile) && s.spawnedMonsters > 0) || s.spawnedMonsters > 2 {
		return false
	}
	return s.nextPossibleSpawn <= tn
}
