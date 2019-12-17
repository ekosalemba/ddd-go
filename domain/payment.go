package domain

import "strings"

type SetPaymentRequest struct {
	BookingCode string  `json:"bookingCode"`
	Payment     Payment `json:"payment"`
	FFInfo      FFInfo  `json:"FFInfo"`
}
type SetPaymentResponse struct {
	BaseResponse
	Data SetPaymentData `json:"data" xml:"data"`
}
type Payment struct {
	PaymentMethod   string         `json:"paymentMethod"`
	PaymentCardName string         `json:"paymentCardName"`
	PaymentTenor    string         `json:"paymentTenor"`
	PaymentEmail    string         `json:"paymentEmail"`
	PaymentContact  PaymentContact `json:"paymentContact"`
}
type PaymentContact struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type FFInfo struct {
	FFType     string `json:"FFType"`
	FFNumber   string `json:"FFNumber"`
	FFPassword string `json:"FFPassword"`
	FFPoint    int    `json:"FFPoint"`
}

func (data SetPaymentRequest) Validate() (bool, ErrorResponse) {
	valid := true
	var message string
	var messages []ErrorMessage
	if (len(strings.TrimSpace(data.BookingCode)) == 0) {
		messages = append(messages, ErrorMessage{"", "bookingCode is required!"})
	}
	if (len(strings.TrimSpace(data.Payment.PaymentMethod)) == 0) {
		messages = append(messages, ErrorMessage{"", "paymentMethod is required!"})
	}
	if len(messages) > 0 {
		valid = false
		message = "Some field are required"
	}
	response := ErrorResponse{BaseResponse{false, "", message}, messages}
	return valid, response
}

type SetPaymentData struct {
	PaymentCode              string  `json:"paymentCode"`
	PaymentMethod            string  `json:"paymentMethod"`
	PaymentMethodDescription string  `json:"paymentMethodDescription"`
	PaymentAmount            float64 `json:"paymentAmount"`
	FFDetail                 string  `json:"FFDetail"`
	SpecialRequest           string  `json:"specialRequest"`
	BNIWOW                   string  `json:"BNIWOW"`
}
