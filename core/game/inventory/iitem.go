package inventory

type IItem interface {
	GetItemID() int
	GetFlag() int
	SetFlag(int)
	GetPosition() int
	SetPosition(int)
	GetOwner() string
	GetQuantity() int
	SetQuantity(int)
	GetExpiration() int64
	GetSN() int
	SetSN(int)
	GetUID() int64
	SetUID(int64)
	GetType() int
	Copy() IItem
}
