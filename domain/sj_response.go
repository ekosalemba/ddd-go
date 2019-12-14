package domain

type SjSearchResponse struct {
	Data [][]struct {
		CityFrom            string `json:"CityFrom"`
		CityFromName        string `json:"CityFromName"`
		CityTo              string `json:"CityTo"`
		CityToName          string `json:"CityToName"`
		ClassesAvailableB2C struct {
			Economy []struct {
				Availability string `json:"Availability"`
				Class        string `json:"Class"`
				ClassType    string `json:"ClassType"`
				Currency     string `json:"Currency"`
				Key          string `json:"Key"`
				Price        string `json:"Price"`
				PriceDetail  []struct {
					FareComponent []struct {
						Amount             string `json:"Amount"`
						CurrencyCode       string `json:"CurrencyCode"`
						FareChargeTypeCode string `json:"FareChargeTypeCode"`
						FareChargeTypeDesc string `json:"FareChargeTypeDesc"`
					} `json:"FareComponent"`
					Nta1        string `json:"Nta_1"`
					PaxCategory string `json:"PaxCategory"`
					Total1      string `json:"Total_1"`
				} `json:"PriceDetail"`
				PriceDetailY struct {
					Availability string `json:"Availability"`
					Class        string `json:"Class"`
					Currency     string `json:"Currency"`
					Key          string `json:"Key"`
					Price        string `json:"Price"`
					PriceDetail  []struct {
						FareComponent []struct {
							Amount             string `json:"Amount"`
							CurrencyCode       string `json:"CurrencyCode"`
							FareChargeTypeCode string `json:"FareChargeTypeCode"`
							FareChargeTypeDesc string `json:"FareChargeTypeDesc"`
						} `json:"FareComponent"`
						Nta1        string `json:"Nta_1"`
						PaxCategory string `json:"PaxCategory"`
						Total1      string `json:"Total_1"`
					} `json:"PriceDetail"`
					SeatAvail   string `json:"SeatAvail"`
					StatusAvail string `json:"StatusAvail"`
				} `json:"PriceDetail_Y"`
				SeatAvail   string `json:"SeatAvail"`
				StatusAvail string `json:"StatusAvail"`
			} `json:"ECONOMY"`
		} `json:"ClassesAvailable_B2C"`
		FlightTime string `json:"FlightTime"`
		Keterangan string `json:"Keterangan"`
		Segments   []struct {
			ArrivalStation       string `json:"ArrivalStation"`
			ArrivalStationName   string `json:"ArrivalStationName"`
			CarrierCode          string `json:"CarrierCode"`
			DepartureStation     string `json:"DepartureStation"`
			DepartureStationName string `json:"DepartureStationName"`
			Legs                 []struct {
				ArrivalStation       string `json:"ArrivalStation"`
				ArrivalStationName   string `json:"ArrivalStationName"`
				DepartureStation     string `json:"DepartureStation"`
				DepartureStationName string `json:"DepartureStationName"`
				Sta                  string `json:"Sta"`
				Sta_                 string `json:"Sta_"`
				Std                  string `json:"Std"`
				Std_                 string `json:"Std_"`
			} `json:"Legs"`
			NoFlight    string `json:"NoFlight"`
			RouteStatus string `json:"RouteStatus"`
			Sta         string `json:"Sta"`
			Sta_        string `json:"Sta_"`
			Std         string `json:"Std"`
			Std_        string `json:"Std_"`
		} `json:"Segments"`
		Sta  string `json:"Sta"`
		Sta_ string `json:"Sta_"`
		Std  string `json:"Std"`
		Std_ string `json:"Std_"`
	} `json:"DATA"`
	ErrorCode    string `json:"ERROR_CODE" xml""ERROR_CODE"`
	ErrorMessage string `json:"ERROR_MESSAGE" xml"ERROR_MESSAGE"`
	StiStatus    string `json:"STI_STATUS" xml""STI_STATUS"`
	SearchKey    string `json:"SearchKey" xml"SearchKey"`
	TimeInt      string `json:"TIME_INT" xml""TIME_INT"`
}

type SjBookInfoResponse struct {
	DATA struct {
		BookingCode          string `json:"BookingCode"`
		CHECKINDATE          string `json:"CHECKIN_DATE"`
		CHECKINFLAG          string `json:"CHECKIN_FLAG"`
		ErrorCode            string `json:"ErrorCode"`
		ErrorMessage         string `json:"ErrorMessage"`
		NUMERICBOOKINGCODE   string `json:"NUMERIC_BOOKING_CODE"`
		PAYMENTMETHOD        string `json:"PAYMENT_METHOD"`
		PromoCode            string `json:"PromoCode"`
		Username             string `json:"Username"`
		YourItineraryDetails struct {
			AdditionalInformation struct{} `json:"AdditionalInformation"`
			AgentDetails          struct {
				BookedBy string `json:"BookedBy"`
				IssuedBy string `json:"IssuedBy"`
			} `json:"AgentDetails"`
			BookingRemarks []struct {
				CommentText string `json:CommentText`
				CreatedBy   string `json:CreatedBy`
				CreatedDate string `json:CreatedDate`
				IpAddress   string `json:IpAddress`
			} `json:"BookingRemarks"`
			ContactList []struct {
				Description string `json:"Description"`
				Type        string `json:"Type"`
				Value       string `json:"Value"`
			} `json:"ContactList"`
			ItineraryDetails struct {
				Journey []struct {
					Segment []struct {
						CheckInStatus     string `json:"CheckInStatus"`
						CityFrom          string `json:"CityFrom"`
						CityFromName      string `json:"CityFromName"`
						CityTo            string `json:"CityTo"`
						CityToName        string `json:"CityToName"`
						Class             string `json:"Class"`
						FlightNo          string `json:"FlightNo"`
						FlownDate         string `json:"FlownDate"`
						ReservationStatus string `json:"ReservationStatus"`
						StaLT             string `json:"StaLT"`
						StdLT             string `json:"StdLT"`
					} `json:"Segment"`
				} `json:"Journey"`
			} `json:"ItineraryDetails"`
			PassengerDetails []struct {
				FirstName           string `json:"FirstName"`
				FrequentFlyerNumber string `json:"FrequentFlyerNumber"`
				LastName            string `json:"LastName"`
				No                  string `json:"No"`
				SeatNo              string `json:"SeatNo"`
				SeatQty             string `json:"SeatQty"`
				SpecialRequest      string `json:"SpecialRequest"`
				Suffix              string `json:"Suffix"`
				TicketNumber        string `json:"TicketNumber"`
			} `json:"PassengerDetails"`
			PaymentDetails struct {
				AddOn         string `json:"AddOn"`
				BasicFare     string `json:"BasicFare"`
				CurrencyCode  string `json:"CurrencyCode"`
				DirectVoucher string `json:"DirectVoucher"`
				Nta           string `json:"Nta"`
				Others        string `json:"Others"`
				Sti           string `json:"Sti"`
				Total         string `json:"Total"`
			} `json:"PaymentDetails"`
			ReservationDetails struct {
				BalanceDue        string `json:"BalanceDue"`
				BalanceDueRemarks string `json:"BalanceDueRemarks"`
				BookingCode       string `json:"BookingCode"`
				BookingDate       string `json:"BookingDate"`
				CurrencyCode      string `json:"CurrencyCode"`
				Status            string `json:"Status"`
				Time              string `json:"Time"`
				TimeDescription   string `json:"TimeDescription"`
			} `json:"ReservationDetails"`
		} `json:"YourItineraryDetails"`
		StatusShow string `json:"status_show"`
	} `json:"DATA"`
	ERRORCODE    string `json:"ERROR_CODE"`
	ERRORMESSAGE string `json:"ERROR_MESSAGE"`
}

type SjSetPaymentResponse struct {
	Data []struct {
		PaymentCode              string  `json:"PAYMENT_CODE"`
		PaymentMethod            string  `json:"PAYMENT_METHOD"`
		PaymentMethodDescription string  `json:"PAYMENT_METHOD_DESCRIPTION"`
		Amount                   float64 `json:"AMOUNT"`
		FFDetail                 string  `json:"FF_DETAIL"`
		SpecialRequest           string  `json:"SPECIAL_REQUEST"`
		BNIWOW                   string  `json:"BNIWOW"`
	} `json:"DATA"`
	ErrorCode    string `json:"ERROR_CODE" xml""ERROR_CODE"`
	ErrorMessage string `json:"ERROR_MESSAGE" xml"ERROR_MESSAGE"`
}
