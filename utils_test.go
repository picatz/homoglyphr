package homoglyphr

import "testing"

func TestContains(t *testing.T) {
	var ex = []string{"p", "i", "c", "a", "t"}
	got := contains(ex, "p")
	if got != true {
		t.Error("Unable to find string in slice of strings.")
	}
}

func TestRemoveDuplicates(t *testing.T) {
	var ex = []string{"p", "i", "p", "i"}
	var exr = removeDuplicates(ex)
	var pCounter = 0
	for _, c := range exr {
		if c == "p" {
			pCounter++
		}
	}
	if pCounter >= 2 {
		t.Error("Unable to successfully remove duplicates from slice of strings.")
	}
}
