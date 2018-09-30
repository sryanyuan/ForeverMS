package models

import "github.com/juju/errors"

/*
-- ----------------------------
-- Table structure for `characters`
-- ----------------------------
DROP TABLE IF EXISTS `characters`;
CREATE TABLE `characters` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `accountid` int(11) NOT NULL DEFAULT '0',
  `world` int(11) NOT NULL DEFAULT '0',
  `name` varchar(13) NOT NULL DEFAULT '',
  `level` int(11) NOT NULL DEFAULT '0',
  `exp` int(11) NOT NULL DEFAULT '0',
  `str` int(11) NOT NULL DEFAULT '0',
  `dex` int(11) NOT NULL DEFAULT '0',
  `luk` int(11) NOT NULL DEFAULT '0',
  `intt` int(11) NOT NULL DEFAULT '0',
  `hp` int(11) NOT NULL DEFAULT '0',
  `mp` int(11) NOT NULL DEFAULT '0',
  `maxhp` int(11) NOT NULL DEFAULT '0',
  `maxmp` int(11) NOT NULL DEFAULT '0',
  `meso` int(11) NOT NULL DEFAULT '0',
  `hpApUsed` int(11) NOT NULL DEFAULT '0',
  `mpApUsed` int(11) NOT NULL DEFAULT '0',
  `job` int(11) NOT NULL DEFAULT '0',
  `skincolor` int(11) NOT NULL DEFAULT '0',
  `gender` int(11) NOT NULL DEFAULT '0',
  `fame` int(11) NOT NULL DEFAULT '0',
  `hair` int(11) NOT NULL DEFAULT '0',
  `face` int(11) NOT NULL DEFAULT '0',
  `ap` int(11) NOT NULL DEFAULT '0',
  `sp` int(11) NOT NULL DEFAULT '0',
  `map` int(11) NOT NULL DEFAULT '0',
  `spawnpoint` int(11) NOT NULL DEFAULT '0',
  `gm` int(11) NOT NULL DEFAULT '0',
  `party` int(11) NOT NULL DEFAULT '0',
  `buddyCapacity` int(11) NOT NULL DEFAULT '25',
  `createdate` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `rank` int(10) unsigned NOT NULL DEFAULT '1',
  `rankMove` int(11) NOT NULL DEFAULT '0',
  `jobRank` int(10) unsigned NOT NULL DEFAULT '1',
  `jobRankMove` int(11) NOT NULL DEFAULT '0',
  `guildid` int(10) unsigned NOT NULL DEFAULT '0',
  `guildrank` int(10) unsigned NOT NULL DEFAULT '5',
  `allianceRank` int(10) unsigned NOT NULL DEFAULT '5',
  `messengerid` int(10) unsigned NOT NULL DEFAULT '0',
  `messengerposition` int(10) unsigned NOT NULL DEFAULT '4',
  `reborns` int(11) NOT NULL DEFAULT '0',
  `pvpkills` int(9) NOT NULL DEFAULT '0',
  `pvpdeaths` int(9) NOT NULL DEFAULT '0',
  `clan` tinyint(1) NOT NULL DEFAULT '-1',
  `mountlevel` int(9) NOT NULL DEFAULT '1',
  `mountexp` int(9) NOT NULL DEFAULT '0',
  `mounttiredness` int(9) NOT NULL DEFAULT '0',
  `married` int(10) unsigned NOT NULL DEFAULT '0',
  `partnerid` int(10) unsigned NOT NULL DEFAULT '0',
  `cantalk` int(10) unsigned NOT NULL DEFAULT '1',
  `zakumlvl` int(10) unsigned NOT NULL DEFAULT '0',
  `marriagequest` int(10) unsigned NOT NULL DEFAULT '0',
  `omok` int(4) DEFAULT NULL,
  `matchcard` int(4) DEFAULT NULL,
  `omokwins` int(4) DEFAULT NULL,
  `omoklosses` int(4) DEFAULT NULL,
  `omokties` int(4) DEFAULT NULL,
  `matchcardwins` int(4) DEFAULT NULL,
  `matchcardlosses` int(4) DEFAULT NULL,
  `matchcardties` int(4) DEFAULT NULL,
  `MerchantMesos` int(11) NOT NULL DEFAULT '0',
  `HasMerchant` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `gmtext` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `equipslots` int(11) NOT NULL DEFAULT '24',
  `useslots` int(11) NOT NULL DEFAULT '24',
  `setupslots` int(11) NOT NULL DEFAULT '24',
  `etcslots` int(11) NOT NULL DEFAULT '24',
  `bosspoints` int(11) NOT NULL DEFAULT '0',
  `bossrepeats` int(11) NOT NULL DEFAULT '0',
  `nextBQ` bigint(20) unsigned NOT NULL DEFAULT '0',
  `LeaderPoints` int(11) NOT NULL,
  `pqPoints` int(11) NOT NULL,
  `votePoints` int(11) NOT NULL DEFAULT '0',
  `occupation` int(11) NOT NULL DEFAULT '1',
  `jqpoints` int(11) NOT NULL DEFAULT '0',
  `CashPoints` int(11) NOT NULL DEFAULT '0',
  `jqrank` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `accountid` (`accountid`),
  KEY `party` (`party`),
  KEY `ranking1` (`level`,`exp`),
  KEY `ranking2` (`gm`,`job`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=latin1;
*/

type Character struct {
	ID                int64
	AccountID         int64
	World             int32
	Name              string
	Level             int
	Exp               int
	Str               int
	Dex               int
	Luk               int
	Intt              int
	HP                int
	MP                int
	MaxHP             int
	MaxMP             int
	Meso              int
	HpApUsed          int
	MpApUsed          int
	Job               int
	SkinColor         int
	Gender            int
	Fame              int
	Hair              int
	Face              int
	Ap                int
	Sp                int
	MapID             int
	SpawnPoint        int
	GM                int
	Party             int
	BuddyCapacity     int
	CreateDate        *int64
	Rank              int
	RankMove          int
	JobRank           int
	JobRankMove       int
	GuildID           int
	GuildRank         int
	AllianceRank      int
	MessagerID        int
	MessengerPosition int
	Reborns           int
	PvpKills          int
	PvpDeaths         int
	Clan              int
	MountLevel        int
	Mountexp          int
	MountTiredness    int
	Married           int
	PartnerID         int
	CanTalk           int
	ZakumLvl          int
	MarriageQuest     int
	Omok              int
	MatchCard         int
	OmokWins          int
	OmokLosses        int
	OmokTies          int
	MatchCardWins     int
	MatchCardLosses   int
	MatchCardTies     int
	MerchantMesos     int
	HasMerchant       int
	GMText            int
	EquipSlots        int
	UseSlots          int
	SetupSlots        int
	EtcSlots          int
	BossPoints        int
	BossRepeats       int
	NextBQ            int64
	LeaderPoints      int
	PQPoints          int
	VotePoints        int
	Occupation        int
	JQPoints          int
	CashPoints        int
	JQRank            int
}

const (
	characterFields = `
	id,
	accountid,
	world,
	name,
	level,
	exp,
	str,
	dex,
	luk,
	intt,
	hp,
	mp,
	maxhp,
	maxmp,
	meso,
	hpApUsed,
	mpApUsed,
	job,
	skincolor,
	gender,
	fame,
	hair,
	face,
	ap,
	sp,
	mapid,
	spawnpoint,
	gm,
	party,
	buddyCapacity,
	createdate,
	rank,
	rankMove,
	jobRank,
	jobRankMove,
	guildid,
	guildrank,
	allianceRank,
	messengerid,
	messengerposition,
	reborns,
	pvpkills,
	pvpdeaths,
	clan,
	mountlevel,
	mountexp,
	mounttiredness,
	married,
	partnerid,
	cantalk,
	zakumlvl,
	marriagequest,
	omok,
	matchcard,
	omokwins,
	omoklosses,
	omokties,
	matchcardwins,
	matchcardlosses,
	matchcardties,
	MerchantMesos,
	HasMerchant,
	gmtext,
	equipslots,
	useslots,
	setupslots,
	etcslots,
	bosspoints,
	bossrepeats,
	nextBQ,
	LeaderPoints,
	pqPoints,
	votePoints,
	occupation,
	jqpoints,
	CashPoints,
	jqrank
	`
)

func scanCharacterFullRowData(s rowScanner, r *Character) error {
	if err := s.Scan(
		&r.ID,
		&r.AccountID,
		&r.World,
		&r.Name,
		&r.Level,
		&r.Exp,
		&r.Str,
		&r.Dex,
		&r.Luk,
		&r.Intt,
		&r.HP,
		&r.MP,
		&r.MaxHP,
		&r.MaxMP,
		&r.Meso,
		&r.HpApUsed,
		&r.MpApUsed,
		&r.Job,
		&r.SkinColor,
		&r.Gender,
		&r.Fame,
		&r.Hair,
		&r.Face,
		&r.Ap,
		&r.Sp,
		&r.MapID,
		&r.SpawnPoint,
		&r.GM,
		&r.Party,
		&r.BuddyCapacity,
		&r.CreateDate,
		&r.Rank,
		&r.RankMove,
		&r.JobRank,
		&r.JobRankMove,
		&r.GuildID,
		&r.GuildRank,
		&r.AllianceRank,
		&r.MessagerID,
		&r.MessengerPosition,
		&r.Reborns,
		&r.PvpKills,
		&r.PvpDeaths,
		&r.Clan,
		&r.MountLevel,
		&r.Mountexp,
		&r.MountTiredness,
		&r.Married,
		&r.PartnerID,
		&r.CanTalk,
		&r.ZakumLvl,
		&r.MarriageQuest,
		&r.Omok,
		&r.MatchCard,
		&r.OmokWins,
		&r.OmokLosses,
		&r.OmokTies,
		&r.MatchCardWins,
		&r.MatchCardLosses,
		&r.MatchCardTies,
		&r.MerchantMesos,
		&r.HasMerchant,
		&r.GMText,
		&r.EquipSlots,
		&r.UseSlots,
		&r.SetupSlots,
		&r.EtcSlots,
		&r.BossPoints,
		&r.BossRepeats,
		&r.NextBQ,
		&r.LeaderPoints,
		&r.PQPoints,
		&r.VotePoints,
		&r.Occupation,
		&r.JQPoints,
		&r.CashPoints,
		&r.JQRank,
	); nil != err {
		return errors.Trace(err)
	}
	return nil
}

func SelectCharacterByCharacterID(charID int32) (*Character, error) {
	row := GetGlobalDB().QueryRow(`SELECT `+characterFields+` FROM characters WHERE id = ?`,
		charID)

	var c Character
	if err := scanCharacterFullRowData(row, &c); nil != err {
		return nil, errors.Trace(err)
	}

	return &c, nil
}

func SelectCharacterByAccountID(accountID int32) (*Character, error) {
	row := GetGlobalDB().QueryRow(`SELECT `+characterFields+` FROM characters WHERE accountid = ?`,
		accountID)

	var c Character
	if err := scanCharacterFullRowData(row, &c); nil != err {
		return nil, errors.Trace(err)
	}

	return &c, nil
}

func SelectCharacterByAccountIDWorldID(accountID int32, worldID int32) ([]*Character, error) {
	rows, err := GetGlobalDB().Query(`SELECT `+characterFields+` FROM characters WHERE accountid = ? AND world = ?`,
		accountID, worldID)
	if nil != err {
		return nil, errors.Trace(err)
	}
	defer rows.Close()

	res := make([]*Character, 0, 4)
	for rows.Next() {
		var c Character
		if err = scanCharacterFullRowData(rows, &c); nil != err {
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
		ch.Name,
		ch.AccountID,
		ch.World,
		ch.Face,
		ch.Hair,
		ch.SkinColor,
		ch.Gender,
		ch.Str,
		ch.Dex,
		ch.Intt,
		ch.Luk)
	if nil != err {
		return 0, errors.Trace(err)
	}
	insertID, err := res.LastInsertId()
	return insertID, errors.Trace(err)
}

func DeleteCharacter(charID int64) error {
	_, err := GetGlobalDB().Exec("DELETE FROM characters WHERE id = ?", charID)
	return errors.Trace(err)
}
