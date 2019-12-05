package infrastructure

import (
	"encoding/json"
	"ddd-go/domain"
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

func SjSplBook() {

}
func SjSplSetPayment() {

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
		//searchResponse = domain.SearchResponse{domain.BaseResponse{true, "OK", ""}, searchResponseData}
		bookInfoResponse = domain.BookInfoResponse{BaseResponse: domain.BaseResponse{true, "OK", ""}, Data: bookInfoData}
	}
	//return responseSpl, err
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
