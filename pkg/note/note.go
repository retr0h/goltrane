package note

import (
	"strconv"
)

var (
	alphabet = []string{
		"A", "B", "C", "D", "E", "F", "G",
	}
	notes = [][]string{
		{"B#", "C", "Dbb"},
		{"B##", "C#", "Db"},
		{"C##", "D", "Ebb"},
		{"D#", "Eb", "Fbb"},
		{"D##", "E", "Fb"},
		{"E#", "F", "Gbb"},
		{"E##", "F#", "Gb"},
		{"F##", "G", "Abb"},
		{"G#", "Ab"},
		{"G##", "A", "Bbb"},
		{"A#", "Bb", "Cbb"},
		{"A##", "B", "Cb"},
	}
	// standard names for each note, ordered identical to the indexes in the
	// notes array.
	intervals = [][]string{
		{"P1", "d2"}, // Perfect unison   Diminished second
		{"m2", "A1"}, // Minor second     Augmented unison
		{"M2", "d3"}, // Major second     Diminished third
		{"m3", "A2"}, // Minor third      Augmented second
		{"M3", "d4"}, // Major third      Diminished fourth
		{"P4", "A3"}, // Perfect fourth   Augmented third
		{"d5", "A4"}, // Diminished fifth Augmented fourth
		{"P5", "d6"}, // Perfect fifth    Diminished sixth
		{"m6", "A5"}, // Minor sixth      Augmented fifth
		{"M6", "d7"}, // Major sixth      Diminished seventh
		{"m7", "A6"}, // Minor seventh    Augmented sixth
		{"M7", "d8"}, // Major seventh    Diminished octave
		{"P8", "A7"}, // Perfect octave   Augmented seventh
	}
)

// findAlphabetIndex takes a sequence of notes (alphabet), and a note to
// search for, and returns the index
func findAlphabetIndex(alphabet []string, searchNote string) int {
	for i, item := range alphabet {
		if item == searchNote {
			return i
		}
	}

	return -1
}

// findAlphabetIndex takes a sequence of notes (notes), and a note to
// search for, and returns the index
func findNoteIndex(notes [][]string, searchNote string) int {
	for i, note := range notes {
		for _, item := range note {
			if item == searchNote {
				return i
			}
		}
	}

	return -1
}

// leftRotateAlphabet left-rotate a sequence of notes (alphabet) by n positions
func leftRotateAlphabet(alphabet []string, rotation int) []string {
	var newArray []string
	size := len(alphabet)

	for i := 0; i < rotation; i++ {
		newArray = alphabet[1:size]
		newArray = append(newArray, alphabet[0])
		alphabet = newArray
	}

	return alphabet
}

// leftRotateNotes left-rotate a sequence of notes (notes) by n positions
func leftRotateNotes(notes [][]string, rotation int) [][]string {
	var newArray [][]string
	size := len(notes)

	for i := 0; i < rotation; i++ {
		newArray = notes[1:size]
		newArray = append(newArray, notes[0])
		notes = newArray
	}

	return notes
}

// chromaticScale generate a chromatic scale in a given key
func chromaticScale(key string) [][]string {
	// Determine how much to rotate the notes and return the rotated version
	numRotations := findNoteIndex(notes, key)

	return leftRotateNotes(notes, numRotations)
}

// makeIntervalsStandard generate a map of all interval names to the
// right note names in a given key
func makeIntervalsStandard(key string) map[string]string {
	var note string
	labels := make(map[string]string)

	// Step 1: Generate a chromatic scale in our desired key
	cs := chromaticScale(key)

	// The alphabets starting at provided key's alphabet
	numRotations := findAlphabetIndex(alphabet, key)
	alphabetKey := leftRotateAlphabet(alphabet, numRotations)

	// Step 2: Find the notes to search through based on degree
	for i, intervalList := range intervals {
		notesToSearch := cs[i%len(cs)]

		for _, intervalName := range intervalList {
			sz := len(intervalName)
			lastChar, _ := strconv.Atoi(intervalName[sz-1:])

			// Get the interval degree
			degree := int(lastChar - 1) //  e.g. M3 --> 3, m7 --> 7

			// Get the alphabet to look for
			alphabetToSearch := alphabetKey[degree%len(alphabetKey)]

			for _, x := range notesToSearch {
				if x[0:1] == alphabetToSearch {
					note = x
				}
			}

			if note == "" {
				note = notesToSearch[0]
			}

			labels[intervalName] = note
		}
	}

	// fmt.Println(labels)
	return labels
}
