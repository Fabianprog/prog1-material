package tests

func Quadrat(n, i int) int {
	if n == 0 {
		return -1
	}
	if i == 0 {
		return -1
	}
	if n < 0 {
		return -1
	}
	if i*i == n {
		return i
	}
	return Quadrat(n, i-1)
}
