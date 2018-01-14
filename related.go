package homoglyphr

var NoRelatedCharacters = []string{}

func RelatedAsciiCharacters(str string) []string {
	switch str {
	case "a":
		return []string{"o"}
	case "b":
		return []string{"d", "lo"}
	case "d":
		return []string{"cl", "c1"}
	case "g":
		return []string{"q"}
	case "i":
		return []string{"l"}
	case "j":
		return []string{"i"}
	case "l":
		return []string{"1"}
	case "m":
		return []string{"nn", "rn"}
	case "o":
		return []string{"0"}
	case "p":
		return []string{"q"}
	case "q":
		return []string{"p"}
	case "u":
		return []string{"v"}
	case "w":
		return []string{"vv"}
	case "y":
		return []string{"v"}
	case "1":
		return []string{"l"}
	case "0":
		return []string{"o"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedAsciiCharacters(domain string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedAsciiCharacters(domain)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedGurmukhiCharacters(r rune) []string {
	switch r {
	case 'a':
		return []string{"\u0A30"}
	case 'o':
		return []string{"\u0A66"}
	case 'g':
		return []string{"\u0A67"}
	case 'q':
		return []string{"\u0A67"}
	case 'u':
		return []string{"\u0A2A", "\u0A6B"}
	case 's':
		return []string{"\u0A1F"}
	case '3':
		return []string{"\u0A69", "\u0A24"}
	case '5':
		return []string{"\u0A1F"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedGurmukhiCharacters(c string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedGujaratiCharacters(c)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedArabicCharacters(c string) []string {
	switch string(c) {
	case "o":
		return []string{"\u0665", "\u06C1", "\u06C2", "\u06D5"}
	case "v":
		return []string{"\u0667", "\u06F7"}
	case "j":
		return []string{"\uFEDE"}
	case "q":
		return []string{"\u0669", "\u08B1"}
	case "l":
		return []string{"\u08AD"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedArabicCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedArabicCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedArmenianCharacters(str string) []string {
	switch str {
	case "a":
		return []string{"\u0571"}
	case "d":
		return []string{"\u056E"}
	case "p":
		return []string{"\u0562", "\u0584", "\u0569"}
	case "q":
		return []string{"\u0563", "\u0566"}
	case "g":
		return []string{"\u0566"}
	case "t":
		return []string{"\u0567"}
	case "l":
		return []string{"\u056C"}
	case "o":
		return []string{"\u0585"}
	case "n":
		return []string{"\u0578", "\u057C"}
	case "u":
		return []string{"\u057D"}
	case "w":
		return []string{"\u0561"}
	case "4":
		return []string{"\u056F"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedArmenianCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedArmenianCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedBengaliCharacters(str string) []string {
	switch str {
	case "e":
		return []string{"\u09CE"}
	case "l":
		return []string{"\u09F7"}
	case "b":
		return []string{"\u09EE"}
	case "q":
		return []string{"\u09ED", "\u0980"}
	case "o":
		return []string{"\u09E6"}
	case "8":
		return []string{"\u09EA"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedBengaliCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedBengaliCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedBopomofoCharacters(str string) []string {
	switch str {
	case "l":
		return []string{"\u31B3"}
	case "u":
		return []string{"\u3129"}
	case "n":
		return []string{"\u3107"}
	case "m":
		return []string{"\u31AC"}
	case "4":
		return []string{"\u3110", "\u31A2"}
	case "p":
		return []string{"\u3117"}
	case "y":
		return []string{"\u311A"}
	case "x":
		return []string{"\u3128"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedBopomofoCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedBopomofoCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedCanadianAboriginalCharacters(str string) []string {
	switch str {
	case "s":
		return []string{"\u151D", "\u151E", "\u151F",
			"\u1520", "\u1521", "\u1522", "\u1523", "\u1524"}
	case "r":
		return []string{"\u148B"}
	case "b":
		return []string{"\u1472", "\u1473", "\u147F",
			"\u1480", "\u1481", "\u1482", "\u1488",
			"\u1579", "\u157A", "\u15AF", "\u1624",
			"\u1625", "\u1626", "\u162A", "\u162B",
			"\u162C"}
	case "d":
		return []string{"\u146F", "\u1470", "\u1471", "\u147A",
			"\u147B", "\u147C", "\u147D", "\u1577", "\u1578"}
	case "p":
		return []string{"\u146C", "\u146D", "\u146E",
			"\u1476", "\u1477", "\u1478", "\u1479"}
	case "q":
		return []string{"\u146B", "\u1474", "\u1475", "\u1574"}
	case "x":
		return []string{"\u166E"}
	case "c":
		return []string{"\u1455"}
	case "n":
		return []string{"\u144E"}
	case "u":
		return []string{"\u144C", "\u1640"}
	case "v":
		return []string{"\u1401"}
	case "j":
		return []string{"\u148E", "\u148F", "\u1491", "\u1490", "\u148D"}
	case "a":
		return []string{"\u1610", "\u1611", "\u1612", "\u161C"}
	case "e":
		return []string{"\u1566", "\u1567", "\u1568", "\u1569", "\u1613"}
	case "3":
		return []string{"\u15F1", "\u15F2", "	\u15F3", "\u15F4"}
	default:
		return []string{} // none
	}
}

func StreamRelatedCanadianAboriginalCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedCanadianAboriginalCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedCherokeeCharacters(str string) []string {
	switch str {
	case "i":
		return []string{"\uAB75"}
	case "t":
		return []string{"\uAB80"}
	case "r":
		return []string{"\uAB81"}
	case "w":
		return []string{"\uAB83"}
	case "d":
		return []string{"\uAB84", "\uABB7"}
	case "o":
		return []string{"\uAB8E", "\uABBB"}
	case "h":
		return []string{"\uAB92", "\uABB5"}
	case "g":
		return []string{"\uABBD"}
	case "k":
		return []string{"\uABB6"}
	case "n":
		return []string{"\uAB91"}
	case "z":
		return []string{"\uAB93"}
	case "3":
		return []string{"\uAB9B"}
	case "4":
		return []string{"\uAB9E"}
	case "b":
		return []string{"\uAB9F"}
	case "l":
		return []string{"\uABA3"}
	case "v":
		return []string{"\uABA9"}
	case "s":
		return []string{"\uABAA"}
	case "c":
		return []string{"\uABAF"}
	case "p":
		return []string{"\uABB2"}
	case "9":
		return []string{"\uABBD"}
	case "6":
		return []string{"\uABBE"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedCherokeeCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedCherokeeCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedCopticCharacters(str string) []string {
	switch str {
	case "t":
		return []string{"\u03EF"}
	case "6":
		return []string{"\u03ED"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedCopticCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedCopticCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedEthiopicCharacters(str string) []string {
	switch str {
	case "g":
		return []string{"\u12EB"}
	case "u":
		return []string{"\u1200", "\u1201"}
	case "y":
		return []string{"\u1202", "\u1203", "\u1204"}
	case "b":
		return []string{"\u122A"}
	case "w":
		return []string{"\u1220", "\u1221"}
	case "o":
		return []string{"\u12D0", "\u12D1", "\u12D5",
			"\u1340", "\u1345"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedEthiopicCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedEthiopicCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedGreekCharacters(str string) []string {
	switch str {
	case "a":
		return []string{"\u03B1", "\u1F00", "\u1F01", "\u1F70", "\u1FB1"}
	case "y":
		return []string{"\u03B3"}
	case "d":
		return []string{"\u03B4"}
	case "o":
		return []string{"\u03B8", "\u03BF", "\u03C3", "\u03CC",
			"\u1F40", "\u1F41"}
	case "i":
		return []string{"\u03B9", "\u1F76", "\u1F77", "\u0390", "\u03CA"}
	case "u":
		return []string{"\u03BC", "\u03BD", "\u03C5", "\u03CB",
			"\u03CD", "\u1F50"}
	case "v":
		return []string{"\u03BD"}
	case "w":
		return []string{"\u03C9", "\u03CE", "\u1F60", "\u1F61"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedGreekCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedGreekCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedGujaratiCharacters(str string) []string {
	switch str {
	case "o":
		return []string{"\u0AE6"}
	case "s":
		return []string{"\u0A95", "\u0A99", "\u0A9F", "\u0AA1"}
	case "q":
		return []string{"\u0AE7"}
	case "3":
		return []string{"\u0AE9"}
	case "4":
		return []string{"\u0AEB", "\u0AAA"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedGujaratiCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedGujaratiCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedHanCharacters(str string) []string {
	switch str {
	case "e":
		return []string{"\u2E8B", "\u2E92"}
	case "t":
		return []string{"\u2E8A"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedHanCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedHanCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedHangulCharacters(str string) []string {
	switch str {
	case "o":
		return []string{"\u110B", "\uFFB7"}
	case "-":
		return []string{"\uFFDA"}
	case "l":
		return []string{"\uFFDC"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedHangulCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedHangulCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedHebrewCharacters(str string) []string {
	switch str {
	case "o":
		return []string{"\u05E1", "\uFB41"}
	case "c":
		return []string{"\u05DB"}
	case "7":
		return []string{"\u05DA"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedHebrewCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedHebrewCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedCyrillicCharacters(str string) []string {
	switch str {
	case "a":
		return []string{"\u0430", "\u04D1", "\u04D3", "\u04D8", "\u04D7"}
	case "b":
		return []string{"\u044C", "\u048D"}
	case "bi":
		return []string{"\u044B"}
	case "ae":
		return []string{"\u04D5"}
	case "e":
		return []string{"\u04D7", "\u0435", "\u0450", "\u0451", "\u04BD", "\u04BF"}
	case "c":
		return []string{"\u0441", "\u043E"}
	case "d":
		return []string{"\u0501"}
	case "h":
		return []string{"\u0452", "\u045B", "\u04BB"}
	case "i":
		return []string{"\u0456", "\u0457"}
	case "j":
		return []string{"\u0458"}
	case "k":
		return []string{"\u043A", "\u045C", "\u049B", "\u049D", "\u049F", "\u04C4"}
	case "l":
		return []string{"\u04CF"}
	case "m":
		return []string{"\u04CE", "\u043C"}
	case "n":
		return []string{"\u043F", "\u04C6", "\u0525"}
	case "o":
		return []string{"\u0473", "\u047B", "\u04E7", "\u04E9", "\u04EB", "\u043E"}
	case "p":
		return []string{"\u048F", "\u0440"}
	case "r":
		return []string{"\u0433", "\u0491", "\u0493"}
	case "s":
		return []string{"\u0455"}
	case "v":
		return []string{"\u0475", "\u0477"}
	case "u":
		return []string{"\u0446"}
	case "w":
		return []string{"\u0448", "\u0449", "\u0461", "\u047F", "\u051D"}
	case "y":
		return []string{"\u0447", "\u0443", "\u045E", "\u04EF", "\u04F1", "\u04F3", "\u04B7", "\u04B9", "\u04F5"}
	case "x":
		return []string{"\u0445", "\u04B3", "\u04FF"}
	case "1":
		return []string{"\u04CF"}
	case "3":
		return []string{"\u0437", "\u044D", "\u0454", "\u0499", "\u04DF", "\u04E1"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedCyrillicCharacters(str string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedCyrillicCharacters(str)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func RelatedLatinCharacters(str string) []string {
	switch str {
	case "a":
		return []string{
			"\u0061", "\u00E0", "\u00E1", "\u00E2",
			"\u00E3", "\u00E4", "\u00E5", "\u0101",
			"\u0103", "\u0105", "\u01CE", "\u0201",
			"\u0203", "\u0227", "\u1E01", "\u1E9A",
			"\u1EA1", "\u1EA3"}
	case "b":
		return []string{"\u1E03", "\u1E05", "\u1E07"}
	case "c":
		return []string{"\u00E7", "\u0107", "\u0109", "\u010B", "\u010D"}
	case "d":
		return []string{"\u010F", "\u1E0B", "\u1E0D", "\u1E0F", "\u1E11", "\u1E13"}
	case "e":
		return []string{
			"\u00E8", "\u00E9", "\u00EA", "\u00EB",
			"\u0113", "\u0115", "\u0117", "\u0119",
			"\u011B", "\u0205", "\u0207", "\u0229",
			"\u1E1B", "\u1EB9", "\u1EBB", "\u1EBD"}
	case "f":
		return []string{"\u1E1F"}
	case "g":
		return []string{"\u011D", "\u011F", "\u0121", "\u0123", "\u01E7", "\u01F5", "\u1E21"}
	case "h":
		return []string{"\u0125", "\u021F", "\u1E23", "\u1E25", "\u1E27", "\u1E29", "\u1E2B", "\u1E96"}
	case "i":
		return []string{
			"\u00EC", "\u00ED", "\u00EE", "\u00EF",
			"\u0129", "\u012B", "\u012D", "\u012F",
			"\u01D0", "\u0209", "\u020B", "\u1E2D",
			"\u1EC9", "\u1ECB"}
	case "j":
		return []string{"\u0135", "\u01F0"}
	case "k":
		return []string{"\u0137", "\u01E9", "\u1E31", "\u1E33", "\u1E35"}
	case "l":
		return []string{"\u013A", "\u013C", "\u013E", "\u0140", "\u1E37", "\u1E3B"}
	case "m":
		return []string{"\u1E3F", "\u1E41", "\u1E43"}
	case "n":
		return []string{
			"\u00F1", "\u0144", "\u0146", "\u0148",
			"\u0149", "\u01F9", "\u1E45", "\u1E47",
			"\u1E49", "\u1E4B"}
	case "o":
		return []string{
			"\u00F2", "\u00F3", "\u00F4", "\u00F5",
			"\u00F6", "\u014D", "\u014F", "\u0151",
			"\u01A1", "\u01D2", "\u01EB", "\u020D",
			"\u020F", "\u022F", "\u1ECD", "\u1ECF"}
	case "p":
		return []string{"\u1E55", "\u1E57"}
	case "q":
		return []string{"\u1E55", "\u1E57"}
	case "r":
		return []string{
			"\u0155", "\u0157", "\u0159", "\u0211",
			"\u0213", "\u1E59", "\u1E5B", "\u1E5F"}
	case "s":
		return []string{
			"\u015B", "\u015D", "\u015F", "\u0161",
			"\u0219", "\u1E61", "\u1E63"}
	case "t":
		return []string{
			"\u0163", "\u0165", "\u021B", "\u1E6B",
			"\u1E6D", "\u1E6F", "\u1E71", "\u1E97"}
	case "u":
		return []string{
			"\u00F9", "\u00FA", "\u00FB", "\u00FC",
			"\u0169", "\u016B", "\u016D", "\u016F",
			"\u0171", "\u0173", "\u01B0", "\u01D4",
			"\u0215", "\u0217", "\u1E73", "\u1E75",
			"\u1E77", "\u1EE5", "\u1EE7"}
	case "v":
		return []string{"\u1E7D", "\u1E7F"}
	case "w":
		return []string{"\u0175", "\u1E81", "\u1E83", "\u1E85", "\u1E87", "\u1E89", "\u1E98"}
	case "x":
		return []string{"\u1E8B", "\u1E8D"}
	case "y":
		return []string{
			"\u1EF5", "\u1EF7", "\u1EF9", "\u00FD",
			"\u00FF", "\u0177", "\u0233", "\u1E8F",
			"\u1E99", "\u1EF3"}
	case "z":
		return []string{"\u017A", "\u017C", "\u017E", "\u1E91", "\u1E93", "\u1E95"}
	default:
		return NoRelatedCharacters
	}
}

func StreamRelatedLatinCharacters(domain string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		data := RelatedLatinCharacters(domain)
		for _, d := range data {
			messages <- d
		}
	}()
	return messages
}

func StreamAllRelatedCharacters(domain string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		for d := range StreamRelatedAsciiCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedGurmukhiCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedArabicCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedArmenianCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedBengaliCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedBopomofoCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedCanadianAboriginalCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedCherokeeCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedCopticCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedEthiopicCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedGreekCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedGujaratiCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedHanCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedHangulCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedHebrewCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedCyrillicCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedLatinCharacters(domain) {
			messages <- d
		}
	}()
	return messages
}

func StreamAllRelatedNonAsciiCharacters(domain string) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		for d := range StreamRelatedGurmukhiCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedArabicCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedArmenianCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedBengaliCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedBopomofoCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedCanadianAboriginalCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedCherokeeCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedCopticCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedEthiopicCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedGreekCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedGujaratiCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedHanCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedHangulCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedHebrewCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedCyrillicCharacters(domain) {
			messages <- d
		}
		for d := range StreamRelatedLatinCharacters(domain) {
			messages <- d
		}
	}()
	return messages
}
