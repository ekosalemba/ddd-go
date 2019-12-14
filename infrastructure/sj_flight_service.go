package infrastructure

import (
	"ddd-go/domain"
)

func SjSearch(req domain.SearchRequest) domain.SearchResponse {
	searchResponse, _ := SjSplSearch(req)
	return searchResponse
}

func SjBook() {

}
func SjBookInfo(req domain.BookInfoRequest) domain.BookInfoResponse {
	response, _ := SjSplBookInfo(req)
	return response
}
func SjSetPayment(req domain.SetPaymentRequest) domain.SetPaymentResponse {
	response, _ := SjSplSetPayment(req)
	return response
}
