package utils

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
)

type Response struct {
	Data    interface{}     `json:"data"`
	Message string          `json:"message"`
	Status  int             `json:"status"`
	Errors  []ErrorResponse `json:"errors"`
}

func BuildResponse(data *Response) map[string]interface{} {
	return map[string]interface{}{
		"data":    data.Data,
		"status":  data.Status,
		"message": data.Message,
	}
}
