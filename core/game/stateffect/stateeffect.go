package stateffect

import (
	"container/list"
	"image"

	"github.com/sryanyuan/ForeverMS/core/consts"
	"github.com/sryanyuan/ForeverMS/core/consts/buffstat"
	"github.com/sryanyuan/ForeverMS/core/consts/monsterstatus"
	"github.com/sryanyuan/ForeverMS/core/game/dataloader"
	"github.com/sryanyuan/ForeverMS/core/wz"
)

type BuffStatValue struct {
	Stat int64
	Val  int
}

func NewBuffStatVal(stat int64, val int) *BuffStatValue {
	return &BuffStatValue{
		Stat: stat,
		Val:  val,
	}
}

type StatEffect struct {
	AttackCount   int
	bulletCount   int
	bulletConsume int
	Cooldown      int
	cureDebuffs   list.List
	Duration      int
	fixDamage     int

	HP, MP int

	hpR, mpR float64

	isMorph            bool
	itemCon, itemConNo int
	lt                 *image.Point
	rb                 *image.Point
	mastery, rang      int
	mobCount           int
	moneyCon           int

	MonsterStatus map[int]int

	morphId int
	moveTo  int

	MPCon, HPCon int
	overTime     bool
	prop         float64
	remark       string
	ret          interface{}
	skill        bool
	sourceID     int
	statups      list.List

	Watk, Matk, Wdef, Mdef, Acc, Avoid, Hands, Speed, Jump int

	X, Y, Z int
	Damage  int
}

func (s *StatEffect) GetRemark() string {
	return s.remark
}

func LoadSkillEffectFromData(source wz.MapleData, skillID int, overTime bool, lvl string) *StatEffect {
	return LoadFromData(source, skillID, true, overTime, "Level "+lvl)
}

func LoadFromData(source wz.MapleData, sourceID int, skill bool, overTime bool, remark string) *StatEffect {
	eff := &StatEffect{}
	eff.Duration = dataloader.ConvertPathIntDefault("time", source, -1)
	eff.HP = dataloader.ConvertPathIntDefault("hp", source, 0)
	eff.hpR = float64(dataloader.ConvertPathIntDefault("hpR", source, 0)) / 100.0
	eff.MP = dataloader.ConvertPathIntDefault("mp", source, 0)
	eff.mpR = float64(dataloader.ConvertPathIntDefault("mpR", source, 0)) / 100.0
	eff.MPCon = dataloader.ConvertPathIntDefault("mpCon", source, 0)
	eff.HPCon = dataloader.ConvertPathIntDefault("hpCon", source, 0)
	var iprop = dataloader.ConvertPathIntDefault("prop", source, 100)
	eff.prop = float64(iprop) / 100.0
	eff.AttackCount = dataloader.ConvertPathIntDefault("attackCount", source, 1)
	eff.mobCount = dataloader.ConvertPathIntDefault("mobCount", source, 1)
	eff.Cooldown = dataloader.ConvertPathIntDefault("cooltime", source, 0)
	eff.morphId = dataloader.ConvertPathIntDefault("morph", source, 0)
	eff.isMorph = false
	if eff.morphId > 0 {
		eff.isMorph = true
	}
	eff.remark = remark
	eff.sourceID = sourceID
	eff.skill = skill

	if !eff.skill && eff.Duration > -1 {
		eff.overTime = true
	} else {
		eff.Duration *= 1000 // items have their times stored in ms, of course
		eff.overTime = overTime
	}

	eff.Watk = dataloader.ConvertPathIntDefault("pad", source, 0)
	eff.Wdef = dataloader.ConvertPathIntDefault("pdd", source, 0)
	eff.Watk = dataloader.ConvertPathIntDefault("mad", source, 0)
	eff.Mdef = dataloader.ConvertPathIntDefault("mdd", source, 0)
	eff.Acc = dataloader.ConvertPathIntDefault("acc", source, 0)
	eff.Avoid = dataloader.ConvertPathIntDefault("ava", source, 0)
	eff.Speed = dataloader.ConvertPathIntDefault("speed", source, 0)
	eff.Jump = dataloader.ConvertPathIntDefault("jump", source, 0)

	if eff.overTime && eff.GetSummonMovementType() == consts.SummonMovementType.None {
		AddBuffStatValToListIfNotZero(&eff.statups, buffstat.Watk, eff.Watk)
		AddBuffStatValToListIfNotZero(&eff.statups, buffstat.Wdef, eff.Wdef)
		AddBuffStatValToListIfNotZero(&eff.statups, buffstat.Matk, eff.Matk)
		AddBuffStatValToListIfNotZero(&eff.statups, buffstat.Mdef, eff.Mdef)
		AddBuffStatValToListIfNotZero(&eff.statups, buffstat.Acc, eff.Acc)
		AddBuffStatValToListIfNotZero(&eff.statups, buffstat.Avoid, eff.Avoid)
		AddBuffStatValToListIfNotZero(&eff.statups, buffstat.Speed, eff.Speed)
		AddBuffStatValToListIfNotZero(&eff.statups, buffstat.Jump, eff.Jump)
		AddBuffStatValToListIfNotZero(&eff.statups, buffstat.Morph, eff.morphId)
	}

	ltd := source.ChildByPath("lt")
	if nil != ltd {
		eff.lt = wz.GetPoint(ltd)
		eff.rb = wz.GetPoint(source.ChildByPath("rb"))
	}

	x := dataloader.ConvertPathIntDefault("x", source, 0)
	eff.X = x
	eff.Y = dataloader.ConvertPathIntDefault("y", source, 0)
	eff.Z = dataloader.ConvertPathIntDefault("z", source, 0)
	eff.Damage = dataloader.ConvertPathIntDefault("damage", source, 100)
	eff.bulletCount = dataloader.ConvertPathIntDefault("bulletCount", source, 1)
	eff.bulletConsume = dataloader.ConvertPathIntDefault("bulletConsume", source, 0)
	eff.moneyCon = dataloader.ConvertPathIntDefault("moneyCon", source, 0)

	eff.itemCon = dataloader.ConvertPathIntDefault("itemCon", source, 0)
	eff.itemConNo = dataloader.ConvertPathIntDefault("itemConNo", source, 0)
	eff.fixDamage = dataloader.ConvertPathIntDefault("fixdamage", source, 0)
	eff.moveTo = dataloader.ConvertPathIntDefault("moveTo", source, -1)

	eff.mastery = dataloader.ConvertPathIntDefault("mastery", source, 0)
	eff.rang = dataloader.ConvertPathIntDefault("range", source, 0)

	if dataloader.ConvertPathIntDefault("poison", source, 0) > 0 {
		eff.cureDebuffs.PushBack(consts.Disease.Poison)
	}
	if dataloader.ConvertPathIntDefault("seal", source, 0) > 0 {
		eff.cureDebuffs.PushBack(consts.Disease.Seal)
	}
	if dataloader.ConvertPathIntDefault("darkness", source, 0) > 0 {
		eff.cureDebuffs.PushBack(consts.Disease.Darkness)
	}
	if dataloader.ConvertPathIntDefault("weakness", source, 0) > 0 {
		eff.cureDebuffs.PushBack(consts.Disease.Weaken)
	}
	if dataloader.ConvertPathIntDefault("curse", source, 0) > 0 {
		eff.cureDebuffs.PushBack(consts.Disease.Curse)
	}

	monsterStatus := make(map[int]int)

	if eff.skill {
		switch eff.sourceID {
		case 2001002: // 魔法守卫
		case 12001001:
			eff.statups.PushBack(NewBuffStatVal(buffstat.MagicGuard, x))
			break
		case 2301003: // 无敌
			eff.statups.PushBack(NewBuffStatVal(buffstat.Invincible, x))
			break
		case 9001004: // 隐藏
			eff.Duration = 60 * 120 * 1000
			eff.overTime = true
			fallthrough
			//夜行者 NIGHT_KNIGHT
		case 14001003: // 隐身
			eff.statups.PushBack(NewBuffStatVal(buffstat.Darksight, x))
			break
		case 4001003: // darksight
			eff.statups.PushBack(NewBuffStatVal(buffstat.Darksight, x))
			break
		case 4211003: // pickpocket
			eff.statups.PushBack(NewBuffStatVal(buffstat.Pickpocket, x))
			break
		case 4211005: // mesoguard
			eff.statups.PushBack(NewBuffStatVal(buffstat.Mesoguard, x))
			break
		case 4111001: // mesoup
			eff.statups.PushBack(NewBuffStatVal(buffstat.Mesoup, x))
			break
		case 4111002: // 影分身
		case 14111000:
			eff.statups.PushBack(NewBuffStatVal(buffstat.Shadowpartner, x))
			break
		case 3101004: // 灵魂的箭头
		case 3201004:
		case 21120002: //战神之舞
		case 13101003: //精灵使者-无形箭
			eff.statups.PushBack(NewBuffStatVal(buffstat.Soularrow, x))
			break
		case 2311002: //时空门
			eff.statups.PushBack(NewBuffStatVal(buffstat.Soularrow, x))
			break
		case 1211003:
		case 1211004:
		case 1211005:
		case 1221003: //圣灵之剑
		case 11101002: //终极剑
		case 15111006: //闪光击
		case 1221004: //圣灵之锤
		case 1211006: // 寒冰钝器
		case 21111005: //冰雪矛
		case 1211007:

		case 1211008:
		case 15101006:
			eff.statups.PushBack(NewBuffStatVal(buffstat.WkCharge, x))
			break
		case 21120007: //战神之盾
			eff.statups.PushBack(NewBuffStatVal(buffstat.MagicGuard, x))
			break
		case 21111001: //灵巧击退MAGIC_GUARD
			eff.statups.PushBack(NewBuffStatVal(buffstat.Wdef, x))
			break
		case 21100005: //连环吸血
			eff.statups.PushBack(NewBuffStatVal(buffstat.Infinity, x))
			break
		case 21101003: //抗压
			eff.statups.PushBack(NewBuffStatVal(buffstat.Powerguard, x))
			break
		/*case 21111005: //冰雪矛
		  eff.statups.PushBack(NewBuffStatVal(buffstat.WK_CHARGE,x));
		  eff.duration *= 2; //冰雪矛冰冻时间为2秒
		  //monsterStatus.Add(MonsterStatus.SPEED, eff.x));
		                      monsterStatus.Add(MonsterStatus.FREEZE, 1));
		  break;*/
		case 1101004:
		case 1101005: // booster
		case 1201004:
		case 1201005:
		case 1301004:
		case 1301005: //快速矛
		case 2111005: // spell booster, do these work the same? :做这些工作的助推器,法术一样吗?
		case 2211005:
		case 3101002:
		case 3201002:
		case 4101003:
		case 4201002:
		case 5101006:
		case 5201003:
		case 11101001: //魂骑士-快速剑
		case 12101004: //炎骑士-魔法狂暴
		case 13101001: //精灵使者-快速箭
		case 14101002:
		case 15101002:
		case 21001003:
			eff.statups.PushBack(NewBuffStatVal(buffstat.Booster, x))
			break
		case 5121009:
			eff.statups.PushBack(NewBuffStatVal(buffstat.SpeedInfusion, x))
			fallthrough
		case 15111005:
			eff.statups.PushBack(NewBuffStatVal(buffstat.SpeedInfusion, x))
			break
		case 1101006: // 愤怒
		case 11101003: // 愤怒之火
			eff.statups.PushBack(NewBuffStatVal(buffstat.Wdef, eff.Wdef))
			fallthrough
		case 1121010: // enrage
			eff.statups.PushBack(NewBuffStatVal(buffstat.Watk, eff.Watk))
			break
		case 1301006: // iron will
			eff.statups.PushBack(NewBuffStatVal(buffstat.Mdef, eff.Mdef))
			fallthrough
		case 1001003: // iron body
			eff.statups.PushBack(NewBuffStatVal(buffstat.Wdef, eff.Wdef))
			break
		case 2001003: // magic armor
			eff.statups.PushBack(NewBuffStatVal(buffstat.Wdef, eff.Wdef))
			break
		case 2101001: // meditation
		case 2201001: // meditation
			eff.statups.PushBack(NewBuffStatVal(buffstat.Matk, eff.Matk))
			break
		case 4101004: // 轻功
		case 4201003: // 轻功
		case 9001001: // gm轻功
			eff.statups.PushBack(NewBuffStatVal(buffstat.Speed, eff.Speed))
			eff.statups.PushBack(NewBuffStatVal(buffstat.Jump, eff.Jump))
			break
		case 2301004: //祝福
			eff.statups.PushBack(NewBuffStatVal(buffstat.Wdef, eff.Wdef))
			eff.statups.PushBack(NewBuffStatVal(buffstat.Mdef, eff.Mdef))
			fallthrough
		case 3001003: //二连射
			eff.statups.PushBack(NewBuffStatVal(buffstat.Acc, eff.Acc))
			eff.statups.PushBack(NewBuffStatVal(buffstat.Avoid, eff.Avoid))
			break
		case 9001003: //GM祝福
			eff.statups.PushBack(NewBuffStatVal(buffstat.Matk, eff.Matk))
			//goto case 9001003;
		case 3121008: // 集中精力
			eff.statups.PushBack(NewBuffStatVal(buffstat.Watk, eff.Watk))
			break
		case 5001005: //疾驰
			eff.statups.PushBack(NewBuffStatVal(buffstat.Dash, x))
			eff.statups.PushBack(NewBuffStatVal(buffstat.Jump, eff.Y))
			fallthrough
		case 1101007: //伤害反击
		case 1201007:
		case 21100003:
			eff.statups.PushBack(NewBuffStatVal(buffstat.Powerguard, x))
			break
		case 1301007:
		case 9001008:
			eff.statups.PushBack(NewBuffStatVal(buffstat.Hyperbodyhp, x))
			eff.statups.PushBack(NewBuffStatVal(buffstat.Hyperbodymp, eff.Y))
			break
		case 1001: // recovery
		case 10001001:
		case 20001001:
			eff.statups.PushBack(NewBuffStatVal(buffstat.Recovery, x))
			break
		case 1111002: // combo
		case 11111001:
			eff.statups.PushBack(NewBuffStatVal(buffstat.Combo, 1))
			break
		case 1011:
		case 20001011:
			eff.statups.PushBack(NewBuffStatVal(buffstat.BerserkFury, 1))
			break
		case 1004: // monster riding
		case 10001004:
		case 20001004:
		case 5221006: // 4th Job - Pirate riding
		case 5221008:
			eff.statups.PushBack(NewBuffStatVal(buffstat.MonsterRiding, 1))
			break
		case 1311006: //dragon roar
			eff.hpR = float64(-x) / 100.0
			break
		case 1311008: // dragon blood
			eff.statups.PushBack(NewBuffStatVal(buffstat.Dragonblood, 1))
			break
		case 1121000: // maple warrior, all classes
		case 1221000:
		case 1321000:
		case 2121000:
		case 2221000:
		case 2321000:
		case 3121000:
		case 3221000:
		case 4121000:
		case 4221000:
		case 5121000:
		case 5221000:
		case 21121000:
			eff.statups.PushBack(NewBuffStatVal(buffstat.MapleWarrior, 1))
			break
		case 3121002: // sharp eyes bowmaster
		case 3221002: // sharp eyes marksmen
			eff.statups.PushBack(NewBuffStatVal(buffstat.SharpEyes, eff.X<<8|eff.Y))
			break
		case 1321007: //Beholder
		case 2221005: // ifrit
		case 2311006: // summon dragon
		case 2321003: // bahamut
		case 3121006: // phoenix
		case 5211001: // Pirate octopus summon
		case 5211002: // Pirate bird summon
		case 5220002: // wrath of the octopi
		case 11001004:
		case 12001004:
		case 13001004:
		case 14001005:
		case 15001004:
		case 12111004:
			eff.statups.PushBack(NewBuffStatVal(buffstat.Summon, 1))
			break
		case 2311003: //神圣祈祷
		case 21110000: //属性暴击
		case 9001002: //GM圣化之力
			eff.statups.PushBack(NewBuffStatVal(buffstat.HolySymbol, x))
			break
		case 4121006: // 暗器伤人
			eff.statups.PushBack(NewBuffStatVal(buffstat.ShadowClaw, 0))
			break
		case 2121004:
		case 2221004:
		case 2321004: // Infinity
			eff.statups.PushBack(NewBuffStatVal(buffstat.Infinity, x))
			break
		case 1121002:
		case 1221002:
		case 0000012: //精灵的祝福
		case 21120004: //防守策略
		case 21120009: //(隐藏) 战神之舞- 双重重击
		case 21120010: //(隐藏) 战神之舞 - 三重重击
		case 1321002: // Stance
		case 21121003: //战神的意志
			eff.statups.PushBack(NewBuffStatVal(buffstat.Stance, iprop))
			break
		case 1005: // Echo of Hero
			eff.statups.PushBack(NewBuffStatVal(buffstat.EchoOfHero, eff.X))
			break
		case 2121002: // mana reflection
		case 2221002:
		case 2321002:
			eff.statups.PushBack(NewBuffStatVal(buffstat.ManaReflection, 1))
			break
		case 2321005: // holy shield
			eff.statups.PushBack(NewBuffStatVal(buffstat.HolyShield, x))
			break
		case 3111002: // puppet ranger
		case 3211002: // puppet sniper
			eff.statups.PushBack(NewBuffStatVal(buffstat.Puppet, 1))
			break

		// -----------------------------飓风把! ----------------------------- //

		case 4001002: //混乱
			monsterStatus[monsterstatus.Watk] = eff.X
			monsterStatus[monsterstatus.Wdef] = eff.Y
			break
		case 1201006: // threaten
			monsterStatus[monsterstatus.Watk] = eff.X
			monsterStatus[monsterstatus.Wdef] = eff.Y
			break
		case 1211002: // charged blow
		case 1111008: // shout
		case 4211002: // assaulter
		case 3101005: // arrow bomb
		case 1111005: // coma: sword
		case 1111006: // coma: axe
		case 4221007: // boomerang step
		case 20001005:
		case 5101002: // Backspin Blow
		case 5101003: // Double Uppercut
		case 5121004: // Demolition
		case 14101006:
		case 21110004:
		case 21100004:
		case 5121005: // Snatch
		case 5121007: // Barrage
		case 5201004: // pirate blank shot
		case 11111003:
			monsterStatus[monsterstatus.Stun] = 1
			break
		case 4121003:
		case 4221003:
			monsterStatus[monsterstatus.Taunt] = eff.X
			monsterStatus[monsterstatus.Mdef] = eff.X
			monsterStatus[monsterstatus.Wdef] = eff.X
			break
		case 4121004: // Ninja ambush
		case 4221004: //忍者伏击
			//int damage = 2 * (c.GetPlayer().GetStr() + c.GetPlayer().GetLuk()) * (eff.damage / 100);
			monsterStatus[monsterstatus.NinjaAmbush] = 1
			break
		case 2201004: // 冰冻术
		case 20001013:
		case 2211002: // ice strike
		case 5221003:

		case 3211003: // blizzard
		case 2211006: // il elemental compo
		case 2221007: // 落霜冰破
		case 21120006: //星辰
		case 5211005: // Ice Splitter
		case 2121006: // Paralyze
			monsterStatus[monsterstatus.Freeze] = 1
			eff.Duration *= 2 // 冰冻的时间
			break
		case 2121003: // fire demon
		case 2221003: // ice demon
			monsterStatus[monsterstatus.Poison] = 1
			monsterStatus[monsterstatus.Freeze] = 1
			break
		case 2101003: // fp slow
		case 2201003: // il slow
			monsterStatus[monsterstatus.Speed] = eff.X
			break
		case 2101005: // poison breath
		case 2111006: // fp elemental compo
			monsterStatus[monsterstatus.Poison] = 1
			break
		case 2311005:
			monsterStatus[monsterstatus.Doom] = 1
			break
		case 3111005: // golden hawk
		case 3211005: // golden eagle
		case 13111004:
			eff.statups.PushBack(NewBuffStatVal(buffstat.Summon, 1))
			monsterStatus[monsterstatus.Stun] = 1
			break
		case 2121005: // elquines
		case 3221005: // frostprey
			eff.statups.PushBack(NewBuffStatVal(buffstat.Summon, 1))
			monsterStatus[monsterstatus.Freeze] = 1
			break
		case 2111004: // fp seal
		case 2211004: // il seal
		case 12111002:
			monsterStatus[monsterstatus.Seal] = 1
			break
		case 4111003: // shadow web
		case 14111001:
			monsterStatus[monsterstatus.ShadowWeb] = 1
			break
		case 3121007: // Hamstring
			eff.statups.PushBack(NewBuffStatVal(buffstat.Hamstring, x))
			monsterStatus[monsterstatus.Speed] = x
			break
		case 3221006: // Blind
			eff.statups.PushBack(NewBuffStatVal(buffstat.Blind, x))
			monsterStatus[monsterstatus.Acc] = x
			break
		case 5221009:
			monsterStatus[monsterstatus.Hypnotized] = 1
			break
		default:
			break
		}
	}

	if eff.isMorph && !eff.IsPirateMorph() {
		eff.statups.PushBack(NewBuffStatVal(buffstat.Morph, eff.morphId))
	}

	eff.MonsterStatus = monsterStatus

	return eff
}

func (s *StatEffect) IsPirateMorph() bool {
	return s.skill && (s.sourceID == 5111005 || s.sourceID == 5121003)
}

func (s *StatEffect) GetSummonMovementType() int {
	if !s.skill {
		return consts.SummonMovementType.None
	}
	switch s.sourceID {
	case 3211002: // puppet sniper
	case 3111002: // puppet ranger
	case 5211001: // octopus - pirate
	case 5220002: // advanced octopus - pirate
		return consts.SummonMovementType.Stationary
	case 3211005: // golden eagle
	case 3111005: // golden hawk
	case 2311006: // summon dragon
	case 3221005: // frostprey
	case 3121006: // phoenix
	case 5211002: // bird - pirate
		return consts.SummonMovementType.CircleFollow
	case 1321007: // 灵魂助力
	case 2121005: // 冰破魔兽
	case 2221005: // 火魔兽
	case 5221010:
	case 21121008:
	case 2321003: // 强化圣龙

		break
	case 11001004: //魂精灵
	case 12001004: //炎精灵
	case 13001004: //风精灵
	case 14001005: //夜精灵
	case 15001004: //雷精灵
	case 12111004: //火魔兽
		return consts.SummonMovementType.Follow
	}
	return consts.SummonMovementType.None
}

func AddBuffStatValToListIfNotZero(ls *list.List, buffStat int64, val int) {
	if val == 0 {
		return
	}
	ls.PushBack(&BuffStatValue{
		Stat: buffStat,
		Val:  val,
	})
}
