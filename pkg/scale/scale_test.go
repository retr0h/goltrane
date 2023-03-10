package scale

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetScalesByName", func() {
	BeforeEach(func() {})

	Context("initially", func() {
		It("should return a slice of scale names", func() {
			s := Scale{}
			got := s.GetScalesByName()
			Expect(len(got)).To(Equal(10))

			want := []string{
				"Major",
				"Minor",
				"Natural Minor",
				"Ionian",
				"Dorian",
				"Phrygian",
				"Lydian",
				"Mixolydian",
				"Aeolian",
				"Locrian",
			}
			Expect(want).To(Equal(got))
		})
	})
})

// ---
// scales:

//   C:
//     root: C
//     tones:
//       1: C
//       2: D
//       3: E
//       4: F
//       5: G
//       6: A
//       7: B

//   C major:
//     root: C
//     tones:
//       1: C
//       2: D
//       3: E
//       4: F
//       5: G
//       6: A
//       7: B

//   C minor:
//     root: C
//     tones:
//       1: C
//       2: D
//       3: Eb
//       4: F
//       5: G
//       6: Ab
//       7: Bb

//   C natural minor:
//     root: C
//     tones:
//       1: C
//       2: D
//       3: Eb
//       4: F
//       5: G
//       6: Ab
//       7: Bb

//   C diminished:
//     root: C
//     tones:
//       1: C
//       2: D
//       3: Eb
//       4: F
//       5: Gb
//       6: Ab
//       7: A
//       8: B

//   C augmented:
//     root: C
//     tones:
//       1: C
//       2: D#
//       3: E
//       4: G
//       5: G#
//       6: B

//   C melodic minor ascend:
//     root: C
//     tones:
//       1: C
//       2: D
//       3: Eb
//       4: F
//       5: G
//       6: A
//       7: B

//   C melodic minor descend:
//     root: C
//     tones:
//       1: C
//       2: D
//       3: Eb
//       4: F
//       5: G
//       6: Ab
//       7: Bb

//   C harmonic minor:
//     root: C
//     tones:
//       1: C
//       2: D
//       3: Eb
//       4: F
//       5: G
//       6: Ab
//       7: B

//   C ionian:
//     root: C
//     tones:
//       1: C
//       2: D
//       3: E
//       4: F
//       5: G
//       6: A
//       7: B

//   D dorian:
//     root: D
//     tones:
//       1: D
//       2: E
//       3: F
//       4: G
//       5: A
//       6: B
//       7: C

//   E phrygian:
//     root: E
//     tones:
//       1: E
//       2: F
//       3: G
//       4: A
//       5: B
//       6: C
//       7: D

//   F major:
//     root: F
//     tones:
//       1: F
//       2: G
//       3: A
//       4: Bb
//       5: C
//       6: D
//       7: E

//   F lydian:
//     root: F
//     tones:
//       1: F
//       2: G
//       3: A
//       4: B
//       5: C
//       6: D
//       7: E

//   G mixolydian:
//     root: G
//     tones:
//       1: G
//       2: A
//       3: B
//       4: C
//       5: D
//       6: E
//       7: F

//   A aeolian:
//     root: A
//     tones:
//       1: A
//       2: B
//       3: C
//       4: D
//       5: E
//       6: F
//       7: G

//   B locrian:
//     root: B
//     tones:
//       1: B
//       2: C
//       3: D
//       4: E
//       5: F
//       6: G
//       7: A
