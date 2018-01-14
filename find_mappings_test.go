package homoglyphr

import "testing"

//TODO: Clean up this test lol
func TestFindAllMappings(t *testing.T) {
	var mappings = FindAllMappings("google")
	var counter = 0
	for _ = range mappings {
		counter++
	}
	if counter != 4 {
		t.Error("Unable to find all related character mappings")
	}
	counter = 0
	for _ = range mappings["g"] {
		counter++
	}
	if counter != 11 {
		t.Error("Unable to find correct number of related character mappings for this 'fella 'g'")
	}
	counter = 0
	for _ = range mappings["o"] {
		counter++
	}
	if counter != 44 {
		t.Error("Unable to find correct number of related character mappings for this 'fella 'o'")
	}
	counter = 0
	for _ = range mappings["l"] {
		counter++
	}
	if counter != 13 {
		t.Error("Unable to find correct number of related character mappings for this 'fella 'l'")
	}
	counter = 0
	for _ = range mappings["e"] {
		counter++
	}
	if counter != 30 {
		t.Error("Unable to find correct number of related character mappings for this 'fella 'e'")
	}
}
