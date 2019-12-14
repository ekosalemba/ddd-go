package infrastructure

import (
	"ddd-go/domain"
	"strconv"
	"strings"
	"time"
)

func MapSjSearchRequest(searchRequest domain.SearchRequest, path string) *strings.Reader {
	returnStatus := "NO"
	if (searchRequest.ReturnStatus) {
		returnStatus = "YES"
	}
	departDate, _ := time.Parse("2006-01-02", searchRequest.DepartDate)
	departDateStr := departDate.Format("02-Jan-2006")
	returnDate, _ := time.Parse("2006-01-02", searchRequest.ReturnDate)
	returnDateStr := returnDate.Format("02-Jan-2006")

	payload := strings.NewReader(
		"CityFrom=" + searchRequest.Origin +
			"&CityTo=" + searchRequest.Destination +
			"&DepartDate=" + departDateStr +
			"&ReturnStatus=" + returnStatus +
			"&ReturnDate=" + returnDateStr +
			"&PromoCode=" +
			"&Adult=" + strconv.Itoa(searchRequest.PaxCount.Adt) +
			"&Child=" + strconv.Itoa(searchRequest.PaxCount.Chd) +
			"&Infant==" + strconv.Itoa(searchRequest.PaxCount.Inf) +
			"&DEVICE_ID=" + SjConfigGetDeviceId() +
			"&SUBSCRIBE_ID=" + SjConfigGetSubscribeId() +
			"&USERNAME=" + SjConfigGetUserName() +
			"&PASSWORD=" + SjConfigGetPassword() +
			"&OS=" + SjConfigGetOs() +
			"&APPS_NAME=" + SjConfigGetAppsName() +
			"&APPS_VERSION=" + SjConfigGetAppsVersion() +
			"&API_URL=" + SjConfigGetAPibaseUrl() + path +
			"&USER_LOGIN=")
	return payload

}

func MapSjSearchResponse(responseSpl domain.SjSearchResponse) domain.SearchResponseData {
	searchResponseData := domain.SearchResponseData{}
	schedules := []domain.Schedule{}
	for key, result := range responseSpl.Data {
		schedule := domain.Schedule{}
		if (key == 0) {
			schedule.JourneyType = "Departure"
		} else {
			schedule.JourneyType = "Return"
		}
		for i, dataJourney := range result {
			if i == 0 {
				schedule.Origin = dataJourney.CityFrom
				schedule.OriginName = dataJourney.CityFromName
				schedule.Destination = dataJourney.CityTo
				schedule.DestinationName = dataJourney.CityToName
			}
			journey := domain.Journey{}
			journey.Origin = dataJourney.CityFrom
			journey.OriginName = dataJourney.CityFromName
			journey.Destination = dataJourney.CityTo
			journey.DestinationName = dataJourney.CityToName
			stdDate, _ := time.Parse("02-Jan-2006 15:04", dataJourney.Std_)
			journey.Std = stdDate.Format("2006-01-02 15:04")
			staDate, _ := time.Parse("02-Jan-2006 15:04", dataJourney.Sta_)
			journey.Sta = staDate.Format("2006-01-02 15:04")
			journey.FlightDuration, _ = strconv.Atoi(dataJourney.FlightTime)
			journey.Description = dataJourney.Keterangan

			for _, dataSegment := range dataJourney.Segments {
				segment := domain.Segment{}
				segment.CarrierCode = dataSegment.CarrierCode
				segment.FlightNo = dataSegment.NoFlight
				segment.Origin = dataSegment.DepartureStation
				segment.OriginName = dataSegment.DepartureStationName
				segment.Destination = dataSegment.ArrivalStation
				segment.DestinationName = dataSegment.ArrivalStationName
				stdDate, _ := time.Parse("02-Jan-2006 15:04", dataSegment.Std_)
				segment.Std = stdDate.Format("2006-01-02 15:04")
				staDate, _ := time.Parse("02-Jan-2006 15:04", dataSegment.Sta_)
				segment.Sta = staDate.Format("2006-01-02 15:04")
				segment.RouteStatus = dataSegment.RouteStatus

				for _, dataLeg := range dataSegment.Legs {
					leg := domain.Leg{}
					leg.Origin = dataLeg.DepartureStation
					leg.OriginName = dataLeg.DepartureStationName
					leg.Destination = dataLeg.ArrivalStation
					leg.DestinationName = dataLeg.ArrivalStationName

					stdDate, _ := time.Parse("02-Jan-2006 15:04", dataLeg.Std_)
					leg.Std = stdDate.Format("2006-01-02 15:04")
					staDate, _ := time.Parse("02-Jan-2006 15:04", dataLeg.Sta_)
					leg.Sta = staDate.Format("2006-01-02 15:04")
					segment.Legs = append(segment.Legs, leg)
				}
				journey.Segments = append(journey.Segments, segment)
			}

			availableCabinClass := domain.AvailableCabinClass{}
			availableCabinClass.CabinClass = "ECONOMY"
			for _, dataSubClass := range dataJourney.ClassesAvailableB2C.Economy {
				availableSubClass := domain.AvailableSubClass{}

				availableSubClass.SubClass = dataSubClass.Class
				availableSubClass.SubClassKey = dataSubClass.Key
				availableSubClass.SeatAvail = dataSubClass.SeatAvail
				availableSubClass.Availability, _ = strconv.Atoi(dataSubClass.Availability)
				availableSubClass.StatusAvail = dataSubClass.StatusAvail
				availableSubClass.ClassType = dataSubClass.ClassType
				availableSubClass.Currency = dataSubClass.Currency
				availableSubClass.Price, _ = strconv.ParseFloat(dataSubClass.Price, 64)
				for _, dataPriceDetail := range dataSubClass.PriceDetail {
					priceDetail := domain.PriceDetail{}
					priceDetail.PaxType = dataPriceDetail.PaxCategory
					priceDetail.TotalFare, _ = strconv.ParseFloat(dataPriceDetail.Total1, 64)
					priceDetail.NtaFare, _ = strconv.ParseFloat(dataPriceDetail.Nta1, 64)
					for _, dataFareComponent := range dataPriceDetail.FareComponent {
						fareComponent := domain.FareComponent{}
						fareComponent.FareCode = dataFareComponent.FareChargeTypeCode
						fareComponent.FareDesc = dataFareComponent.FareChargeTypeDesc
						fareComponent.Currency = dataFareComponent.CurrencyCode
						fareComponent.Amount, _ = strconv.ParseFloat(dataFareComponent.Amount, 64)
						priceDetail.FareComponents = append(priceDetail.FareComponents, fareComponent)
					}
					availableSubClass.PriceDetails = append(availableSubClass.PriceDetails, priceDetail)
				}
				availableCabinClass.AvailableSubClasses = append(availableCabinClass.AvailableSubClasses, availableSubClass)
			}
			journey.AvailableCabinClasses = append(journey.AvailableCabinClasses, availableCabinClass)

			schedule.Journeys = append(schedule.Journeys, journey)
		}
		schedules = append(schedules, schedule)
	}
	searchResponseData = domain.SearchResponseData{responseSpl.SearchKey, schedules}
	return searchResponseData
}

func MapSjBookRequest(bookRequest domain.BookRequest, path string) *strings.Reader {
	payload := strings.NewReader(
		"BookingCode=" + bookRequest.Journey.Origin +
			"&DEVICE_ID=" + SjConfigGetDeviceId() +
			"&SUBSCRIBE_ID=" + SjConfigGetSubscribeId() +
			"&USERNAME=" + SjConfigGetUserName() +
			"&PASSWORD=" + SjConfigGetPassword() +
			"&OS=" + SjConfigGetOs() +
			"&APPS_NAME=" + SjConfigGetAppsName() +
			"&APPS_VERSION=" + SjConfigGetAppsVersion() +
			"&API_URL=" + SjConfigGetAPibaseUrl() + path +
			"&USER_LOGIN=")
	return payload
}

func MapSjBookInfoRequest(bookInfoRequest domain.BookInfoRequest, path string) *strings.Reader {
	payload := strings.NewReader(
		"BookingCode=" + bookInfoRequest.BookingCode +
			"&DEVICE_ID=" + SjConfigGetDeviceId() +
			"&SUBSCRIBE_ID=" + SjConfigGetSubscribeId() +
			"&USERNAME=" + SjConfigGetUserName() +
			"&PASSWORD=" + SjConfigGetPassword() +
			"&OS=" + SjConfigGetOs() +
			"&APPS_NAME=" + SjConfigGetAppsName() +
			"&APPS_VERSION=" + SjConfigGetAppsVersion() +
			"&API_URL=" + SjConfigGetAPibaseUrl() + path +
			"&USER_LOGIN=")
	return payload
}

func MapSjBookInfoResponse(responseSpl domain.SjBookInfoResponse) domain.BookInfoData {
	bookInfoData := domain.BookInfoData{}
	bookInfoData.BookingCode = responseSpl.DATA.BookingCode
	bookInfoData.NumericBookingCode = responseSpl.DATA.NUMERICBOOKINGCODE
	bookInfoData.PromoCode = responseSpl.DATA.PromoCode
	bookInfoData.PaymentMethod = responseSpl.DATA.PAYMENTMETHOD
	bookInfoData.StatusShow = responseSpl.DATA.StatusShow
	bookInfoData.CheckInFlag = responseSpl.DATA.CHECKINFLAG
	bookInfoData.CheckInDate = responseSpl.DATA.CHECKINDATE
	bookInfoData.ItineraryDetails.ReservationDetails.BookingCode = responseSpl.DATA.YourItineraryDetails.ReservationDetails.BookingCode
	bookInfoData.ItineraryDetails.ReservationDetails.BookingDate = responseSpl.DATA.YourItineraryDetails.ReservationDetails.BookingDate
	bookInfoData.ItineraryDetails.ReservationDetails.CurrencyCode = responseSpl.DATA.YourItineraryDetails.ReservationDetails.CurrencyCode
	bookInfoData.ItineraryDetails.ReservationDetails.BalanceDue, _ = strconv.ParseFloat(responseSpl.DATA.YourItineraryDetails.ReservationDetails.BalanceDue, 64)
	bookInfoData.ItineraryDetails.ReservationDetails.BalanceDueRemarks = responseSpl.DATA.YourItineraryDetails.ReservationDetails.BalanceDueRemarks
	bookInfoData.ItineraryDetails.ReservationDetails.BookingStatus = responseSpl.DATA.YourItineraryDetails.ReservationDetails.Status
	bookInfoData.ItineraryDetails.ReservationDetails.TimeInfo = responseSpl.DATA.YourItineraryDetails.ReservationDetails.Time
	bookInfoData.ItineraryDetails.ReservationDetails.TimeInfoDescription = responseSpl.DATA.YourItineraryDetails.ReservationDetails.TimeDescription

	for _, paxDetail := range responseSpl.DATA.YourItineraryDetails.PassengerDetails {
		passengerDetail := domain.PassengerDetail{}
		passengerDetail.No, _ = strconv.Atoi(paxDetail.No)
		passengerDetail.Suffix = paxDetail.FirstName
		passengerDetail.FirstName = paxDetail.FirstName
		passengerDetail.LastName = paxDetail.LastName
		passengerDetail.SeatQty, _ = strconv.Atoi(paxDetail.SeatQty)
		passengerDetail.TicketNumber = paxDetail.TicketNumber
		passengerDetail.FrequentFlyerNumber = paxDetail.FrequentFlyerNumber
		passengerDetail.SpecialRequest = paxDetail.SpecialRequest
		passengerDetail.SeatNo = paxDetail.SeatNo
		bookInfoData.ItineraryDetails.PassengerDetails = append(bookInfoData.ItineraryDetails.PassengerDetails, passengerDetail)
	}

	bookInfoData.ItineraryDetails.PaymentDetails.CurrencyCode = responseSpl.DATA.YourItineraryDetails.PaymentDetails.CurrencyCode
	bookInfoData.ItineraryDetails.PaymentDetails.BasicFare, _ = strconv.ParseFloat(responseSpl.DATA.YourItineraryDetails.PaymentDetails.BasicFare, 64)
	bookInfoData.ItineraryDetails.PaymentDetails.OtherFare, _ = strconv.ParseFloat(responseSpl.DATA.YourItineraryDetails.PaymentDetails.Others, 64)
	bookInfoData.ItineraryDetails.PaymentDetails.STIFare, _ = strconv.ParseFloat(responseSpl.DATA.YourItineraryDetails.PaymentDetails.Sti, 64)
	bookInfoData.ItineraryDetails.PaymentDetails.TotalFare, _ = strconv.ParseFloat(responseSpl.DATA.YourItineraryDetails.PaymentDetails.Total, 64)
	bookInfoData.ItineraryDetails.PaymentDetails.DirectVoucher, _ = strconv.ParseFloat(responseSpl.DATA.YourItineraryDetails.PaymentDetails.DirectVoucher, 64)
	bookInfoData.ItineraryDetails.PaymentDetails.AddOnFare, _ = strconv.ParseFloat(responseSpl.DATA.YourItineraryDetails.PaymentDetails.AddOn, 64)
	bookInfoData.ItineraryDetails.PaymentDetails.NtaFare, _ = strconv.ParseFloat(responseSpl.DATA.YourItineraryDetails.PaymentDetails.Nta, 64)

	for i, splJourney := range responseSpl.DATA.YourItineraryDetails.ItineraryDetails.Journey {
		journey := domain.BookJourney{}
		if i == 0 {
			journey.JourneyType = "Departure"
		} else {
			journey.JourneyType = "Return"
		}
		for _, splSegment := range splJourney.Segment {
			segment := domain.BookSegment{}
			segment.FlownDate = splSegment.FlownDate
			segment.CarrierCode = splSegment.FlightNo[0:2]
			segment.FlightNo = splSegment.FlightNo
			segment.Origin = splSegment.CityFrom
			segment.OriginName = splSegment.CityFromName
			segment.Destination = splSegment.CityTo
			segment.DestinationName = splSegment.CityToName
			segment.StdLT = splSegment.StdLT
			segment.StaLT = splSegment.StaLT
			segment.ReservationStatus = splSegment.ReservationStatus
			segment.SubClass = splSegment.Class
			segment.CheckInStatus = splSegment.CheckInStatus

			journey.Segments = append(journey.Segments, segment)
		}
		bookInfoData.ItineraryDetails.Itineraries.Journeys = append(bookInfoData.ItineraryDetails.Itineraries.Journeys, journey)
	}

	for _, splContact := range responseSpl.DATA.YourItineraryDetails.ContactList {
		contact := domain.Contact{}
		contact.Type = splContact.Type
		contact.Value = splContact.Value
		contact.Description = splContact.Description
		bookInfoData.ItineraryDetails.ContactList = append(bookInfoData.ItineraryDetails.ContactList, contact)
	}

	bookInfoData.ItineraryDetails.AgentDetails.BookedBy = responseSpl.DATA.YourItineraryDetails.AgentDetails.BookedBy
	bookInfoData.ItineraryDetails.AgentDetails.IssuedBy = responseSpl.DATA.YourItineraryDetails.AgentDetails.IssuedBy

	for _, splRemark := range responseSpl.DATA.YourItineraryDetails.BookingRemarks {
		remark := domain.BookingRemark{}
		remark.CommentText = splRemark.CommentText
		remark.CreatedBy = splRemark.CreatedBy
		remark.CreatedDate = splRemark.CreatedDate
		remark.IPAddress = splRemark.IpAddress
		bookInfoData.ItineraryDetails.BookingRemarks = append(bookInfoData.ItineraryDetails.BookingRemarks, remark)
	}
	bookInfoData.ItineraryDetails.AdditionalInformation = responseSpl.DATA.YourItineraryDetails.AdditionalInformation
	return bookInfoData
}

func MapSjSetPaymentRequest(setPaymentRequest domain.SetPaymentRequest, path string) *strings.Reader {
	payload := strings.NewReader(
		"BOOKING_CODE=" + setPaymentRequest.BookingCode +
			"&NAMA_PEMEGANG_KARTU=" + setPaymentRequest.Payment.PaymentCardName +
			"&RECEIVED_PHONE=" + setPaymentRequest.Payment.PaymentContact.Phone +
			"&PAYMENT_METHOD=" + setPaymentRequest.Payment.PaymentMethod +
			"&EMAIL=" + setPaymentRequest.Payment.PaymentEmail +
			"&EMAIL_CP=" + setPaymentRequest.Payment.PaymentEmail +
			"&TENOR=" + setPaymentRequest.Payment.PaymentTenor +
			"&FFNO=" + setPaymentRequest.FFInfo.FFNumber +
			"&PASSWORD_FF=" + setPaymentRequest.FFInfo.FFPassword +
			"&POINT=" + setPaymentRequest.FFInfo.FFPoint +
			"&FF_TYPE=" + setPaymentRequest.FFInfo.FFType +

			"&DEVICE_ID=" + SjConfigGetDeviceId() +
			"&SUBSCRIBE_ID=" + SjConfigGetSubscribeId() +
			"&USERNAME=" + SjConfigGetUserName() +
			"&PASSWORD=" + SjConfigGetPassword() +
			"&OS=" + SjConfigGetOs() +
			"&APPS_NAME=" + SjConfigGetAppsName() +
			"&APPS_VERSION=" + SjConfigGetAppsVersion() +
			"&API_URL=" + SjConfigGetAPibaseUrl() + path +
			"&USER_LOGIN=")
	return payload
}

func MapSjSetPaymentResponse(responseSpl domain.SjSetPaymentResponse) domain.SetPaymentData {
	setPaymentData := domain.SetPaymentData{}
	setPaymentData.PaymentCode = responseSpl.Data[0].PaymentCode
	setPaymentData.PaymentMethod = responseSpl.Data[0].PaymentMethod
	setPaymentData.PaymentMethodDescription = responseSpl.Data[0].PaymentMethodDescription
	setPaymentData.PaymentAmount = responseSpl.Data[0].Amount
	setPaymentData.FFDetail = responseSpl.Data[0].FFDetail
	setPaymentData.SpecialRequest = responseSpl.Data[0].SpecialRequest
	setPaymentData.BNIWOW = responseSpl.Data[0].BNIWOW
	return setPaymentData
}
