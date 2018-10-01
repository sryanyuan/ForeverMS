package world

const (
	MaxOid   = 20000
	dropLife = 180000
)

type Map struct {
	rangedMapobjectTypes []int
	channelID            int
	mapEffect            *MapEffect
	mapTimer             *MapTimer

	monsterRate  float64
	monsterSpawn []*SpawnPoint

	boat bool

	CannotInvincible bool

	canVipRock bool

	docked        bool
	dropsDisabled bool

	hasEvent    bool
	origMobRate float64

	runningOid int

	MapID       int
	ReturnMapID int

	objects map[int]IMapObject
}

func NewMap(mapid int, channel int, returnMapID int, monsterRate float64) *Map {
	m := &Map{
		canVipRock:  true,
		runningOid:  100,
		MapID:       mapid,
		ReturnMapID: returnMapID,
		monsterRate: monsterRate,
		origMobRate: monsterRate,
		objects:     make(map[int]IMapObject),
	}
	// TODO: Add monster spawn timer processor
	return m
}

func (m *Map) AddMapObject(obj IMapObject) {

}
