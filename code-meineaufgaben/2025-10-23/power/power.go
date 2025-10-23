package power

func Pow2(n int) int {
	if n == 0 {
		return 1
	}
	return Pow2(n-1) * 2

}
