package digitalocean

type Actions struct {
	Actions []Action     `json:"actions"`
	Meta    DropletsMeta `json:"meta"`
}

type Action struct {
	Id           int    `json:"id"`
	Status       string `json:"status"`
	Type         string `json:"type"`
	StartedAt    string `json:"started_at"`
	CompletedAt  string `json:"completed_at"`
	ResourceId   int    `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	Region       string `json:"region"`
}
