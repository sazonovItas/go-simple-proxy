package entity

type Proxy struct {
	ID        string `json:"id,omitempty"`
	Status    string `json:"status,omitempty"`
	Address   string `json:"address,omitempty"`
	StartedAt string `json:"started_at,omitempty"`
}
