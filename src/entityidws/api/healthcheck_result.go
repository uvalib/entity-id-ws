package api

type HealthCheckResult struct {
	Healthy bool   `json:"healthy"`
	Message string `json:"message,omitempty"`
}
