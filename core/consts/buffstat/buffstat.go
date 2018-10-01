package buffstat

const (
	Morph          = 0x2 //变形
	Recovery       = 0x4 //恢复
	MapleWarrior   = 0x8 //冒险岛勇士
	Stance         = 0x10
	SharpEyes      = 0x20             //火眼晶晶 - 赋予组队成员针对敌人寻找弱点并给予敌人致命伤的能力
	ManaReflection = 0x40             //魔法反击
	ShadowClaw     = 0x100            // 暗器伤人
	Infinity       = 0x20000000000000 //终极无限 - 一定时间内搜集周围的魔力,不消耗魔法值
	HolyShield     = 0x400            //圣灵之盾
	Hamstring      = 0x800
	Blind          = 0x1000
	Concentrate    = 0x2000 // another no op buff
	EchoOfHero     = 0x8000
	GhostMorph     = 0x20000    // ??? Morphs you into a ghost - no idea what for
	Dash           = 0x60000000 //0x60000000
	BerserkFury    = 0x8000000
	EnergyCharge   = 0x800000000
	MonsterRiding  = 0x10000000000
	Watk           = 0x100000000
	Wdef           = 0x200000000
	Matk           = 0x400000000
	Mdef           = 0x800000000
	Acc            = 0x1000000000
	Avoid          = 0x2000000000
	Hands          = 0x4000000000
	Speed          = 0x8000000000
	Jump           = 0x10000000000
	MagicGuard     = 0x20000000000
	Darksight      = 0x40000000000 // also used by gm hide
	Booster        = 0x80000000000
	SpeedInfusion  = 0x800000000000
	Powerguard     = 0x100000000000
	Hyperbodyhp    = 0x200000000000
	Hyperbodymp    = 0x400000000000
	Invincible     = 0x800000000000
	Soularrow      = 0x1000000000000
	Stun           = 0x2000000000000
	Poison         = 0x4000000000000
	Seal           = 0x8000000000000
	Darkness       = 0x10000000000000
	Combo          = 0x20000000000000
	Summon         = 0x20000000000000 //hack buffstat for summons ^.- =does/should not increase damage... hopefully <3
	WkCharge       = 0x40000000000000
	Dragonblood    = 0x80000000000000 // another funny buffstat...
	HolySymbol     = 0x100000000000000
	Mesoup         = 0x200000000000000
	Shadowpartner  = 0x400000000000000
	//0x8000000000000
	Pickpocket = 0x800000000000000
	Puppet     = 0x800000000000000 // HACK - shares buffmask with pickpocket - odin special ^.-
	Mesoguard  = 0x1000000000000000
	Weaken     = 0x4000000000000000 //SWITCH_CONTROLS=0x8000000000000L
)
