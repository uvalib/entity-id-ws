package api

//
// HealthCheckResult -- used for a health check result; we can have many of these
//
type HealthCheckResult struct {
	Healthy bool   `json:"healthy"`
	Message string `json:"message,omitempty"`
}

//
// end of file
//
