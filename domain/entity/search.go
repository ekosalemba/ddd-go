package entity

import (
	"strings"
	"time"
)

type SearchRequest struct {
	Origin       string   `json:"origin" xml:"origin"`
	Destination  string   `json:"destination" xml:"destination"`
	DepartDate   string   `json:"departDate" xml:"departDate"`
	ReturnDate   string   `json:"returnDate" xml:"returnDate"`
	ReturnStatus bool     `json:"returnStatus" xml:"returnStatus"`
	PaxCount     PaxCount `json:"paxCount" xml:"paxCount"`
}
type PaxCount struct {
	Adt int `json:"adt" xml:"adt"`
	Chd int `json:"chd" xml:"chd"`
	Inf int `json:"inf" xml:"inf"`
}

func (data SearchRequest) Validate() (bool, ErrorResponse) {
	valid := true
	var message string
	var messages []ErrorMessage

	if (len(strings.TrimSpace(data.Origin)) == 0) {
		valid = false
		messages = append(messages, ErrorMessage{"", "origin is required!"})
	}
	if (len(strings.TrimSpace(data.Destination)) == 0) {
		valid = false
		messages = append(messages, ErrorMessage{"", "destination is required!"})
	}
	if (len(strings.TrimSpace(data.DepartDate)) == 0) {
		valid = false
		messages = append(messages, ErrorMessage{"", "departDate is required!"})
	} else {
		_, err := time.Parse("2006-01-02", data.DepartDate)
		if err != nil {
			valid = false
			messages = append(messages, ErrorMessage{"", "departDate not valid!"})
		}
	}
	if (data.ReturnStatus) {
		if (len(strings.TrimSpace(data.ReturnDate)) == 0) {
			valid = false
			messages = append(messages, ErrorMessage{"", "returnDate is required!"})
		} else {
			_, err := time.Parse("2006-01-02", data.ReturnDate)
			if err != nil {
				valid = false
				messages = append(messages, ErrorMessage{"", "returnDate not valid!"})
			}
		}
	}
	if (data.PaxCount.Adt+data.PaxCount.Chd > 9) {
		valid = false
		messages = append(messages, ErrorMessage{"", "adt+chd <= 9"})
	}
	if (data.PaxCount.Adt == 0) {
		valid = false
		messages = append(messages, ErrorMessage{"", "adt is required!"})
	}
	if (data.PaxCount.Inf > data.PaxCount.Adt) {
		valid = false
		messages = append(messages, ErrorMessage{"", "inf <= adt"})
	}
	if len(messages) > 0 {
		valid = false
		message = "Some field are required"
	}
	response := ErrorResponse{BaseResponse{false, "", message}, messages}
	return valid, response
}

type SearchResponse struct {
	BaseResponse
	Data SearchResponseData `json:"data" xml:"data"`
}

type SearchResponseData struct {
	SearchKey string     `json:"searchKey" xml:"searchKey"`
	Schedules []Schedule `json:"schedules" xml:"schedules"`
}
type Schedule struct {
	JourneyType     string    `json:"journeyType" xml:"journeyType"`
	Origin          string    `json:"origin" xml:"origin"`
	OriginName      string    `json:"originName" xml:"originName"`
	Destination     string    `json:"destination" xml:"destination"`
	DestinationName string    `json:"destinationName" xml:"destinationName"`
	Journeys        []Journey `json:"journeys" xml:"journeys"`
}

type Journey struct {
	Origin                string                `json:"origin" xml:"origin"`
	OriginName            string                `json:"originName" xml:"originName"`
	Destination           string                `json:"destination" xml:"destination"`
	DestinationName       string                `json:"destinationName" xml:"destinationName"`
	Std                   string                `json:"std" xml:"std"`
	Sta                   string                `json:"sta" xml:"sta"`
	FlightDuration        int                   `json:"flightDuration" xml:"flightDuration"`
	Description           string                `json:"description" xml:"description"`
	Segments              []Segment             `json:"segments" xml:"segments"`
	AvailableCabinClasses []AvailableCabinClass `json:"availableCabinClasses" xml:"availableCabinClasses"`
}

type Segment struct {
	CarrierCode     string `json:"carrierCode" xml:"carrierCode"`
	FlightNo        string `json:"flightNo" xml:"flightNo"`
	Origin          string `json:"origin" xml:"origin"`
	OriginName      string `json:"originName" xml:"originName"`
	Destination     string `json:"destination" xml:"destination"`
	DestinationName string `json:"destinationName" xml:"destinationName"`
	Std             string `json:"std" xml:"std"`
	Sta             string `json:"sta" xml:"sta"`
	RouteStatus     string `json:"routeStatus" xml:"routeStatus"`
	Legs            []Leg  `json:"legs" xml:"legs"`
}

type Leg struct {
	Origin          string `json:"origin" xml:"origin"`
	OriginName      string `json:"originName" xml:"originName"`
	Destination     string `json:"destination" xml:"destination"`
	DestinationName string `json:"destinationName" xml:"destinationName"`
	Std             string `json:"std" xml:"std"`
	Sta             string `json:"sta" xml:"sta"`
}
type AvailableCabinClass struct {
	CabinClass          string              `json:"cabinClass" xml:"cabinClass"`
	AvailableSubClasses []AvailableSubClass `json:"availableSubClasses" xml:"availableSubClasses"`
}
type AvailableSubClass struct {
	SubClass     string        `json:"subClass" xml:"subClass"`
	SubClassKey  string        `json:"subClassKey" xml:"subClassKey"`
	SeatAvail    string        `json:"seatAvail" xml:"seatAvail"`
	Availability int           `json:"availability" xml:"availability"`
	StatusAvail  string        `json:"statusAvail" xml:"statusAvail"`
	ClassType    string        `json:"classType" xml:"classType"`
	Currency     string        `json:"currency" xml:"currency"`
	Price        float64       `json:"price" xml:"price"`
	PriceDetails []PriceDetail `json:"priceDetails" xml:"priceDetails"`
}
type PriceDetail struct {
	PaxType        string          `json:"paxType" xml:"paxType"`
	TotalFare      float64         `json:"totalFare" xml:"totalFare"`
	NtaFare        float64         `json:"ntaFare" xml:"ntaFare"`
	FareComponents []FareComponent `json:"fareComponents" xml:"fareComponents"`
}
type FareComponent struct {
	FareCode string  `json:"fareCode" xml:"fareCode"`
	FareDesc string  `json:"fareDesc" xml:"fareDesc"`
	Currency string  `json:"currency" xml:"currency"`
	Amount   float64 `json:"amount" xml:"amount"`
}
