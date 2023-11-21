package medication

type CreateScheduleRequest struct {
	DailyOfWeek *int   `json:"daily_of_week"`
	Time        string `json:"time"`
}

type CreateScheduleResponse struct {
	ID          uint64 `json:"id"`
	DailyOfWeek *int   `json:"daily_of_week"`
	LiteralDay  string `json:"literal_day"`
	Time        string `json:"time"`
}
