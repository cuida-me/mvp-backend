package patient

import "time"

type Patient struct {
	ID             uint64
	Name           string
	DateBirth      *time.Time
	ProfilePicture string
	Sex            string
}
