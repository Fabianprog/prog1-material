package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MaxElements.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein!
*/

// MaxElements erwartet zwei int-Listen und liefert eine Liste, die an jeder Position
// jeweils das größere der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt, gilt dieses Element als das größere.
func MaxElements(l1, l2 []int) []int {
	liste := []int{}

	var länger int
	if len(l1) > len(l2) {
		länger = len(l1)
	} else {
		länger = len(l2)
	}
	for i := 0; i < länger; i++ {
		if l1[i] < l2[i] {
			liste = append(liste, l2[i])
		}
		if l2[i] < l1[i] {
			liste = append(liste, l1[i])
		}
		if l1[i] == l2[i] {
			liste = append(liste, l1[i])
		}
	}

	return liste
}
