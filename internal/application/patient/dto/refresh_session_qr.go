package patient

type RefreshSessionQRRequest struct {
	OldQR string `json:"old_qr"`
}

type RefreshSessionQRResponse struct {
	Token string `json:"token"`
}
