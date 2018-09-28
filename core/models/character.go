package models

import "github.com/juju/errors"

type Character struct {
	charID          int32
	userID          int32
	worldID         int32
	name            string
	gender          byte
	skin            byte
	face            int32
	hair            int32
	level           byte
	job             int16
	str             int16
	dex             int16
	intt            int16
	luk             int16
	hp              int16
	maxHP           int16
	mp              int16
	maxMP           int16
	ap              int16
	sp              int16
	exp             int32
	fame            int16
	currentMap      int32
	currentMapPos   byte
	previousMap     int32
	feeMarketReturn int32
	mesos           int32
	equipSlotSize   byte
	useSlotSize     byte
	setupSlotSize   byte
	etcSlotSize     byte
	cashSlotSize    byte

	items []Item

	skills map[int32]int32

	x        int16
	y        int16
	foothold int16
	state    byte
	chairID  int32

	omokWins, omokTies, omokLosses int32
}

func (c *Character) GetSkills() map[int32]int32 {
	val := c.skills

	return val
}

func (c *Character) SetSkills(val map[int32]int32) {
	c.skills = val
}

func (c *Character) UpdateSkill(id, level int32) {
	c.skills[id] = level
}

func (c *Character) GetItems() []Item {
	val := c.items

	return val
}

func (c *Character) SetItems(val []Item) {
	c.items = val
}

func (c *Character) GetCharID() int32 {
	val := c.charID

	return val
}

func (c *Character) SetCharID(val int32) {
	c.charID = val
}

func (c *Character) GetUserID() int32 {
	val := c.userID

	return val
}

func (c *Character) SetUserID(val int32) {
	c.userID = val
}

func (c *Character) GetWorldID() int32 {
	val := c.worldID

	return val
}

func (c *Character) SetWorldID(val int32) {
	c.worldID = val
}

func (c *Character) GetName() string {
	val := c.name

	return val
}

func (c *Character) SetName(val string) {
	c.name = val
}

func (c *Character) GetGender() byte {
	val := c.gender

	return val
}

func (c *Character) SetGender(val byte) {

	c.gender = val

}

func (c *Character) GetSkin() byte {

	val := c.skin

	return val
}

func (c *Character) SetSkin(val byte) {

	c.skin = val

}

func (c *Character) GetFace() int32 {

	val := c.face

	return val
}

func (c *Character) SetFace(val int32) {

	c.face = val

}

func (c *Character) GetHair() int32 {

	val := c.hair

	return val
}

func (c *Character) SetHair(val int32) {

	c.hair = val

}

func (c *Character) GetLevel() byte {

	val := c.level

	return val
}

func (c *Character) SetLevel(val byte) {

	c.level = val

}

func (c *Character) GetJob() int16 {

	val := c.job

	return val
}

func (c *Character) SetJob(val int16) {

	c.job = val

}

func (c *Character) GetStr() int16 {

	val := c.str

	return val
}

func (c *Character) SetStr(val int16) {

	c.str = val

}

func (c *Character) GetDex() int16 {

	val := c.dex

	return val
}

func (c *Character) SetDex(val int16) {
	c.dex = val
}

func (c *Character) GetInt() int16 {

	val := c.intt

	return val
}

func (c *Character) SetInt(val int16) {
	c.intt = val
}

func (c *Character) GetLuk() int16 {

	val := c.luk

	return val
}

func (c *Character) SetLuk(val int16) {
	c.luk = val
}

func (c *Character) GetHP() int16 {
	val := c.hp

	return val
}

func (c *Character) SetHP(val int16) {
	c.hp = val
}

func (c *Character) GetMaxHP() int16 {
	val := c.maxHP

	return val
}

func (c *Character) SetMaxHP(val int16) {
	c.maxHP = val
}

func (c *Character) GetMP() int16 {

	val := c.mp

	return val
}

func (c *Character) SetMP(val int16) {

	c.mp = val

}

func (c *Character) GetMaxMP() int16 {

	val := c.maxMP

	return val
}

func (c *Character) SetMaxMP(val int16) {

	c.maxMP = val

}

func (c *Character) GetAP() int16 {

	val := c.ap

	return val
}

func (c *Character) SetAP(val int16) {

	c.ap = val

}
func (c *Character) GetSP() int16 {

	val := c.sp

	return val
}

func (c *Character) SetSP(val int16) {

	c.sp = val

}

func (c *Character) GetEXP() int32 {

	val := c.exp

	return val
}

func (c *Character) SetEXP(val int32) {

	c.exp = val

}

func (c *Character) GetFame() int16 {

	val := c.fame

	return val
}

func (c *Character) SetFame(val int16) {

	c.fame = val

}

func (c *Character) GetCurrentMap() int32 {

	val := c.currentMap

	return val
}

func (c *Character) SetCurrentMap(val int32) {

	c.currentMap = val

}

func (c *Character) GetCurrentMapPos() byte {

	val := c.currentMapPos

	return val
}

func (c *Character) SetCurrentMapPos(val byte) {

	c.currentMapPos = val

}

func (c *Character) GetPreviousMap() int32 {

	val := c.previousMap

	return val
}

func (c *Character) SetPreviousMap(val int32) {

	c.previousMap = val

}

func (c *Character) GetFeeMarketReturn() int32 {

	val := c.feeMarketReturn

	return val
}

func (c *Character) SetFreeMarketReturn(val int32) {

	c.feeMarketReturn = val

}

func (c *Character) GetMesos() int32 {

	val := c.mesos

	return val
}

func (c *Character) SetMesos(val int32) {

	c.mesos = val

}

func (c *Character) GetEquipSlotSize() byte {

	val := c.equipSlotSize

	return val
}
func (c *Character) SetEquipSlotSize(val byte) {

	c.equipSlotSize = val

}

func (c *Character) GetUsetSlotSize() byte {

	val := c.useSlotSize

	return val
}

func (c *Character) SetUseSlotSize(val byte) {

	c.useSlotSize = val

}

func (c *Character) GetSetupSlotSize() byte {

	val := c.setupSlotSize

	return val
}
func (c *Character) SetSetupSlotSize(val byte) {

	c.setupSlotSize = val

}

func (c *Character) GetEtcSlotSize() byte {

	val := c.etcSlotSize

	return val
}
func (c *Character) SetEtcSlotSize(val byte) {

	c.etcSlotSize = val

}

func (c *Character) GetCashSlotSize() byte {

	val := c.cashSlotSize

	return val
}
func (c *Character) SetCashSlotSize(val byte) {

	c.cashSlotSize = val

}

func (c *Character) GetX() int16 {

	val := c.x

	return val
}

func (c *Character) SetX(val int16) {

	c.x = val

}

func (c *Character) GetY() int16 {

	val := c.y

	return val
}

func (c *Character) SetY(val int16) {

	c.y = val

}

func (c *Character) GetFoothold() int16 {

	val := c.foothold

	return val
}

func (c *Character) SetFoothold(val int16) {

	c.foothold = val

}

func (c *Character) GetState() byte {

	val := c.state

	return val
}

func (c *Character) SetState(val byte) {

	c.state = val

}

func (c *Character) GetChairID() int32 {

	val := c.chairID

	return val
}

func (c *Character) SetChairID(val int32) {

	c.chairID = val

}

func (c *Character) GetOmokWins() int32 {

	val := c.omokWins

	return val
}

func (c *Character) SetOmokWins(val int32) {

	c.omokWins = val

}

func (c *Character) GetOmokTies() int32 {

	val := c.omokTies

	return val
}

func (c *Character) SetOmokTies(val int32) {

	c.omokTies = val

}

func (c *Character) GetOmokLosses() int32 {

	val := c.omokLosses

	return val
}

func (c *Character) SetOmokLosses(val int32) {

	c.omokLosses = val

}

const (
	characterFields = `
	id,
	userID,
	worldID,
	name,
	gender,
	skin,
	hair,
	face,
	level,
	job,
	str,
	dex,
	intt,
	luk,
	hp,
	maxHP,
	mp,
	maxMP,
	ap,
	sp,
	exp,
	fame,
	mapID,
	mapPos,
	previousMapID,
	mesos,
	equipSlotSize,
	useSlotSize,
	setupSlotSize,
	etcSlotSize,
	cashSlotSize
	`
)

func scanCharacterRowData(s rowScanner, r *Character) error {
	if err := s.Scan(
		&r.charID,
		&r.userID,
		&r.worldID,
		&r.name,
		&r.gender,
		&r.skin,
		&r.hair,
		&r.face,
		&r.level,
		&r.job,
		&r.str,
		&r.dex,
		&r.intt,
		&r.luk,
		&r.hp,
		&r.maxHP,
		&r.mp,
		&r.maxMP,
		&r.ap,
		&r.sp,
		&r.exp,
		&r.fame,
		&r.currentMap,
		&r.currentMapPos,
		&r.previousMap,
		&r.mesos,
		&r.equipSlotSize,
		&r.useSlotSize,
		&r.setupSlotSize,
		&r.etcSlotSize,
		&r.cashSlotSize,
	); nil != err {
		return errors.Trace(err)
	}
	return nil
}

func SelectCharacterByCharacterID(charID int32) (*Character, error) {
	row := GetGlobalDB().QueryRow(`SELECT `+characterFields+` FROM characters WHERE charID = ?`,
		charID)

	var c Character
	if err := scanCharacterRowData(row, &c); nil != err {
		return nil, errors.Trace(err)
	}

	return &c, nil
}

func SelectCharacterByUserID(userID int32) (*Character, error) {
	row := GetGlobalDB().QueryRow(`SELECT `+characterFields+` FROM characters WHERE userID = ?`,
		userID)

	var c Character
	if err := scanCharacterRowData(row, &c); nil != err {
		return nil, errors.Trace(err)
	}

	return &c, nil
}

func SelectCharacterByUserIDWorldID(userID int32, worldID int32) ([]*Character, error) {
	rows, err := GetGlobalDB().Query(`SELECT `+characterFields+` FROM characters WHERE userID = ? AND worldID = ?`,
		userID, worldID)
	if nil != err {
		return nil, errors.Trace(err)
	}
	defer rows.Close()

	res := make([]*Character, 0, 4)
	for rows.Next() {
		var c Character
		if err = scanCharacterRowData(rows, &c); nil != err {
			return nil, errors.Trace(err)
		}
		res = append(res, &c)
	}
	return res, nil
}

func SelectCharacterNameCount(name string) (int, error) {
	row := GetGlobalDB().QueryRow("SELECT COUNT(*) FROM characters WHERE name = ?", name)
	var cnt int
	if err := row.Scan(&cnt); nil != err {
		return 0, errors.Trace(err)
	}
	return cnt, nil
}

func InsertCharacter(ch *Character) (int64, error) {
	res, err := GetGlobalDB().Exec(`INSERT INTO characters (
		name,
		userID,
		worldID,
		face,
		hair,
		skin,
		gender,
		str,
		dex,
		intt,
		luk) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		ch.name,
		ch.userID,
		ch.worldID,
		ch.face,
		ch.hair,
		ch.skin,
		ch.gender,
		ch.str,
		ch.dex,
		ch.intt,
		ch.luk)
	if nil != err {
		return 0, errors.Trace(err)
	}
	insertID, err := res.LastInsertId()
	return insertID, errors.Trace(err)
}

func DeleteCharacter(charID int64) error {
	_, err := GetGlobalDB().Exec("DELETE FROM characters WHERE charID = ?", charID)
	return errors.Trace(err)
}
