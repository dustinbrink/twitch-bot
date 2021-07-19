package main

import (
	"encoding/json"
	"io"
	"net/http"
)


func fetchSwansonQuote() (quote string) {
	log("Fetch Swanson Quote - fetching "+SWANSON_URI);
	resp, err := http.Get(SWANSON_URI)

	if err != nil {
		log("Fetch Swanson Quote - error connecting to " + SWANSON_URI)
		log(err.Error())
		return ""
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log("Fetch Swanson Quote - error reading response ")
		log(err.Error())
		return ""
	}

	log("Fetch Swanson Quote - response "+string(body));

	var quotes []string
	json.Unmarshal([]byte(body), &quotes)
	return quotes[0]
}