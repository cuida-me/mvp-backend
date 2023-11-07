package patient

type NewPatientSessionRequest struct {
	Ip       string `json:"ip"`
	DeviceID string `json:"device_id"`
}

type NewPatientSessionResponse struct {
	Token string `json:"token"`
}
