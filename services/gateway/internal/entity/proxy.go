package entity

type Proxy struct {
	Status    string `json:"status,omitempty"`
	Address   string `json:"address,omitempty"`
	StartedAt string `json:"started_at,omitempty"`
}
