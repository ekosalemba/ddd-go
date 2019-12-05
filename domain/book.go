package domain

import "strings"

type BookInfoRequest struct {
	BookingCode string `json:"bookingCode" xml:"bookingCode" form:"bookingCode"`
}

type BookInfoResponse struct {
	BaseResponse
	Data BookInfoData `json:"data" xml:"data"`
}

func (data BookInfoRequest) Validate() (bool, ErrorResponse) {
	valid := true
	var message string
	var messages []ErrorMessage
	if (len(strings.TrimSpace(data.BookingCode)) == 0) {
		valid = false
		messages = append(messages, ErrorMessage{"", "bookingCode is required!"})
		message = "Some field are required"
	}
	response := ErrorResponse{BaseResponse{false, "", message}, messages}
	return valid, response
}

type BookInfoData struct {
	BookingCode        string `json:"bookingCode" xml:"bookingCode"`
	NumericBookingCode string `json:"numericBookingCode" xml:"numericBookingCode"`
	PromoCode          string `json:"promoCode" xml:"promoCode"`
	PaymentMethod      string `json:"paymentMethod" xml:"paymentMethod"`
	StatusShow         string `json:"statusShow" xml:"statusShow"`
	CheckInFlag        string `json:"checkInFlag" xml:"checkInFlag"`
	CheckInDate        string `json:"checkInDate" xml:"checkInDate"`
	ItineraryDetails   struct {
		ReservationDetails struct {
			BookingCode         string  `json:"bookingCode"`
			BookingDate         string  `json:"bookingDate"`
			CurrencyCode        string  `json:"currencyCode"`
			BalanceDue          float64 `json:"balanceDue"`
			BalanceDueRemarks   string  `json:"balanceDueRemarks"`
			BookingStatus       string  `json:"bookingStatus"`
			TimeInfo            string  `json:"timeInfo"`
			TimeInfoDescription string  `json:"timeInfoDescription"`
		} `json:"reservationDetails"`
		PassengerDetails []PassengerDetail `json:"passengerDetails"`
		Itineraries      struct {
			Journeys []BookJourney `json:"journeys"`
		} `json:"itineraries"`
		PaymentDetails struct {
			CurrencyCode  string  `json:"currencyCode"`
			BasicFare     float64 `json:"basicFare"`
			OtherFare     float64 `json:"otherFare"`
			STIFare       float64 `json:"STIFare"`
			TotalFare     float64 `json:"totalFare"`
			DirectVoucher float64 `json:"directVoucher"`
			AddOnFare     float64 `json:"addOnFare"`
			NtaFare       float64 `json:"ntaFare"`
		} `json:"paymentDetails"`
		ContactList  []Contact `json:"contactList"`
		AgentDetails struct {
			BookedBy string `json:"bookedBy"`
			IssuedBy string `json:"issuedBy"`
		} `json:"agentDetails"`
		BookingRemarks        []BookingRemark `json:"bookingRemarks"`
		AdditionalInformation struct {
		} `json:"additionalInformation"`
	} `json:"itineraryDetails"`
}
type PassengerDetail struct {
	No                  int    `json:"no"`
	Suffix              string `json:"suffix"`
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	SeatQty             int    `json:"seatQty"`
	TicketNumber        string `json:"ticketNumber"`
	FrequentFlyerNumber string `json:"frequentFlyerNumber"`
	SpecialRequest      string `json:"specialRequest"`
	SeatNo              string `json:"seatNo"`
}
type BookingRemark struct {
	CommentText string `json:"commentText"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	IPAddress   string `json:"ipAddress"`
}
type Contact struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

type BookSegment struct {
	FlownDate         string `json:"flownDate"`
	CarrierCode       string `json:"carrierCode"`
	FlightNo          string `json:"flightNo"`
	Origin            string `json:"origin"`
	OriginName        string `json:"originName"`
	Destination       string `json:"destination"`
	DestinationName   string `json:"destinationName"`
	StdLT             string `json:"stdLT"`
	StaLT             string `json:"staLT"`
	ReservationStatus string `json:"reservationStatus"`
	SubClass          string `json:"subClass"`
	CheckInStatus     string `json:"checkInStatus"`
}

type BookJourney struct {
	JourneyType string        `json:"journeyType"`
	Segments    []BookSegment `json:"segments"`
}
