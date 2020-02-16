package shiftmodels

import (
	"time"

	"github.com/johnfg10/scheduletap/internal/sharedmodels"
)

type Position struct {
	ID               string               `json:"id"`
	Company          sharedmodels.Company `json:"company"`
	Name             string               `json:"name"`
	Description      string               `json:"description"`
	StartTime        time.Time            `json:"start_time"`
	Duration         time.Duration        `json:"duration"`
	DocstoreRevision interface{}          `json:"-"`
}
