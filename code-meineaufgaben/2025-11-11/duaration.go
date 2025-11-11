package duration

type Duration int

func (z Duration) Sekunden() int {
	return int(z)
}
func (z Duration) Minuten() int {
	return z.Sekunden() / 60
}
func (z Duration) Stunden() int {
	return z.Minuten() / 60
}
func (z Duration) Tage() int {
	return z.Stunden() / 24
}
func (z Duration) Wochen() int {
	return z.Tage() / 7
}
func (z Duration) Jahre() int {
	return z.Tage() / 365
}

func FromSekunden(z int) Duration {
	return Duration(z)
}
func FromMinuten(z int) Duration {
	return Duration(z) * 60
}
func FromStunden(z int) Duration {
	return Duration(z) * 60 * 60
}
func FromTage(z int) Duration {
	return Duration(z) * 60 * 60 * 24
}
func FromWochen(z int) Duration {
	return Duration(z) * 60 * 60 * 24 * 7
}
func FromJahre(z int) Duration {
	return Duration(z) * 60 * 60 * 24 * 365
}
