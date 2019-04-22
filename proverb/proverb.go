package proverb

// Proverb takes a list of inputs and generates relevant proverbs.
func Proverb(rhyme []string) []string {
	proverbs := make([]string, 0)

	for i, v := range rhyme {
		if i+1 >= len(rhyme) {
			proverbs = append(proverbs,
				"And all for the want of a "+rhyme[0]+".")
			break
		}
		proverbs = append(proverbs,
			"For want of a "+v+" the "+rhyme[i+1]+" was lost.")
	}

	return proverbs
}
