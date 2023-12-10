package scheduling

import "time"

type DailyScheduling struct {
	Day        int          `json:"day"`
	DayName    string       `json:"day_name"`
	MonthName  string       `json:"month_name"`
	Date       time.Time    `json:"date"`
	DayWeek    int          `json:"day_week"`
	DayColors  []string     `json:"day_colors"`
	Scheduling []Scheduling `json:"schedulings"`
}

type Scheduling struct {
	Id                  uint64    `json:"id"`
	Name                string    `json:"name"`
	MedicationTime      time.Time `json:"medication_time"`
	MedicationTakenTime time.Time `json:"medication_taken_time"`
	Dosage              string    `json:"dosage"`
	Quantity            int       `json:"quantity"`
	Status              string    `json:"status"`
	Image               string    `json:"image"`
	Color               string    `json:"color"`
}

func (slice DailyScheduling) Len() int {
	return len(slice.Scheduling)
}

func (slice DailyScheduling) Less(i, j int) bool {
	return slice.Scheduling[i].MedicationTime.After(slice.Scheduling[j].MedicationTime)
}

func (slice DailyScheduling) Swap(i, j int) {
	slice.Scheduling[i], slice.Scheduling[j] = slice.Scheduling[j], slice.Scheduling[i]
}
