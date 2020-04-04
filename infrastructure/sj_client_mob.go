package infrastructure

import (
	"ddd-go/domain/entity"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func SjSplSearchV2(searchRequest *entity.SearchRequest) (entity.SearchResponse, error) {
	path := SjSearchPath()
	payload := MapSjSearchRequest(*searchRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = SearchMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl entity.SjSearchResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if (errParse != nil) {
		fmt.Println("Parse failed with error : \n", errParse)
	}
	searchResponseData := entity.SearchResponseData{}
	var searchResponse entity.SearchResponse
	if err != nil {
		searchResponse = entity.SearchResponse{entity.BaseResponse{false, "", "Internal Server Error"}, entity.SearchResponseData{}}
	} else {
		if (responseSpl.ErrorCode == "EC:0000") {
			searchResponseData = MapSjSearchResponse(responseSpl)
			searchResponse = entity.SearchResponse{entity.BaseResponse{true, "OK", ""}, searchResponseData}
		} else {
			searchResponse = entity.SearchResponse{entity.BaseResponse{false, responseSpl.ErrorMessage, ""}, entity.SearchResponseData{}}
		}
	}
	return searchResponse, err
}

func SjSplSearch(searchRequest entity.SearchRequest) (entity.SearchResponse, error) {
	path := SjSearchPath()
	payload := MapSjSearchRequest(searchRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = SearchMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl entity.SjSearchResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if (errParse != nil) {
		fmt.Println("Parse failed with error : \n", errParse)
	}
	searchResponseData := entity.SearchResponseData{}
	var searchResponse entity.SearchResponse
	if err != nil {
		searchResponse = entity.SearchResponse{entity.BaseResponse{false, "", "Internal Server Error"}, entity.SearchResponseData{}}
	} else {
		if (responseSpl.ErrorCode == "EC:0000") {
			searchResponseData = MapSjSearchResponse(responseSpl)
			searchResponse = entity.SearchResponse{entity.BaseResponse{true, "OK", ""}, searchResponseData}
		} else {
			searchResponse = entity.SearchResponse{entity.BaseResponse{false, responseSpl.ErrorMessage, ""}, entity.SearchResponseData{}}
		}
	}
	return searchResponse, err
}

func SjSplSetPayment(setPaymentRequest entity.SetPaymentRequest) (entity.SetPaymentResponse, error) {
	path := SjSetPaymentPath()
	payload := MapSjSetPaymentRequest(setPaymentRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = SetPaymentMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl entity.SjSetPaymentResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if errParse != nil {
		fmt.Println("Parse failed with error: ", errParse)
	}
	setPaymentData := entity.SetPaymentData{}
	var setPaymentResponse entity.SetPaymentResponse
	if err != nil {
		setPaymentResponse = entity.SetPaymentResponse{entity.BaseResponse{false, "", "Internal Server Error"}, entity.SetPaymentData{}}
	} else {
		if responseSpl.ErrorCode == "EC:0000" {
			setPaymentData = MapSjSetPaymentResponse(responseSpl)
			setPaymentResponse = entity.SetPaymentResponse{BaseResponse: entity.BaseResponse{true, "OK", ""}, Data: setPaymentData}
		} else {
			setPaymentResponse = entity.SetPaymentResponse{BaseResponse: entity.BaseResponse{Success: false, Message: responseSpl.ErrorMessage, ErrorMessage: ""}, Data: entity.SetPaymentData{}}
		}
	}
	return setPaymentResponse, err
}

func SjSplSetPaymentV2(setPaymentRequest *entity.SetPaymentRequest) (entity.SetPaymentResponse, error) {
	path := SjSetPaymentPath()
	payload := MapSjSetPaymentRequest(*setPaymentRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = SetPaymentMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl entity.SjSetPaymentResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if errParse != nil {
		fmt.Println("Parse failed with error: ", errParse)
	}
	setPaymentData := entity.SetPaymentData{}
	var setPaymentResponse entity.SetPaymentResponse
	if err != nil {
		setPaymentResponse = entity.SetPaymentResponse{entity.BaseResponse{false, "", "Internal Server Error"}, entity.SetPaymentData{}}
	} else {
		if responseSpl.ErrorCode == "EC:0000" {
			setPaymentData = MapSjSetPaymentResponse(responseSpl)
			setPaymentResponse = entity.SetPaymentResponse{BaseResponse: entity.BaseResponse{true, "OK", ""}, Data: setPaymentData}
		} else {
			setPaymentResponse = entity.SetPaymentResponse{BaseResponse: entity.BaseResponse{Success: false, Message: responseSpl.ErrorMessage, ErrorMessage: ""}, Data: entity.SetPaymentData{}}
		}
	}
	return setPaymentResponse, err
}

func SjSplBookInfo(bookInfoRequest entity.BookInfoRequest) (entity.BookInfoResponse, error) {
	path := SjBookInfoPath()
	payload := MapSjBookInfoRequest(bookInfoRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = BookInfoMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl entity.SjBookInfoResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if (errParse != nil) {
		fmt.Println("Parse failed with error: ", errParse)
	}
	bookInfoData := entity.BookInfoData{}
	var bookInfoResponse entity.BookInfoResponse
	if err != nil {
		bookInfoResponse = entity.BookInfoResponse{entity.BaseResponse{false, "", "Internal Server Error"}, entity.BookInfoData{}}
	} else {
		if responseSpl.ERRORCODE == "EC:0000" {
			bookInfoData = MapSjBookInfoResponse(responseSpl)
			bookInfoResponse = entity.BookInfoResponse{BaseResponse: entity.BaseResponse{true, "OK", ""}, Data: bookInfoData}
		} else {
			bookInfoResponse = entity.BookInfoResponse{BaseResponse: entity.BaseResponse{false, responseSpl.ERRORMESSAGE, ""}, Data: entity.BookInfoData{}}
		}
	}
	return bookInfoResponse, err
	//LZYHFE
}

func SjSplBookInfoV2(bookInfoRequest *entity.BookInfoRequest) (entity.BookInfoResponse, error) {
	path := SjBookInfoPath()
	payload := MapSjBookInfoRequest(*bookInfoRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = BookInfoMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl entity.SjBookInfoResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if (errParse != nil) {
		fmt.Println("Parse failed with error: ", errParse)
	}
	bookInfoData := entity.BookInfoData{}
	var bookInfoResponse entity.BookInfoResponse
	if err != nil {
		bookInfoResponse = entity.BookInfoResponse{entity.BaseResponse{false, "", "Internal Server Error"}, entity.BookInfoData{}}
	} else {
		if responseSpl.ERRORCODE == "EC:0000" {
			bookInfoData = MapSjBookInfoResponse(responseSpl)
			bookInfoResponse = entity.BookInfoResponse{BaseResponse: entity.BaseResponse{true, "OK", ""}, Data: bookInfoData}
		} else {
			bookInfoResponse = entity.BookInfoResponse{BaseResponse: entity.BaseResponse{false, responseSpl.ERRORMESSAGE, ""}, Data: entity.BookInfoData{}}
		}
	}
	return bookInfoResponse, err
	//LZYHFE
}

func SjSplBook(bookRequest entity.BookRequest) (entity.BookInfoResponse, error) {
	path := SjBookPath()
	payload := MapSjBookRequest(bookRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = BookMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl entity.SjBookInfoResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if errParse != nil {
		fmt.Println("Parse failed with error: ", errParse)
	}
	bookInfoData := entity.BookInfoData{}
	var bookInfoResponse entity.BookInfoResponse
	if err != nil {
		bookInfoResponse = entity.BookInfoResponse{entity.BaseResponse{false, "", "Internal Server Error"}, entity.BookInfoData{}}
	} else {
		if responseSpl.ERRORCODE == "EC:0000" {
			bookInfoData = MapSjBookInfoResponse(responseSpl)
			bookInfoResponse = entity.BookInfoResponse{BaseResponse: entity.BaseResponse{true, "OK", ""}, Data: bookInfoData}
		} else {
			bookInfoResponse = entity.BookInfoResponse{BaseResponse: entity.BaseResponse{false, responseSpl.ERRORMESSAGE, ""}, Data: entity.BookInfoData{}}
		}
	}
	return bookInfoResponse, err
	//LZYHFE
}

func SjSplBookV2(bookRequest *entity.BookRequest) (entity.BookInfoResponse, error) {
	path := SjBookPath()
	payload := MapSjBookRequest(*bookRequest, path)

	var response []byte = nil
	var err error = nil
	if SjUseMock() {
		response, err = BookMock(), nil
	} else {
		response, err = SendRequest(path, "POST", payload)
	}

	var responseSpl entity.SjBookInfoResponse
	errParse := json.Unmarshal(response, &responseSpl)
	if errParse != nil {
		fmt.Println("Parse failed with error: ", errParse)
	}
	bookInfoData := entity.BookInfoData{}
	var bookInfoResponse entity.BookInfoResponse
	if err != nil {
		bookInfoResponse = entity.BookInfoResponse{entity.BaseResponse{false, "", "Internal Server Error"}, entity.BookInfoData{}}
	} else {
		if responseSpl.ERRORCODE == "EC:0000" {
			bookInfoData = MapSjBookInfoResponse(responseSpl)
			bookInfoResponse = entity.BookInfoResponse{BaseResponse: entity.BaseResponse{true, "OK", ""}, Data: bookInfoData}
		} else {
			bookInfoResponse = entity.BookInfoResponse{BaseResponse: entity.BaseResponse{false, responseSpl.ERRORMESSAGE, ""}, Data: entity.BookInfoData{}}
		}
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
