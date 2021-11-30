package fillingBlocks

var FillingBlocks = fillingBlocks

type head struct {
	layer int
	teeth int
}

var cache map[head]int
var setMap map[int][]int

const (
	none = iota
	all
	upper
	lower
	outter
	inner
)

func init() {

	cache = map[head]int{}

	setMap = map[int][]int{
		none:   {none, all, lower, upper, outter},
		all:    {none},
		upper:  {none, lower},
		lower:  {none, upper},
		outter: {none, inner},
		inner:  {outter},
	}

}

func fillingBlocks(n int) int {

	var rf func(head, int) int
	rf = func(in head, d int) int {

		d *= 1

		var ans int
		var ok bool

		if ans, ok = cache[in]; ok {
			return ans
		}

		if in.layer == 0 {
			if in.teeth == none {
				ans = 1
			} else {
				ans = 0
			}
		} else {
			ans = 0
			for _, end := range setMap[in.teeth] {
				next := head{in.layer - 1, end}
				q := rf(next, d+1)
				ans += q
			}
			cache[in] = ans
		}

		return ans

	}

	first := head{n, none}
	ans := rf(first, 0)
	return ans

}
