package listprops

// ContainsNRow liefert true, falls die Liste l
// String x mindestens n mal hintereinander enthält.
func ContainsNRow(l []string, x string, n int) bool {
	y := 0

	for i := 0; i < len(l); i++ {

		if l[i] == x {
			y++
			if y == n {
				return true
			}
		} else {
			y = 0
		}

	}
	return false

}
