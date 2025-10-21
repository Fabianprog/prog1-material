package listprops

// ContainsOnly liefert true, falls die Liste l
// ausschließlich den String x enthält.
func ContainsOnly(l []string, x string) bool {
	for i := 0; i < len(l); i++ {
		if l[i] != x {
			return false
		}
	}
	return true
}
