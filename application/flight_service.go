package application

import (
	"ddd-go/domain/entity"
	"ddd-go/infrastructure"
	"fmt"
	"net/http"
	"strings"
)

func Search(vendorCode string, request entity.SearchRequest) (int, *entity.SearchResponse, entity.ErrorResponse) {
	vendorCode = strings.ToUpper(vendorCode)
	switch vendorCode {
	case "SJ":
		valid, response := request.Validate()
		if !valid {
			return http.StatusBadRequest, &entity.SearchResponse{}, response
		} else {
			SjService := infrastructure.NewSjFlightServiceImpl()
			return http.StatusOK, SjService.Search(&request), entity.ErrorResponse{}
		}
	default:
		message := fmt.Sprintf("Vendor Code '%s' is not available", vendorCode)
		sr := entity.SearchResponse{entity.BaseResponse{false, message, ""}, entity.SearchResponseData{}}
		return http.StatusOK, &sr, entity.ErrorResponse{}
	}
}

func Book(request entity.BookRequest) (int, *entity.BookInfoResponse, entity.ErrorResponse) {
	valid, response := request.Validate()
	if !valid {
		return http.StatusBadRequest, &entity.BookInfoResponse{}, response
	} else {
		SjService := infrastructure.NewSjFlightServiceImpl()
		return http.StatusOK, SjService.Book(&request), entity.ErrorResponse{}
	}
}

func BookInfo(request entity.BookInfoRequest) (int, *entity.BookInfoResponse, entity.ErrorResponse) {
	valid, response := request.Validate()
	if !valid {
		return http.StatusBadRequest, &entity.BookInfoResponse{}, response
	} else {
		SjService := infrastructure.NewSjFlightServiceImpl()
		return http.StatusOK, SjService.BookInfo(&request), entity.ErrorResponse{}
	}
}

func SetPayment(request entity.SetPaymentRequest) (int, *entity.SetPaymentResponse, entity.ErrorResponse) {
	valid, response := request.Validate()
	if !valid {
		return http.StatusBadRequest, &entity.SetPaymentResponse{}, response
	} else {
		SjService := infrastructure.NewSjFlightServiceImpl()
		return http.StatusOK, SjService.SetPayment(&request), entity.ErrorResponse{}
	}
}
