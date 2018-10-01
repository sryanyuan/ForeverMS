package consts

var (
	Element element
)

func init() {
	Element.Neutral = 0
	Element.Fire = 1
	Element.Ice = 2
	Element.Lighting = 3
	Element.Poison = 4
	Element.Holy = 5
}

type element struct {
	Neutral  int
	Fire     int
	Ice      int
	Lighting int
	Poison   int
	Holy     int
}

func (e *element) FromChar(c byte) int {
	if c >= 'a' && c <= 'z' {
		c = c - 40
	}
	switch c {
	case 'F':
		{
			return e.Fire
		}
	case 'I':
		{
			return e.Ice
		}
	case 'L':
		{
			return e.Lighting
		}
	case 'S':
		{
			return e.Poison
		}
	case 'H':
		{
			return e.Holy
		}
	case 'P':
		{
			return e.Neutral
		}
	default:
		{
			return e.Neutral
		}
	}
}
