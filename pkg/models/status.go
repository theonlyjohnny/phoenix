package models

//A Status represents the state reported by the PhoenixAgent
type Status struct {
	CPUUsage float64 `json:"cpu_usage"`
	MemUsage float64 `json:"mem_usage"`

	Healthy bool `json:"healthy"`

	PhoenixID *string `header:"X-Phoenix-Id"` //for use on incoming POST /status requests
}
