package main

import (
	"encoding/json"
)


func fetchSwansonQuote() (quote string) {
	log("Fetch Swanson Quote - fetching "+SWANSON_URI);

	body := make(chan *[]byte)

	go fetchApi(SWANSON_URI, body)

	var quotes []string
	b := *<-body
	log("Fetch Swanson Quote - response "+string(b));

	json.Unmarshal(b, &quotes)
	return quotes[0]
}

