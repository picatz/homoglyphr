package homoglyphr

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsRunes(s [][]rune, e rune) bool {
	for _, a := range s {
		for _, r := range a {
			if r == e {
				return true
			}
		}
	}
	return false
}

func removeDuplicates(s []string) []string {
	reducedParts := []string{}
	for _, char := range s {
		if contains(reducedParts, char) {
			continue
		} else {
			reducedParts = append(reducedParts, char)
		}
	}
	return reducedParts
}
