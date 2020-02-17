package irisextensions

// APIResponse a
type APIResponse struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

func NewApiResponse(Type string, description string) APIResponse {
	response := APIResponse{}
	response.Type = Type
	response.Description = description
	return response
}

// NewErrorResponse creates a new APIResponse with the error type
func NewErrorResponse(description string) APIResponse {
	return NewApiResponse("error", description)
}

func NewSucessResponse(description string) APIResponse {
	return NewApiResponse("success", description)
}

func NewNotFoundResponse(description string) APIResponse {
	return NewApiResponse("notfound", description)
}

func NewInternalErrorResponse(err error) APIResponse {
	return NewApiResponse("interalerror", err.Error())
}
