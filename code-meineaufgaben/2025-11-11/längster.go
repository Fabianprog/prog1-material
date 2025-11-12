package duration

func Längster(a []string) string {
	Längster := 0
	Länge := 0

	for i := 0; i < len(a); i++ {
		if Länge < len(a[i]) {
			Länge = len(a[i])
			Längster = i
		}
	}
	return a[Längster]
}
