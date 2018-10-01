package consts

var (
	ItemType itemType
)

func init() {
	ItemType.Equip = 1
	ItemType.Item = 2
	ItemType.Pet = 3
}

type itemType struct {
	Equip int
	Item  int
	Pet   int
}
