package domain

import (
	"strconv"
	"strings"
	"time"
)

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

type BookRequest struct {
	Journey             SearchRequest       `json:"journey" xml:"journey" form:"journey"`
	Contact             BookContact         `json:"contact" xml:"contact" form:"contact"`
	Pax                 Pax                 `json:"pax" xml:"pax" form:"pax"`
	SelectedItineraries SelectedItineraries `json:"selectedItineraries" xml:"selectedItineraries" form:"selectedItineraries"`
}

func (data BookRequest) Validate() (bool, ErrorResponse) {
	valid := true
	var message string
	var messages []ErrorMessage
	if len(strings.TrimSpace(data.Journey.Origin)) == 0 {
		messages = append(messages, ErrorMessage{"", "origin is required!"})
	}
	if len(strings.TrimSpace(data.Journey.Destination)) == 0 {
		messages = append(messages, ErrorMessage{"", "destination is required!"})
	}
	if len(strings.TrimSpace(data.Journey.DepartDate)) == 0 {
		messages = append(messages, ErrorMessage{"", "departDate is required!"})
	} else {
		_, err := time.Parse("2006-01-02", data.Journey.DepartDate)
		if err != nil {
			messages = append(messages, ErrorMessage{"", "departDate not valid!"})
		}
	}
	if data.Journey.PaxCount.Adt+data.Journey.PaxCount.Chd > 9 {
		messages = append(messages, ErrorMessage{"", "adt+chd <= 9"})
	}
	if data.Journey.PaxCount.Adt == 0 {
		messages = append(messages, ErrorMessage{"", "adt is required!"})
	}
	if data.Journey.PaxCount.Inf > data.Journey.PaxCount.Adt {
		messages = append(messages, ErrorMessage{"", "inf <= adt"})
	}

	// contact
	if len(strings.TrimSpace(data.Contact.Phone)) == 0 {
		messages = append(messages, ErrorMessage{"", "contact phone is required!"})
	}
	if len(strings.TrimSpace(data.Contact.Email)) == 0 {
		messages = append(messages, ErrorMessage{"", "contact email is required!"})
	}
	if len(strings.TrimSpace(data.Contact.Name)) == 0 {
		messages = append(messages, ErrorMessage{"", "contact name is required!"})
	}

	// pax
	if len(data.Pax.Adt) != data.Journey.PaxCount.Adt {
		messages = append(messages, ErrorMessage{"", "pax adt is not valid!"})
	} else {
		for i := 0; i < len(data.Pax.Adt); i++ {
			if len(data.Pax.Adt[i].Suffix) == 0 {
				messages = append(messages, ErrorMessage{"", "pax adt [" + strconv.Itoa(i) + "] suffix is not valid!"})
			}
			if len(data.Pax.Adt[i].FirstName) == 0 {
				messages = append(messages, ErrorMessage{"", "pax adt [" + strconv.Itoa(i) + "] first name is not valid!"})
			}
			if len(data.Pax.Adt[i].LastName) == 0 {
				messages = append(messages, ErrorMessage{"", "pax adt [" + strconv.Itoa(i) + "] last name is not valid!"})
			}
		}
	}
	if len(data.Pax.Chd) != data.Journey.PaxCount.Chd {
		messages = append(messages, ErrorMessage{"", "pax adt is not valid!"})
	} else {

		for i := 0; i < len(data.Pax.Chd); i++ {
			if len(data.Pax.Chd[i].Suffix) == 0 {
				messages = append(messages, ErrorMessage{"", "pax chd [" + strconv.Itoa(i) + "] suffix is not valid!"})
			}
			if len(data.Pax.Chd[i].FirstName) == 0 {
				messages = append(messages, ErrorMessage{"", "pax chd [" + strconv.Itoa(i) + "] first name is not valid!"})
			}
			if len(data.Pax.Chd[i].LastName) == 0 {
				messages = append(messages, ErrorMessage{"", "pax chd [" + strconv.Itoa(i) + "] last name is not valid!"})
			}
			_, err := time.Parse("2006-01-02", data.Pax.Chd[i].Dob)
			if err != nil {
				messages = append(messages, ErrorMessage{"", "pax chd [" + strconv.Itoa(i) + "] date of birth is not valid!"})
			}
		}
	}
	if len(data.Pax.Inf) != data.Journey.PaxCount.Inf {
		messages = append(messages, ErrorMessage{"", "pax inf is not valid!"})
	} else {
		var adtAssoc []int
		for i := 0; i < len(data.Pax.Inf); i++ {
			if len(data.Pax.Inf[i].Suffix) == 0 {
				messages = append(messages, ErrorMessage{"", "pax inf [" + strconv.Itoa(i) + "] suffix is not valid!"})
			}
			if len(data.Pax.Inf[i].FirstName) == 0 {
				messages = append(messages, ErrorMessage{"", "pax inf [" + strconv.Itoa(i) + "] first name is not valid!"})
			}
			if len(data.Pax.Inf[i].LastName) == 0 {
				messages = append(messages, ErrorMessage{"", "pax inf [" + strconv.Itoa(i) + "] last name is not valid!"})
			}
			_, err := time.Parse("2006-01-02", data.Pax.Inf[i].Dob)
			if err != nil {
				messages = append(messages, ErrorMessage{"", "pax inf [" + strconv.Itoa(i) + "] date of birth is not valid!"})
			}

			if data.Pax.Inf[i].AdtAssoc <= 0 || data.Pax.Inf[i].AdtAssoc > data.Journey.PaxCount.Adt {
				messages = append(messages, ErrorMessage{"", "pax inf [" + strconv.Itoa(i) + "] adt assoc is not valid!"})
			} else {
				isExistAdtAssoc := false
				for i := 0; i < len(adtAssoc); i++ {
					if adtAssoc[i] == data.Pax.Inf[i].AdtAssoc {
						isExistAdtAssoc = true
					}
				}
				if isExistAdtAssoc {
					adtAssoc = append(adtAssoc, data.Pax.Inf[i].AdtAssoc)
				} else {
					messages = append(messages, ErrorMessage{"", "pax inf [" + strconv.Itoa(i) + "] adt assoc duplicate!"})
				}
			}
		}
	}
	// selected itineraries
	if len(strings.TrimSpace(data.SelectedItineraries.SearchKey)) == 0 {
		messages = append(messages, ErrorMessage{"", "selected itineraries -> search key is required!"})
	}
	if len(data.SelectedItineraries.Journeys) == 0 {
		messages = append(messages, ErrorMessage{"", "selected itineraries -> journeys is required!"})
	} else {
		for i := 0; i < len(data.SelectedItineraries.Journeys); i++ {
			if len(data.SelectedItineraries.Journeys[i].JourneyType) == 0 {
				messages = append(messages, ErrorMessage{"", "selected itineraries -> journeys" + strconv.Itoa(i) + "-> journey type is required!"})
			} else {
				var validJT = strings.ToUpper(data.SelectedItineraries.Journeys[i].JourneyType) == "DEPARTURE" || strings.ToUpper(data.SelectedItineraries.Journeys[i].JourneyType) == "RETURN"
				if !validJT {
					messages = append(messages, ErrorMessage{"", "selected itineraries -> journeys[" + strconv.Itoa(i) + "]-> journey type is only Departure and Return!"})
				}
				for iCabin := 0; iCabin < len(data.SelectedItineraries.Journeys[i].SelectedCabin); iCabin++ {
					if len(data.SelectedItineraries.Journeys[i].SelectedCabin[iCabin].Key) == 0 {
						messages = append(messages, ErrorMessage{"", "selected itineraries -> journeys[" + strconv.Itoa(i) + "]-> selected cabin[" + strconv.Itoa(iCabin) + "] -> key is required!"})
					}
					if len(data.SelectedItineraries.Journeys[i].SelectedCabin[iCabin].SubClass) == 0 {
						messages = append(messages, ErrorMessage{"", "selected itineraries -> journeys[" + strconv.Itoa(i) + "]-> selected cabin[" + strconv.Itoa(iCabin) + "] -> subclass is required!"})
					}
				}
			}
		}
	}
	if data.Journey.ReturnStatus == true && len(data.SelectedItineraries.Journeys) < 2 {
		messages = append(messages, ErrorMessage{"", "selected itineraries -> journeys is not complete!"})
	}

	if len(messages) > 0 {
		valid = false
		message = "Some field are required"
	}
	response := ErrorResponse{BaseResponse{false, "", message}, messages}
	return valid, response
}

type BookContact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type Pax struct {
	Adt []struct {
		Suffix    string `json:"suffix"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		FFNumber  string `json:"FFNumber"`
	} `json:"adt"`
	Chd []struct {
		Suffix    string `json:"suffix"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Dob       string `json:"dob"`
	} `json:"chd"`
	Inf []struct {
		Suffix    string `json:"suffix"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Dob       string `json:"dob"`
		AdtAssoc  int    `json:"adtAssoc"`
	} `json:"inf"`
}
type SelectedItineraries struct {
	SearchKey string `json:"searchKey"`
	Journeys  []struct {
		JourneyType   string `json:"journeyType"`
		SelectedCabin []struct {
			SubClass string `json:"subClass"`
			Key      string `json:"key"`
		} `json:"selectedCabin"`
	} `json:"journeys"`
}
