package api

type StatisticsResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Details Statistics `json:"statistics"`
}
