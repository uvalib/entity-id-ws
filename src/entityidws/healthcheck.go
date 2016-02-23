package main

type HealthCheckResult struct {
	Healthy        bool   `json:"healthy"`
	Message        string `json:"message,omitempty"`
}

type HealthCheckResponse struct {
	CheckType      HealthCheckResult `json:"ezid"`
}

