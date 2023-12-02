package patient

import (
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain"
)

const (
	CREATED   = "CREATED"
	CANCELLED = "CANCELLED"
)

type Patient struct {
	ID        uint64
	Name      string
	BirthDate *time.Time
	Avatar    string
	Sex       domain.Sex
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
