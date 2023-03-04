package scale

var (
	// W-W-H-W-W-W-H
	// 2 2 1 2 2 2 1
	ionianIntervals = ModeIntervals{2, 2, 1, 2, 2, 2}
	// W-H-W-W-W-H-W
	// 2 1 2 2 2 1 2
	dorianIntervals = ModeIntervals{2, 1, 2, 2, 2, 1}
	// H-W-W-W-H-W-W
	// 1 2 2 2 1 2 2
	phrygianIntervals = ModeIntervals{1, 2, 2, 2, 1, 2}
	// W-W-W-H-W-W-H
	// 2 2 2 1 2 2 1
	lydianIntervals = ModeIntervals{2, 2, 2, 1, 2, 2}
	// W-W-H-W-W-H-W
	// 2 2 1 2 2 1 2
	mixolydianIntervals = ModeIntervals{2, 2, 1, 2, 2, 1}
	// natural minor
	// W-H-W-W-H-W-W
	// 2 1 2 2 1 2 2
	aeolianIntervals = ModeIntervals{2, 1, 2, 2, 1, 2}
	// H-W-W-H-W-W-W
	// 1 2 2 1 2 2 2
	locrianIntervals = ModeIntervals{1, 2, 2, 1, 2, 2}
)

// Modal Scales
var modes = []Mode{
	{
		Name: "Major",
		set:  ionianIntervals,
	},
	{
		Name: "Minor",
		set:  aeolianIntervals,
	},
	{
		Name: "Natural Minor",
		set:  aeolianIntervals,
	},
	{
		Name: "Ionian",
		set:  ionianIntervals,
	},
	{
		Name: "Dorian",
		set:  dorianIntervals,
	},
	{
		Name: "Phrygian",
		set:  phrygianIntervals,
	},
	{
		Name: "Lydian",
		set:  lydianIntervals,
	},
	{
		Name: "Mixolydian",
		set:  mixolydianIntervals,
	},
	{
		Name: "Aeolian",
		set:  aeolianIntervals,
	},
	{
		Name: "Locrian",
		set:  locrianIntervals,
	},
}

// NewScale creates an instance of Scale.
func NewScale() *Scale {
	return &Scale{}
}

func (s *Scale) GetScalesByName() []string {
	scales := s.GetScales()
	scaleModeList := make([]string, 0, len(scales))

	for _, f := range modes {
		scaleModeList = append(scaleModeList, f.Name)
	}

	return scaleModeList
}

func (s *Scale) GetScales() []Mode {
	return modes
}
