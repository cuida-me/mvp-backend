package domain

type Sex uint8

const (
	UNDEFINED Sex = iota
	MALE
	FEMALE
)

func (s Sex) String() string {
	switch s {
	case UNDEFINED:
		return "undefined"
	case MALE:
		return "male"
	case FEMALE:
		return "female"
	}
	return "unknown"
}
