package main

import (
	"time"

	"github.com/johnfg10/scheduletap/internal/sharedmodels"
)

// Shift is a
type Shift struct {
	ID               string                      `json:"id"`
	Name             string                      `json:"name"`
	Location         sharedmodels.OfficeLocation `json:"location"`
	StartTime        time.Time                   `json:"start_time"`
	Duration         time.Duration               `json:"duration"`
	DocstoreRevision interface{}                 `json:"-"`
}
