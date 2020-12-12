package gender

type Type uint8

const (
	Female Type = iota
	Male
)

func (t Type) String() string {
	s := "unknown"
	switch t {
	case Female:
		s = "female"
	case Male:
		s = "male"
	}
	return s
}
