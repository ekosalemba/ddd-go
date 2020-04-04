package service

import "ddd-go/domain/entity"

type FlightServiceInterface interface {
	Search(*entity.SearchRequest) *entity.SearchResponse
	Book(*entity.BookRequest) *entity.BookInfoResponse
	BookInfo(*entity.BookInfoRequest) *entity.BookInfoResponse
	SetPayment(*entity.SetPaymentRequest) *entity.SetPaymentResponse
}
