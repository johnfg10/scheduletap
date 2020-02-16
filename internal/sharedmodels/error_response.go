package sharedmodels

// ErrorResponse is a struct representing a recoverable error occuring
type ErrorResponse struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
