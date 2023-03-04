package scale

type Scale struct{}

type ModeIntervals []int

type Mode struct {
	Name string
	set  ModeIntervals
}

type List []string
