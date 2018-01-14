package homoglyphr

import "strings"

func FindAllMappings(str string) map[string][]string {
	parts := strings.Split(str, "")
	reducedParts := removeDuplicates(parts)
	similarsMapping := make(map[string][]string)
	for _, part := range reducedParts {
		similarsMapping[part] = []string{}
		for c := range StreamRelatedAsciiCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedLatinCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedCyrillicCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedGurmukhiCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedArmenianCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedBengaliCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedBopomofoCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedCanadianAboriginalCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedCherokeeCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedCopticCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedEthiopicCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedGreekCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedGujaratiCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedHanCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedHangulCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedHebrewCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
	}
	return similarsMapping
}

func FindNonAsciiMappings(str string) map[string][]string {
	parts := strings.Split(str, "")
	reducedParts := removeDuplicates(parts)
	similarsMapping := make(map[string][]string)
	for _, part := range reducedParts {
		similarsMapping[part] = []string{}
		for c := range StreamRelatedLatinCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedCyrillicCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedGurmukhiCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedArmenianCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedBengaliCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedBopomofoCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedCanadianAboriginalCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedCherokeeCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedCopticCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedEthiopicCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedGreekCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedGujaratiCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedHanCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedHangulCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
		for c := range StreamRelatedHebrewCharacters(part) {
			similarsMapping[part] = append(similarsMapping[part], c)
		}
	}
	return similarsMapping
}
