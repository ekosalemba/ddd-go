package domain

type BaseResponse struct {
	Success      bool   `json:"success" xml:"success"`
	Message      string `json:"message" xml:"message"`
	ErrorMessage string `json:"errorMessage"  xml:"errorMessage"`
}

type ErrorResponse struct {
	BaseResponse
	Data []ErrorMessage `json:"data" xml:"data"`
}
type ErrorMessage struct {
	ErrorCode    string `json:"errorCode" xml:"errorCode"`
	ErrorMessage string `json:"errorMessage" xml:"errorMessage"`
}
