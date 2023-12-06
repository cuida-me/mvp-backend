package domain

type Sex uint8

type DailyOfWeek uint8

const (
	MALE Sex = iota
	FEMALE
)

const (
	DOMINGO DailyOfWeek = iota
	SEGUNDA
	TERCA
	QUARTA
	QUINTA
	SEXTA
	SABADO
)

func (s DailyOfWeek) String() string {
	switch s {
	case DOMINGO:
		return "domingo"
	case SEGUNDA:
		return "segunda"
	case TERCA:
		return "terca"
	case QUARTA:
		return "quarta"
	case QUINTA:
		return "quinta"
	case SEXTA:
		return "sexta"
	case SABADO:
		return "sabado"
	}
	return "unknown"
}

func (s Sex) String() *string {
	male := "male"
	female := "female"
	switch s {
	case MALE:
		return &male
	case FEMALE:
		return &female
	}
	return nil
}
