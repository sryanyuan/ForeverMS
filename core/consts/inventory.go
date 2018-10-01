package consts

var (
	InventoryType inventoryType
)

func init() {
	InventoryType.Undefined = 0
	InventoryType.Equip = 1
	InventoryType.Use = 2
	InventoryType.Setup = 3
	InventoryType.Etc = 4
	InventoryType.Cash = 5
	InventoryType.Equipped = 6
	InventoryType.Total = 7
}

type inventoryType struct {
	Undefined int
	Equip     int
	Use       int
	Setup     int
	Etc       int
	Cash      int
	Equipped  int
	Total     int
}

func (i *inventoryType) FromString(str string) int {
	switch str {
	case "Install":
		{
			return InventoryType.Setup
		}
	case "Consume":
		{
			return InventoryType.Use
		}
	case "Etc":
		{
			return InventoryType.Etc
		}
	case "Cash":
		{
			return InventoryType.Cash
		}
	case "Pet":
		{
			return InventoryType.Cash
		}
	default:
		{
			return InventoryType.Undefined
		}
	}
}
