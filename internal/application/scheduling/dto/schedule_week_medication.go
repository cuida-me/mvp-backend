package scheduling

type JobResponse struct {
	TotalToProcess       int
	ProcessedWithSuccess int
	ProcessedWithError   int
	Error                error
}
