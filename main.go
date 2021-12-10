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

	options := cycletls.Options{
		Timeout:   handleGetRequest.Timeout,
		Body:      "",
		Headers:   handleGetRequest.Headers,
		Ja3:       handleGetRequest.Ja3,
		UserAgent: handleGetRequest.UserAgent,
	}

	if handleGetRequest.Proxy != "" {
		options.Proxy = handleGetRequest.Proxy
	}

	resp, err := client.Do(handleGetRequest.Url, options, "GET")

	var responseText string
	var handleGetResponse Response.HandleGetResponse

	if err != nil {
		handleGetResponse.Success = false
		handleGetResponse.Error = err.Error()

		json.NewEncoder(responseWriter).Encode(handleGetResponse)
	}

	switch resp.Response.Headers["Content-Encoding"] {
	case "gzip":
		reader, _ := gzip.NewReader(strings.NewReader(resp.Response.Body))
		readerResponse, _ := ioutil.ReadAll(reader)
		responseText = string(readerResponse)
		defer reader.Close()
	default:
		responseText = resp.Response.Body
	}

	handleGetResponse.Success = true
	handleGetResponse.Payload = &Response.HandleGetResponsePayload{
		Text:    responseText,
		Headers: resp.Response.Headers,
		Status:  resp.Response.Status,
	}

	json.NewEncoder(responseWriter).Encode(handleGetResponse)
}
