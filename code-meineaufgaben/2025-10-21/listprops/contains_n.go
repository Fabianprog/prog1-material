package listprops

// ContainsN liefert true, falls die Liste l
// String x mindestens n mal enthält.
func ContainsN(l []string, x string, n int) bool {
	y := 0
	for i := 0; i < len(l); i++ {

		if l[i] == x {
			y++
		}
	}
	return n <= y

}
