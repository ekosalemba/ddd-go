package infrastructure

import (
	"ddd-go/domain"
)

func SjSearch(searchRequest domain.SearchRequest) domain.SearchResponse {
	searchResponse, _ := SjSplSearch(searchRequest)
	return searchResponse
}

func SjBook() {

}
func SjBookInfo(bookInfoRequest domain.BookInfoRequest) domain.BookInfoResponse {
	response, _ := SjSplBookInfo(bookInfoRequest)
	return response
}
func SjSetPayment() {

}
