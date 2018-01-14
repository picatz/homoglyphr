package homoglyphr

import "testing"

func TestRelatedAsciiCharacters(t *testing.T) {
	got := contains(RelatedAsciiCharacters("a"), "o")
	if got != true {
		t.Error("Unable to find related ASCII character")
	}
	got = contains(RelatedAsciiCharacters("b"), "d")
	if got != true {
		t.Error("Unable to find related ASCII character")
	}
}

func TestStreamRelatedAsciiCharacters(t *testing.T) {
	var counter = 0
	for _ = range StreamRelatedAsciiCharacters("d") {
		counter++
	}
	if counter != 2 {
		t.Error("Unable to stream related ASCII character")
	}
}

func TestStreamAllRelatedNonAsciiCharacters(t *testing.T) {
	var counter = 0
	for _ = range StreamAllRelatedNonAsciiCharacters("o") {
		counter++
	}
	if counter != 47 {
		//t.Error(counter)
		t.Error("Unable to stream all related non ascii character")
	}

}
