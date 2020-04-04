package infrastructure

import (
	"ddd-go/domain/entity"
	"ddd-go/domain/service"
)

// FlightServiceImpl Implements service.FlightServiceInterface
type SjFlightServiceImpl struct {
}

// NewNewsRepositoryWithRDB returns initialized NewsRepositoryImpl
func NewSjFlightServiceImpl() service.FlightServiceInterface {
	return &SjFlightServiceImpl{}
}

func (r *SjFlightServiceImpl) Search(req *entity.SearchRequest) *entity.SearchResponse {
	searchResponse, _ := SjSplSearchV2(req)
	return &searchResponse
}

func (r *SjFlightServiceImpl) Book(req *entity.BookRequest) *entity.BookInfoResponse {
	response, _ := SjSplBookV2(req)
	return &response
}

func (r *SjFlightServiceImpl) BookInfo(req *entity.BookInfoRequest) *entity.BookInfoResponse {
	response, _ := SjSplBookInfoV2(req)
	return &response
}

func (r *SjFlightServiceImpl) SetPayment(req *entity.SetPaymentRequest) *entity.SetPaymentResponse {
	response, _ := SjSplSetPaymentV2(req)
	return &response
}

/*
func SjBook(req entity.BookRequest) entity.BookInfoResponse {
	response, _ := SjSplBook(req)
	return response
}

func SjBookInfo(req entity.BookInfoRequest) entity.BookInfoResponse {
	response, _ := SjSplBookInfo(req)
	return response
}
func SjSetPayment(req entity.SetPaymentRequest) entity.SetPaymentResponse {
	response, _ := SjSplSetPayment(req)
	return response
}
*/
