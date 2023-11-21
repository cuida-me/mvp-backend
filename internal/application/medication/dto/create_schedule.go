package medication

type CreateScheduleRequest struct {
	DailyOfWeek *int     `json:"daily_of_week"`
	Times       []string `json:"times"`
}

type CreateScheduleResponse struct {
	ID          uint64   `json:"id"`
	DailyOfWeek *int     `json:"daily_of_week"`
	LiteralDay  string   `json:"literal_day"`
	Enabled     bool     `json:"enabled"`
	Times       []string `json:"times"`
}
