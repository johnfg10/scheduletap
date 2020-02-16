package sharedmodels

// Company is a struct representing a company
type Company struct {
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	Location         OfficeLocation `json:"location"`
	DocstoreRevision interface{}    `json:"-"`
}
