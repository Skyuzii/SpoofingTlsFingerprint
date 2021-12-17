package main

import (
	"Golang/Request"
	"Golang/Response"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/Skyuzii/CycleTLS/cycletls"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/check-status", CheckStatus).Methods("GET")
	router.HandleFunc("/handle", Handle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("The proxy server is running")
}

func CheckStatus(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode("good")
}

func Handle(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	var handleRequest Request.HandleRequest
	json.NewDecoder(request.Body).Decode(&handleRequest)
	client := cycletls.Init()

	resp, err := client.Do(handleRequest.Url, cycletls.Options{
		Cookies:   handleRequest.Cookies,
		Body:      handleRequest.Body,
		Proxy:     handleRequest.Proxy,
		Timeout:   handleRequest.Timeout,
		Headers:   handleRequest.Headers,
		Ja3:       handleRequest.Ja3,
		UserAgent: handleRequest.UserAgent,
	}, handleRequest.Method)

	var handleResponse Response.HandleResponse

	if err != nil {
		fmt.Println(err)
		handleResponse.Success = false
		handleResponse.Error = err.Error()
		json.NewEncoder(responseWriter).Encode(handleResponse)
		return
	}

	handleResponse.Success = true
	handleResponse.Payload = &Response.HandleResponsePayload{
		Text:    DecodeResponse(&resp),
		Headers: resp.Response.Headers,
		Status:  resp.Response.Status,
		Cookies: resp.Response.Cookies,
		Url:     resp.Response.Url,
	}

	json.NewEncoder(responseWriter).Encode(handleResponse)
}

func DecodeResponse(response *cycletls.Response) string {
	switch response.Response.Headers["Content-Encoding"] {
	case "gzip":
		reader, _ := gzip.NewReader(strings.NewReader(response.Response.Body))
		defer reader.Close()
		readerResponse, _ := ioutil.ReadAll(reader)
		return string(readerResponse)
	default:
		return response.Response.Body
	}
}
