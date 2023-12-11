package usecase

import (
	"context"
	"sort"
	"time"

	dto "github.com/cuida-me/mvp-backend/internal/application/scheduling/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

var (
	colors = []string{
		"#FD8247",
		"#296147",
		"#EB7287",
		"#F9C80E",
		"#F86624",
		"#662E9B",
		"#43BCCD",
		"#F9F871",
	}

	weekDays = []string{
		"Dom",
		"Seg",
		"Ter",
		"Qua",
		"Qui",
		"Sex",
		"Sáb",
	}

	months = []string{
		"Janeiro",
		"Fevereiro",
		"Março",
		"Abril",
		"Maio",
		"Junho",
		"Julho",
		"Agosto",
		"Setembro",
		"Outubro",
		"Novembro",
		"Dezembro",
	}
)

type getWeekSchedulingUseCase struct {
	repository           scheduling.Repository
	medicationRepository medication.Repository
	log                  log.Provider
	apiErr               apiErr.Provider
}

func NewGetWeekSchedulingUseCase(
	repository scheduling.Repository,
	medicationRepository medication.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *getWeekSchedulingUseCase {
	return &getWeekSchedulingUseCase{
		repository:           repository,
		medicationRepository: medicationRepository,
		log:                  log,
		apiErr:               apiErr,
	}
}

func (u getWeekSchedulingUseCase) Execute(ctx context.Context, patientID *uint64) ([]*dto.DailyScheduling, *apiErr.Message) {
	u.log.Info(ctx, "get week scheduling", log.Body{
		"patientID": patientID,
	})

	patientMedication, err := u.medicationRepository.FindAllMedicationByPatientID(ctx, patientID)
	if err != nil {
		u.log.Error(ctx, "error to get patient medication", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	response := make([]*dto.DailyScheduling, 0)
	sunday, saturday := u.getRangeOfScheduling()

	u.generateDailyGroupForWeek(&response, sunday, saturday)

	for i, medication := range patientMedication {
		scheduling, err := u.repository.FindSchedulingByMedicationIDAndDateRange(ctx, &medication.ID, sunday, saturday)
		if err != nil {
			u.log.Error(ctx, "error to get week scheduling", log.Body{
				"error": err.Error(),
			})
			return nil, u.apiErr.InternalServerError(err)
		}

		for _, schedule := range scheduling {
			groupIndex := u.getDailyGroup(&response, *schedule.MedicationTime)

			response[groupIndex].Scheduling = append(response[groupIndex].Scheduling, *mapToScheduling(schedule, medication, colors[i]))
		}

	}

	for _, day := range response {
		sort.Sort(response[0])
		dayColors := make([]string, 0)
		for _, schedule := range day.Scheduling {
			if len(dayColors) < 3 {
				dayColors = append(dayColors, schedule.Color)
			}
		}
		day.DayColors = dayColors
	}

	return response, nil
}

func (u getWeekSchedulingUseCase) generateDailyGroupForWeek(response *[]*dto.DailyScheduling, sunday time.Time, saturday time.Time) {
	for sunday.Before(saturday) {
		u.getDailyGroup(response, sunday)
		sunday = sunday.AddDate(0, 0, 1)
	}
}

func (u getWeekSchedulingUseCase) getDailyGroup(response *[]*dto.DailyScheduling, dateOfMedication time.Time) int {
	for i, day := range *response {
		if day.DayWeek == int(dateOfMedication.Weekday()) {
			return i
		}
	}

	scheduling := make([]dto.Scheduling, 0)

	var dailyGroup dto.DailyScheduling
	dailyGroup.Day = dateOfMedication.Day()
	dailyGroup.DayName = weekDays[int(dateOfMedication.Weekday())]
	dailyGroup.Date = dateOfMedication.In(time.FixedZone("UTC-3", -3*3600))
	dailyGroup.DayWeek = int(dateOfMedication.Weekday())
	dailyGroup.MonthName = months[int(dateOfMedication.Month())-1]
	dailyGroup.Scheduling = scheduling

	*response = append(*response, &dailyGroup)

	return len(*response) - 1
}

func (u getWeekSchedulingUseCase) getRangeOfScheduling() (time.Time, time.Time) {
	now := time.Now().In(time.FixedZone("UTC-3", -3*3600))

	sunday := time.Date(now.Year(), now.Month(), now.Day(), 0o0, 0o0, 0o0, 0, time.FixedZone("UTC-3", -3*3600))
	saturday := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.FixedZone("UTC-3", -3*3600))

	for sunday.Weekday() != time.Sunday {
		sunday = sunday.AddDate(0, 0, -1)
	}

	for saturday.Weekday() != time.Saturday {
		saturday = saturday.AddDate(0, 0, 1)
	}

	return sunday, saturday
}

func mapToScheduling(scheduling *scheduling.Scheduling, medication *medication.Medication, color string) *dto.Scheduling {
	schedulingResponse := dto.Scheduling{
		Id:             scheduling.ID,
		Name:           medication.Name,
		MedicationTime: scheduling.MedicationTime.In(time.FixedZone("UTC-3", -3*3600)),
		Dosage:         scheduling.Dosage,
		Quantity:       scheduling.Quantity,
		Status:         scheduling.Status,
		Image:          scheduling.Avatar,
		Color:          color,
		MedicationType: medication.Type.Name,
	}

	if scheduling.MedicationTakenAt != nil {
		takenTime := scheduling.MedicationTakenAt.In(time.FixedZone("UTC-3", -3*3600))
		schedulingResponse.MedicationTakenTime = &takenTime
	}

	return &schedulingResponse
}
