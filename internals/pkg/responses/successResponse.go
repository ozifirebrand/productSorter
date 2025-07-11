package responses

type SResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func SuccessResponse(message string, code int, data interface{}) SResponse {
	return SResponse{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  "success",
		},
		Data: data,
	}
}
