package ackermann

func Ackermann(m, n int) int {
	if m == 0 {
		n += 1
	}
	if n == 0 {
		Ackermann(m-1, 1)
	}

	return Ackermann(m-1, Ackermann(m, n-1))

}
