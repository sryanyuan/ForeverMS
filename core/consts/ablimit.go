package consts

var (
	AbilityLimit abilityLimit
)

func init() {
	AbilityLimit.MaxHP = 30000
	AbilityLimit.MaxMP = 30000
	AbilityLimit.MaxSpeed = 140
	AbilityLimit.MaxJump = 123
	AbilityLimit.MaxMagic = 2000
}

type abilityLimit struct {
	MaxHP    int
	MaxMP    int
	MaxSpeed int
	MaxJump  int
	MaxMagic int
}
