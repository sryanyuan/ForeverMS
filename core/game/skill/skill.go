package skill

import (
	"strconv"

	"github.com/sryanyuan/ForeverMS/core/consts"

	"github.com/sryanyuan/ForeverMS/core/game/dataloader"

	"github.com/sryanyuan/ForeverMS/core/wz"

	"github.com/sryanyuan/ForeverMS/core/game/stateffect"
)

type Skill struct {
	SkillID       int
	Element       int
	AnimationTime int
	HasCharge     bool
	effects       []*stateffect.StatEffect
}

func (s *Skill) GetSkillID() int {
	return s.SkillID
}

func (s *Skill) GetEffect(level int) *stateffect.StatEffect {
	if level-1 >= len(s.effects) {
		return nil
	}
	return s.effects[level-1]
}

func (s *Skill) GetMaxLevel() int {
	return len(s.effects)
}

func (s *Skill) CanBeLearned(job int) bool {
	skillForJob := s.SkillID / 10000
	if job/100 != skillForJob/100 && skillForJob/100 != 0 {
		// Wrong job
		return false
	}
	if skillForJob/10%10 > job/10%10 {
		// Wrong 2nd job
		return false
	}
	if skillForJob%10 > job%10 {
		// Wrong 3rd/4th job
		return false
	}
	return true
}

func (s *Skill) Is4thJob() bool {
	return s.SkillID/10000%10 == 2
}

func (s *Skill) IsBeginnerSkill() bool {
	skillStr := strconv.Itoa(s.SkillID)
	if len(skillStr) == 4 || len(skillStr) == 1 {
		return true
	}
	return false
}

func LoadFromData(id int, data wz.MapleData) *Skill {
	res := &Skill{}
	isBuff := false
	skillType := dataloader.ConvertPathIntDefault("skillType", data, -1)
	elem := wz.GetString(data.ChildByPath("elemAttr"))
	if nil != elem && len(*elem) > 0 {
		res.Element = consts.Element.FromChar((*elem)[0])
	} else {
		res.Element = consts.Element.Neutral
	}
	// unfortunatly this is only set for a few skills so we have to do some more to figure out if it's a buff &#65533;.o
	effect := data.ChildByPath("effect")
	if skillType != -1 {
		if skillType == 2 {
			isBuff = true
		}
	} else {
		action := data.ChildByPath("action")
		hit := data.ChildByPath("hit")
		ball := data.ChildByPath("ball")
		isBuff = (effect != nil && hit == nil && ball == nil)
		if action != nil {
			zchild := data.ChildByPath("0")
			if nil != zchild && zchild.Type() == wz.STRING && *wz.GetString(zchild) == "alert2" {
				isBuff = true
			}
		}

		switch id {
		case 1121006: // rush
			fallthrough
		case 1221007: // rush
			fallthrough
		case 1311005: // sacrifice
			fallthrough
		case 1321003: // rush
			fallthrough
		case 2111002: // explosion
			fallthrough
		case 2111003: // poison mist
			fallthrough
		case 2301002: // heal
			fallthrough
		case 3110001: // mortal blow
			fallthrough
		case 3210001: // mortal blow
			fallthrough
		case 4101005: // drain
			fallthrough
		case 4111003: // shadow web
			fallthrough
		case 4201004: // steal
			fallthrough
		case 4221006: // smokescreen
			fallthrough
		case 9101000: // heal + dispel
			fallthrough
		case 1121001: // monster magnet
			fallthrough
		case 1221001: // monster magnet
			fallthrough
		case 1321001: // monster magnet
			fallthrough
		case 5201006: // Recoil Shot
			fallthrough
		case 5111004: // energy drain
			fallthrough
		case 14111001: // Shadow Web
			fallthrough
		case 20001006:
			fallthrough
		case 20001007:
			fallthrough
		case 21000000:
			fallthrough
		case 21001001:
			fallthrough
		case 21000002:
			fallthrough
		case 14101006:
			fallthrough
		case 21100000:
			fallthrough
		case 21100001:
			fallthrough
		case 21120006: //�荤�虫��杈�
			fallthrough
		case 21100002:
			fallthrough
		case 21100004:
			fallthrough
		case 21110000:
			fallthrough
		case 5221010:
			fallthrough
		//case 21110002:
		case 21110003:
			fallthrough
		case 21110006:
			fallthrough
		case 21120001:
			fallthrough
		case 21120002:
			fallthrough
		case 21120004:
			fallthrough
		case 21120005:
			fallthrough
		case 21121008:
			isBuff = false
			break
		//楠�澹��㈢������
		case 12001004: //flame
			fallthrough
		case 11101002: //缁�����
			fallthrough
		case 21100005: //杩����歌�
			fallthrough
		case 11001004: //soul
			fallthrough
		case 14001005: //dark soul?
			fallthrough
		case 13001004: //storm sprite
			fallthrough
		case 15001003:
			fallthrough
		case 5001005:
			fallthrough
		case 14001003: // Dark Sight
			fallthrough
		case 15000000: // Bullet Time
			fallthrough
		case 15001004: // Lightning
			fallthrough
		case 11101001: // Sword Booster
			fallthrough
		case 11101003: // Rage
			fallthrough
		case 11101004: // Soul Blade
			fallthrough
		case 11101005: // Soul Rush
			fallthrough
		case 12101000: // Meditation
			fallthrough
		case 12101001: // Slow
			fallthrough
		case 12101004: // Spell Booster
			fallthrough
		case 12101005: // Elemental Reset
			fallthrough
		case 13101001: // Bow Booster
			fallthrough
		case 13101003: // Soul Arrow : Bow
			fallthrough
		case 13101005: // Storm Brakes
			fallthrough
		case 13101006: // Wind Walk
			fallthrough
		case 14100005: // Vanish
			fallthrough
		case 14101002: // Claw Booster
			fallthrough
		case 14101003: // Haste
			fallthrough
		case 15100004: // Energy Charge
			fallthrough
		case 15101002: // Knuckle Booster
			fallthrough
		case 15101006: // Lightning Charge
			fallthrough
		case 11111001: // Combo Attack
			fallthrough
		case 11111007: // Soul Charge
			fallthrough
		case 12111002: // Seal is one
			fallthrough
		case 13111004: // Puppet
			fallthrough
		case 13111005: // Albatross
			fallthrough
		case 14111000: // Shadow Partner
			fallthrough
		case 15111001: // Energy Drain
			fallthrough
		case 15111002: // Transformation
			fallthrough
		case 15111005: // Speed Infusion
			fallthrough
		case 15111006: // Spark
			fallthrough
		case 15111007: // Shark Wave
			fallthrough
		case 13001002: //focus
			fallthrough
		case 12001001: // Magic Guard
			fallthrough
		case 12001002: // Magic Armor
			fallthrough
		case 11001001: // Iron Body
			fallthrough
		case 12111004: //ifrit KoC
			fallthrough
		case 10000012: // Blessing of the Spirit
			fallthrough
		case 10001000: // Three Snails
			fallthrough
		case 10001001: // Recovery
			fallthrough
		case 10001002: // Nimble Feet
			fallthrough
		case 10001003: // Legendary Spiri
			fallthrough
		case 5221003:
			fallthrough
		case 10001004: // Monster Rider
			fallthrough
		case 10001005: // Echo of Hero
			fallthrough
		case 1001: // recovery
			fallthrough
		case 1002: // nimble feet
			fallthrough
		case 1004: // monster riding
			fallthrough
		case 1005: // echo of hero
			fallthrough
		case 1001003: // iron body
			fallthrough
		case 1101004: // sword booster
			fallthrough
		case 1201004: // sword booster
			fallthrough
		case 1101005: // axe booster
			fallthrough
		case 1201005: // bw booster
			fallthrough
		case 1301004: // spear booster
			fallthrough
		case 1301005: // polearm booster
			fallthrough
		case 3101002: // bow booster
			fallthrough
		case 3201002: // crossbow booster
			fallthrough
		case 4101003: // claw booster
			fallthrough
		case 4201002: // dagger booster
			fallthrough
		case 1101007: // power guard
			fallthrough
		case 1201007: // power guard
			fallthrough
		case 1101006: // rage
			fallthrough
		case 1301006: // iron will
			fallthrough
		case 1301007: // hyperbody
			fallthrough
		case 1111002: // combo attack
			fallthrough
		case 1211006: // blizzard charge bw
			fallthrough
		case 1211004: // fire charge bw
			fallthrough
		case 1211008: // lightning charge bw
			fallthrough
		case 1221004: // divine charge bw
			fallthrough
		case 1211003: // fire charge sword
			fallthrough
		case 1211005: // ice charge sword
			fallthrough
		case 1211007: // thunder charge sword
			fallthrough
		case 1221003: // holy charge sword
			fallthrough
		case 1311008: // dragon blood
			fallthrough
		case 1121000: // maple warrior
			fallthrough
		case 1221000: // maple warrior
			fallthrough
		case 1321000: // maple warrior
			fallthrough
		case 2121000: // maple warrior
			fallthrough
		case 2221000: // maple warrior
			fallthrough
		case 2321000: // maple warrior
			fallthrough
		case 3121000: // maple warrior
			fallthrough
		case 3221000: // maple warrior
			fallthrough
		case 4121000: // maple warrior
			fallthrough
		case 4221000: // maple warrior
			fallthrough
		case 1121002: // power stance
			fallthrough
		case 1221002: // power stance
			fallthrough
		case 1321002: // power stance
			fallthrough
		case 1121010: // enrage
			fallthrough
		case 1321007: // beholder
			fallthrough
		case 1320008: // beholder healing
			fallthrough
		case 1320009: // beholder buff
			fallthrough
		case 2001002: // magic guard
			fallthrough
		case 2001003: // magic armor
			fallthrough
		case 2101001: // meditation
			fallthrough
		case 2201001: // meditation
			fallthrough
		case 2301003: // invincible
			fallthrough
		case 2301004: // bless
			fallthrough
		case 2111005: // spell booster
			fallthrough
		case 2211005: // spell booster
			fallthrough
		case 2311003: // holy symbol
			fallthrough
		case 2311006: // summon dragon
			fallthrough
		case 2121004: // infinity
			fallthrough
		case 2221004: // infinity
			fallthrough
		case 2321004: // infinity
			fallthrough
		case 2321005: // holy shield
			fallthrough
		case 2121005: // elquines
			fallthrough
		case 2221005: // ifrit
			fallthrough
		case 2321003: // bahamut
			fallthrough
		case 3121006: // phoenix
			fallthrough
		case 3221005: // frostprey
			fallthrough
		case 3111002: // puppet
			fallthrough
		case 3211002: // puppet
			fallthrough
		case 3111005: // silver hawk
			fallthrough
		case 3211005: // golden eagle
			fallthrough
		case 3001003: // focus
			fallthrough
		case 3101004: // soul arrow bow
			fallthrough
		case 3201004: // soul arrow crossbow
			fallthrough
		case 3121002: // sharp eyes
			fallthrough
		case 3221002: // sharp eyes
			fallthrough
		case 3121008: // concentrate
			fallthrough
		case 3221006: // blind
			fallthrough
		case 4001003: // dark sight
			fallthrough
		case 4101004: // haste
			fallthrough
		case 4201003: // haste
			fallthrough
		case 4111001: // meso up
			fallthrough
		case 4111002: // shadow partner
			fallthrough
		case 4121006: // shadow stars
			fallthrough
		case 4211003: // pick pocket
			fallthrough
		case 4211005: // meso guard
			fallthrough
		case 5111005: // Transformation (Buccaneer)
			fallthrough
		case 5121003: // Super Transformation (Viper)
			fallthrough
		case 5220002: // wrath of the octopi
			fallthrough
		case 5211001: // Pirate octopus summon
			fallthrough
		case 5211002: // Pirate bird summon
			fallthrough
		case 5221006: // BattleShip
			fallthrough
		case 9001000: // haste
			fallthrough
		case 9101001: // super haste
			fallthrough
		case 9101002: // holy symbol
			fallthrough
		case 9101003: // bless
			fallthrough
		case 9101004: // hide
			fallthrough
		case 9101008: // hyper body
			fallthrough
		case 1121011: // hero's will
			fallthrough
		case 1221012: // hero's will
			fallthrough
		case 1321010: // hero's will
			fallthrough
		case 2321009: // hero's will
			fallthrough
		case 2221008: // hero's will
			fallthrough
		case 2121008: // hero's will
			fallthrough
		case 3121009: // hero's will
			fallthrough
		case 3221008: // hero's will
			fallthrough
		case 4121009: // hero's will
			fallthrough
		case 4221008: // hero's will
			fallthrough
		case 2101003: // slow
			fallthrough
		case 2201003: // slow
			fallthrough
		case 2111004: // seal
			fallthrough
		case 2211004: // seal
			fallthrough
		case 1111007: // armor crash
			fallthrough
		case 1211009: // magic crash
			fallthrough
		case 1311007: // power crash
			fallthrough
		case 2311005: // doom
			fallthrough
		case 2121002: // mana reflection
			fallthrough
		case 2221002: // mana reflection
			fallthrough
		case 2321002: // mana reflection
			fallthrough
		case 2311001: // dispel
			fallthrough
		case 1201006: // threaten
			fallthrough
		case 4121004: // ninja ambush
			fallthrough
		case 4221004: // ninja ambush
			fallthrough
		case 21001003:
			fallthrough
		case 20001001:
			fallthrough
		case 20001002:
			fallthrough
		case 20001004:
			fallthrough
		case 20001005:
			fallthrough

		case 20001010:
			fallthrough
		case 20001011:
			fallthrough
		case 21101003:
			fallthrough
		case 21111001:
			fallthrough

		//  case 1100002:
		case 21121000:
			fallthrough
		case 21121003:
			fallthrough
		case 21120007:
			fallthrough
		//case 21110004:
		case 21111005:
			fallthrough
		case 9001004:
			isBuff = true
			break
		}
	}
	keydown := data.ChildByPath("keydown")
	if nil != keydown {
		res.HasCharge = true
	}
	levels := data.ChildByPath("level")
	if nil != levels && nil != levels.Children() {
		for _, level := range levels.Children() {
			statEffect := stateffect.LoadSkillEffectFromData(level, id, isBuff, level.Name())
			res.effects = append(res.effects, statEffect)
		}
	}
	if effect != nil && nil != effect.Children() {
		for _, effectEntry := range effect.Children() {
			res.AnimationTime += dataloader.ConvertPathIntDefault("delay", effectEntry, 0)
		}
	}

	return res
}
