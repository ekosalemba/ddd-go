package infrastructure

import (
	"ddd-go/domain"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func SjSplSearch(searchRequest domain.SearchRequest) (domain.SearchResponse, error) {
	path := SjSearchPath()
	payload := MapSjSearchRequest(searchRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = SearchMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl domain.SjSearchResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if (errParse != nil) {
		fmt.Println("Parse failed with error : \n", errParse)
	}
	searchResponseData := domain.SearchResponseData{}
	var searchResponse domain.SearchResponse
	if err != nil {
		searchResponse = domain.SearchResponse{domain.BaseResponse{false, "", "Internal Server Error"}, domain.SearchResponseData{}}
	} else {
		searchResponseData = MapSjSearchResponse(responseSpl)
		searchResponse = domain.SearchResponse{domain.BaseResponse{true, "OK", ""}, searchResponseData}
	}
	return searchResponse, err
}

func SjSplSetPayment(setPaymentRequest domain.SetPaymentRequest) (domain.SetPaymentResponse, error) {
	path := SjBookInfoPath()
	payload := MapSjSetPaymentRequest(setPaymentRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = SetPaymentMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl domain.SjSetPaymentResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if (errParse != nil) {
		fmt.Println("Parse failed with error: ", errParse)
	}
	setPaymentData := domain.SetPaymentData{}
	var setPaymentResponse domain.SetPaymentResponse
	if err != nil {
		setPaymentResponse = domain.SetPaymentResponse{domain.BaseResponse{false, "", "Internal Server Error"}, domain.SetPaymentData{}}
	} else {
		setPaymentData = MapSjSetPaymentResponse(responseSpl)
		setPaymentResponse = domain.SetPaymentResponse{BaseResponse: domain.BaseResponse{true, "OK", ""}, Data: setPaymentData}
	}
	return setPaymentResponse, err
}

func SjSplBookInfo(bookInfoRequest domain.BookInfoRequest) (domain.BookInfoResponse, error) {
	path := SjBookInfoPath()
	payload := MapSjBookInfoRequest(bookInfoRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = BookInfoMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl domain.SjBookInfoResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if (errParse != nil) {
		fmt.Println("Parse failed with error: ", errParse)
	}
	bookInfoData := domain.BookInfoData{}
	var bookInfoResponse domain.BookInfoResponse
	if err != nil {
		bookInfoResponse = domain.BookInfoResponse{domain.BaseResponse{false, "", "Internal Server Error"}, domain.BookInfoData{}}
	} else {
		bookInfoData = MapSjBookInfoResponse(responseSpl)
		bookInfoResponse = domain.BookInfoResponse{BaseResponse: domain.BaseResponse{true, "OK", ""}, Data: bookInfoData}
	}
	return bookInfoResponse, err
	//LZYHFE
}

func SjSplBook(bookRequest domain.BookRequest) (domain.BookInfoResponse, error) {
	path := SjBookInfoPath()
	payload := MapSjBookRequest(bookRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = BookMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl domain.SjBookInfoResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if (errParse != nil) {
		fmt.Println("Parse failed with error: ", errParse)
	}
	bookInfoData := domain.BookInfoData{}
	var bookInfoResponse domain.BookInfoResponse
	if err != nil {
		bookInfoResponse = domain.BookInfoResponse{domain.BaseResponse{false, "", "Internal Server Error"}, domain.BookInfoData{}}
	} else {
		bookInfoData = MapSjBookInfoResponse(responseSpl)
		bookInfoResponse = domain.BookInfoResponse{BaseResponse: domain.BaseResponse{true, "OK", ""}, Data: bookInfoData}
	}
	return bookInfoResponse, err
	//LZYHFE
}

func SendRequest(path string, method string, body io.Reader) ([]byte, error) {
	fmt.Printf("Call %v %v, %v \n", method, path, body)
	request, err := http.NewRequest(method, SjConfigGetAPibaseUrl()+path, body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}
	client := new(http.Client)

	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		fmt.Println("The HTTP request failed with error %s\n", err)
		return nil, err
	}
	dataResponse, errParse := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("The HTTP request failed with error %s\n", err)
		return nil, err
	} else {
		if (errParse != nil) {
			fmt.Println("The HTTP request failed with error %s\n", errParse)
		}
		//fmt.Println(string(dataResponse))
		return dataResponse, nil
	}
}
