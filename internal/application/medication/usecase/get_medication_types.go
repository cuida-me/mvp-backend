package medication

import (
	"context"

	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type getMedicationTypesUseCase struct {
	repository medication.TypeRepository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewGetMedicationTypesUseCase(
	repository medication.TypeRepository,
	log log.Provider,
	apiErr apiErr.Provider,
) *getMedicationTypesUseCase {
	return &getMedicationTypesUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u getMedicationTypesUseCase) Execute(ctx context.Context) ([]*medication.Type, *apiErr.Message) {
	u.log.Info(ctx, "getting medication types")

	medicationTypes, err := u.repository.FindAllTypes(ctx)
	if err != nil {
		u.log.Error(ctx, "error to find medication types", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.NotFound(err)
	}

	return medicationTypes, nil
}
