package service_registry

type Service struct {
	id          string     `json:"id"`
	name        string     `json:"name"`
	instances   []Instance `json:"instances"`
	healthCheck string     `json:"healthCheck"`
	ttl         float64    `json:"ttl"`
}
