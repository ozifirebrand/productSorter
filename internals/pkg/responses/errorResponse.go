package responses

type EResponse struct {
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"error"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ErrorResponse(message string, code int, err interface{}) EResponse {
	return EResponse{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  "error",
		},
		Error: err,
	}
}
