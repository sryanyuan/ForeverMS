package consts

var (
	Disease disease
)

func init() {
	Disease.Null = 0
	Disease.Slow = 0x01
	Disease.Seduce = 0x80
	Disease.Fishable = 0x100
	Disease.Curse = 0x200
	Disease.Confuse = 0x80000
	Disease.Stun = 0x2000000000000
	Disease.Poison = 0x4000000000000
	Disease.Darkness = 0x10000000000000
	Disease.Weaken = 0x4000000000000000
}

type disease struct {
	Null     int64
	Slow     int64
	Seduce   int64
	Fishable int64
	Curse    int64
	Confuse  int64
	Stun     int64
	Poison   int64
	Seal     int64
	Darkness int64
	Weaken   int64
}
