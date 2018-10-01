package consts

var (
	SummonMovementType summonMovementType
)

func init() {
	SummonMovementType.Stationary = 0
	SummonMovementType.Follow = 1
	SummonMovementType.CircleFollow = 2
	SummonMovementType.None = 0xff
}

type summonMovementType struct {
	None                             int
	Stationary, Follow, CircleFollow int
}
