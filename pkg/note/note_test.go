package note

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("findAlphabetIndex", func() {
	DescribeTable("find index of given note",
		func(alphabet []string, note string, want int) {
			got := findAlphabetIndex(alphabet, note)

			Expect(got).To(Equal(want))
		},
		Entry("When alphabet contains the note 'A'", alphabet, "A", 0),
		Entry("When note 'inv' is not found in alphabet", alphabet, "inv", -1),
	)
})

var _ = Describe("findNoteIndex", func() {
	DescribeTable("find index of given note",
		func(alphabet []string, note string, want int) {
			got := findNoteIndex(notes, note)

			Expect(got).To(Equal(want))
		},
		Entry("When notes contains the note 'A'", alphabet, "A", 9),
		Entry("When note 'inv' is not found in notes", alphabet, "inv", -1),
	)
})

var _ = Describe("leftRotateAlphabet", func() {
	DescribeTable("rotate left by a given position",
		func(alphabet []string, rotation int, want []string) {
			got := leftRotateAlphabet(alphabet, rotation)

			Expect(got).To(Equal(want))
		},
		Entry("When provided an array",
			[]string{"A", "B", "C", "D", "E", "F", "G"},
			3,
			[]string{"D", "E", "F", "G", "A", "B", "C"},
		),
	)
})

var _ = Describe("leftRotateNotes", func() {
	DescribeTable("rotate left by a given position",
		func(notes [][]string, rotation int, want [][]string) {
			got := leftRotateNotes(notes, rotation)

			Expect(got).To(Equal(want))
		},
		Entry("When provided an array of arrays",
			[][]string{{"A", "B"}, {"B", "C"}, {"D", "E"}},
			2,
			[][]string{{"D", "E"}, {"A", "B"}, {"B", "C"}},
		),
	)
})

var _ = Describe("chromaticScale", func() {
	Context("given key", func() {
		It("should generate scale by rotating notes", func() {
			got := chromaticScale("D")
			want := [][]string{
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
				{"B#", "C", "Dbb"},
				{"B##", "C#", "Db"},
			}

			Expect(got).To(Equal(want))
		})
	})
})

var _ = Describe("makeIntervalsStandard", func() {
	Context("given key", func() {
		It("should generate a map of all interval names", func() {
			got := makeIntervalsStandard("C")
			want := map[string]string{
				"P1": "C",
				"d2": "Dbb",
				"m2": "Db",
				"A1": "C#",
				"M2": "D",
				"d3": "Ebb",
				"m3": "Eb",
				"A2": "D#",
				"M3": "E",
				"d4": "Fb",
				"P4": "F",
				"A3": "E#",
				"d5": "Gb",
				"A4": "F#",
				"P5": "G",
				"d6": "Abb",
				"m6": "Ab",
				"A5": "G#",
				"M6": "A",
				"d7": "Bbb",
				"m7": "Bb",
				"A6": "A#",
				"M7": "B",
				"d8": "Cb",
				"P8": "C",
				"A7": "B#",
			}

			Expect(got).To(Equal(want))
		})
	})
})
