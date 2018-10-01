package consts

var (
	MonsterStatus monsterStatus
)

func init() {
	MonsterStatus.Watk = 0x1
	MonsterStatus.Wdef = 0x2
	MonsterStatus.Matk = 0x4
	MonsterStatus.Mdef = 0x8
	MonsterStatus.Acc = 0x10
	MonsterStatus.Avoid = 0x20
	MonsterStatus.Speed = 0x40
	MonsterStatus.Stun = 0x80 //this is possibly only the bowman stun
	MonsterStatus.Freeze = 0x100
	MonsterStatus.Poison = 0x200
	MonsterStatus.Seal = 0x400
	MonsterStatus.Taunt = 0x800
	MonsterStatus.WeaponAttackUp = 0x1000
	MonsterStatus.WeaponDefenseUp = 0x2000
	MonsterStatus.MagicAttackUp = 0x4000
	MonsterStatus.MagicDefenseUp = 0x8000
	MonsterStatus.Doom = 0x10000
	MonsterStatus.ShadowWeb = 0x20000
	MonsterStatus.WeaponImmunity = 0x40000
	MonsterStatus.MagicImmunity = 0x80000
	MonsterStatus.NinjaAmbush = 0x400000
	MonsterStatus.Hypnotized = 0x10000000
}

type monsterStatus struct {
	Watk, Wdef, Matk, Mdef, Acc, Avoid, Speed, Stun, Freeze, Poison, Seal, Taunt int
	WeaponAttackUp, WeaponDefenseUp, MagicAttackUp, MagicDefenseUp               int
	Doom, ShadowWeb, WeaponImmunity, MagicImmunity, NinjaAmbush, Hypnotized      int
}
