// Package chapter09 contains
// implementations of the algorithms introduced in Chapter 9.
package chapter09

// LzwDecompressor takes a sequence of indices generated by LzwCompressor
// and convert it back into the original text.
func LzwDecompressor(indices []int) string {

	decompressed := ""

	// Each element containing actual strings
	dictionary := make([]string, 256)

	// From dictionary[0] to dictionary[255] will contain
	// all the individual ASCII characters.
	for i := 0; i < 256; i++ {
		dictionary[i] = string(rune(i))
	}

	current := indices[0]
	indices = indices[1:]

	decompressed += dictionary[current]

	for len(indices) > 0 {
		previous := current

		current = indices[0]
		indices = indices[1:]

		// if dictionary[current] already exists, just output it into
		// decompressed
		if current < len(dictionary) {
			s := []rune(dictionary[current])

			decompressed += string(s)

			// previous would have been added to dictionary in LzwCompressor when
			// dictionary[previous]+string(s[0]) wasn't found in dictionary.
			// Since we've now seen dictionary[current],
			// it is the right time to append previous + current[0]
			// to the dictionary of LzwDecompressor.
			dictionary = append(dictionary, dictionary[previous]+string(s[0]))
		} else {
			// In some rare cases, current do not exist in our
			// recreated dictionary yet, because previous entries in indices
			// did not recover dictionary[current] just yet.
			// This situation seems to occur when current was added to indices
			// right after it was added to the dictionary by concatenating
			// dictionary[previous] + current[0], where current[0] has to be
			// same with previous[0].
			// So we deal with this special case by concatenating dictionary[previous]
			// with the first of character of the same string.
			previousEntry := []rune(dictionary[previous])

			s := dictionary[previous] + string(previousEntry[0])

			decompressed += s

			dictionary = append(dictionary, s)
		}
	}

	return decompressed
}
