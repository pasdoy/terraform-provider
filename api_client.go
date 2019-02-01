package main

import (
	"github.com/dghubble/sling"
	"net/http"
	"strconv"
)

func getAPIClient(apiKey string) *sling.Sling {
	return sling.New().
		Client(http.DefaultClient).
		Base("https://api.cloudkarafka.com").
		SetBasicAuth("", apiKey)
}

func getCustomerClient(apiKey string) *sling.Sling {
	return sling.New().
		Client(http.DefaultClient).
		Base("https://customer.cloudkarafka.com").
		SetBasicAuth("", apiKey)
}

func idToString(id interface{}) string {
	return strconv.Itoa(int(id.(float64)))
}
