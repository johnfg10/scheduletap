package sharedmodels

// OfficeLocation is a struct repusenting the location of the office/site
type OfficeLocation struct {
	ID               string      `json:"id"`
	Address          string      `json:"address"`
	DocstoreRevision interface{} `json:"-"`
}
