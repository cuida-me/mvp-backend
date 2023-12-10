package scheduling

import (
	"context"
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type schedulingService struct {
	repository scheduling.Repository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewSchedulingService(
	repository scheduling.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *schedulingService {
	return &schedulingService{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (s schedulingService) CreateAllSchedulingByMedication(ctx context.Context, medication medication.Medication) error {
	for _, schedule := range medication.Schedules {
		if schedule.Enabled {
			for _, time := range medication.Times {
				nextDate, isInNextWeek := s.getNextDayForTime(time.Time, schedule.DailyOfWeek)
				if isInNextWeek {
					continue
				}

				scheduling := scheduling.Scheduling{
					MedicationID:   medication.ID,
					Dosage:         medication.Dosage,
					Quantity:       medication.Quantity,
					MedicationTime: &nextDate,
					Avatar:         medication.Avatar,
					Status:         scheduling.TODO,
				}

				_, err := s.repository.CreateScheduling(ctx, &scheduling)
				if err != nil {
					s.log.Error(ctx, "error to create scheduling", log.Body{
						"error":         err.Error(),
						"medication_id": medication.ID,
						"time":          time.Time,
						"daily_of_week": schedule.DailyOfWeek,
					})

					return err
				}
			}
		}
	}

	return nil
}

func (s schedulingService) getNextDayForTime(timeToGet string, day int) (time.Time, bool) {
	now := time.Now().In(time.FixedZone("UTC-3", -3*3600))

	hour, err := time.Parse("15:04", timeToGet)
	if err != nil {
		panic(err)
	}

	dayInTheWeek := time.Weekday(day)

	nextDate := time.Date(now.Year(), now.Month(), now.Day(), hour.Hour(), hour.Minute(), hour.Second(), 0, time.FixedZone("UTC-3", -3*3600))

	for nextDate.Weekday() != dayInTheWeek {
		nextDate = nextDate.AddDate(0, 0, 1)
	}

	nextSaturday := now.AddDate(0, 0, int(time.Saturday-time.Now().Weekday()+7)%7).In(time.FixedZone("UTC-3", -3*3600))
	nextSaturday = time.Date(nextSaturday.Year(), nextSaturday.Month(), nextSaturday.Day(), 23, 59, 0, 0, nextSaturday.Location())

	return nextDate, nextDate.After(nextSaturday)
}
