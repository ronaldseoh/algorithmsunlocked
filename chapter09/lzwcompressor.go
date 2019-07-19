// Package chapter09 contains
// implementations of the algorithms introduced in Chapter 9.
package chapter09

func findDictionaryEntry(dictionary []string, term string) int {
	for i, v := range dictionary {
		if v == term {
			return i
		}
	}

	return -1
}

// LzwCompressor generates a LZW-compressed sequence of indices into a dictionary.
func LzwCompressor(text string) []int {

	// Each element containing actual strings
	dictionary := make([]string, 256)

	// From dictionary[0] to dictionary[255] will contain
	// all the individual ASCII characters.
	for i := 0; i < 256; i++ {
		dictionary[i] = string(rune(i))
	}

	// Store indices to the dictionary that corresponds to the original text
	indices := make([]int, 0)

	// Convert text into a slice of runes
	textArray := []rune(text)

	// Extract the first character
	s := string(textArray[0])
	textArray = textArray[1:]

	for len(textArray) > 0 {
		c := string(textArray[0])
		textArray = textArray[1:]

		if findDictionaryEntry(dictionary, s+c) != -1 {
			// If s+c is already in the dictionary, make s = s+c
			// so that we can extract one more character and try again
			// in the next iteration
			s = s + c
		} else {
			// if s+c isn't in the dictionary, let's represent what we have
			// in s (which should already exist in the dictionary) by its index
			indices = append(indices, findDictionaryEntry(dictionary, s))

			// Add s+c to the dictionary so that we can represent s+c with this new index
			// when it ever appears again in text
			// So the string to be added to dictionary is always the same as some string
			// already in the dictionary but extended by one character.
			dictionary = append(dictionary, s+c)

			// Since c of s+c was not represented in indices above, let's start building
			// up the string from c again.
			// Notice that at any point in this loop, s itself always exists in the dictionary.
			s = c
		}
	}

	// After the loop above ends, represent the remaining characters in s.
	indices = append(indices, findDictionaryEntry(dictionary, s))

	return indices
}
