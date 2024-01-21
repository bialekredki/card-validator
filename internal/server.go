package internal

import (
	"bialekredki/card-validator/internal/algorithm"
	"bialekredki/card-validator/internal/common"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type includeQueryParameters struct {
	majorIndustry           bool
	majorIndustryIdentifier bool
}

func getIncludeQueryParameters(url url.Values) includeQueryParameters {
	var result includeQueryParameters
	if !url.Has("include") {
		return result
	}
	for _, value := range url["include"] {
		if value == "major-industry" {
			result.majorIndustry = true
		}
		if value == "major-industry-identifier" {
			result.majorIndustryIdentifier = true
		}
	}
	return result
}

func handler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, "Not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := request.URL.Query()
	if !query.Has("card-number") {
		http.Error(writer, "Missing required query parameters: 'card-number'", http.StatusUnprocessableEntity)
		return
	}
	cardNumbers := query["card-number"]
	if len(cardNumbers) > 10 {
		http.Error(writer, "Up to 10 card numbers to be verified are allowed.", http.StatusUnprocessableEntity)
		return
	}
	includeQueryParameters := getIncludeQueryParameters(query)
	var result = make(map[string]map[string]interface{})
	for _, cardNumber := range cardNumbers {
		isValid, _ := algorithm.IsValid(cardNumber)
		result[cardNumber] = make(map[string]interface{})
		result[cardNumber]["is-valid"] = isValid
		if includeQueryParameters.majorIndustry {
			result[cardNumber]["major-industry"] = GetMajorIndustry(cardNumber)
		}
		if firstDigit, err := common.FirstDigit(cardNumber); includeQueryParameters.majorIndustryIdentifier && err == nil {
			result[cardNumber]["major-industry-identifier"] = firstDigit
		}
	}
	response, err := json.Marshal(result)
	if err != nil {
		http.Error(writer, "Internal server errror", http.StatusInternalServerError)
		return
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(response)
}

func Serve(port uint16) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
