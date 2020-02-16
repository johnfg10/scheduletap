
type Position struct {
	ID               string        `json:"id"`
	Company          Company       `json:"company"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	StartTime        time.Time     `json:"start_time"`
	Duration         time.Duration `json:"duration"`
	DocstoreRevision interface{}   `json:"-"`
}