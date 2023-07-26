package helper

type ErrorResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

func ResponseError(message interface{}) ErrorResponse {
	return ErrorResponse{Status: "failed", Message: message}
}

func ResponseSuccess(message interface{}) SuccessResponse {
	return SuccessResponse{Status: "success", Message: message}
}
