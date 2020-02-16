package github.io/johnfg10/scheduletap/internal/shift_models

type Shift struct {
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	Location         OfficeLocation `json:"location"`
	StartTime        time.Time      `json:"start_time"`
	Duration         time.Duration  `json:"duration"`
	DocstoreRevision interface{}    `json:"-"`
}