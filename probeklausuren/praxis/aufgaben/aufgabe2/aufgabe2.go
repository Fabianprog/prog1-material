package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	erste := -1
	letzte := -1
	Liste := []string{}
	for i := 0; i < len(list); i++ {
		if list[i] == first {
			erste = i
		}
		if list[i] == last {
			letzte = i
		}

	}
	if erste == -1 {
		return []string{}
	}
	if letzte == -1 {
		return []string{}
	}
	if letzte < erste {
		return []string{}
	}
	for s := 0; s < len(list); s++ {
		if s < erste || s > letzte {
			Liste = append(Liste, list[s])
		}
	}
	return Liste

}
