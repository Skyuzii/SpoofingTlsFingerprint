package main

import (
	"Golang/Request"
	"Golang/Response"
	"compress/gzip"
	"encoding/json"
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/check-status", CheckStatus).Methods("GET")
	router.HandleFunc("/handle-get", HandleGet).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func CheckStatus(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	success := "true"
	json.NewEncoder(responseWriter).Encode(success)
}

func HandleGet(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	var handleGetRequest Request.HandleGetRequest
	_ = json.NewDecoder(request.Body).Decode(&handleGetRequest)
	client := cycletls.Init()

	resp, err := client.Do(handleGetRequest.Url, cycletls.Options{
		Proxy:     handleGetRequest.Proxy,
		Timeout:   handleGetRequest.Timeout,
		Headers:   handleGetRequest.Headers,
		Ja3:       handleGetRequest.Ja3,
		UserAgent: handleGetRequest.UserAgent,
	}, "GET")

	var handleGetResponse Response.HandleGetResponse

	if err != nil {
		handleGetResponse.Success = false
		handleGetResponse.Error = err.Error()

		json.NewEncoder(responseWriter).Encode(handleGetResponse)
	}

	handleGetResponse.Success = true
	handleGetResponse.Payload = &Response.HandleGetResponsePayload{
		Text:    DecodeResponse(&resp),
		Headers: resp.Response.Headers,
		Status:  resp.Response.Status,
	}

	json.NewEncoder(responseWriter).Encode(handleGetResponse)
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
