package aufgabe6

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
func SymmetricDifference(l1, l2 []int) []int {
	test := false
	list := []int{}
	for i := 0; i < len(l1); i++ {
		for s := 0; s < len(l2); s++ {
			if l1[i] == l2[s] {
				test = true
			}

		}
		if !test {
			list = append(list, l1[i])
		}
		test = false
	}
	for i := 0; i < len(l2); i++ {
		for s := 0; s < len(l1); s++ {
			if l2[i] == l1[s] {
				test = true
			}

		}
		if !test {
			list = append(list, l2[i])
		}
		test = false
	}
	return list
}
