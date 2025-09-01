package dto

type HealthCheckDTO struct {
	Status string `json:"status"`
}

type HealthCheckDTOEnvelope struct {
	Body HealthCheckDTO
}
