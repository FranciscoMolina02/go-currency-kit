package round

type RoundingType string

const (
	Ceiling RoundingType = "ceiling"
	Floor   RoundingType = "floor"
	Nearest RoundingType = "nearest"
)

func (a RoundingType) String() string {
	return string(a)
}
