package consts

var (
	ScrollResult scrollResult
)

func init() {
	ScrollResult.Success = 0
	ScrollResult.Fail = 1
	ScrollResult.Curse = 2
}

type scrollResult struct {
	Success int
	Fail    int
	Curse   int
}
