package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func log(s string) {
	fmt.Printf("[%s UTC] : %s \n", time.Now().UTC().Format("2006-01-02 15:04:05.000"), s)
}

func exitError(s string, err error) {
	log(s)
	log(err.Error())
	os.Exit(1)
}

func fetchApi(url string, c chan *[]byte) {
	resp, err := http.Get(url)

	if err != nil {
		log("Fetch API - error connecting to " + url)
		log(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log("Fetch API - error reading response")
		log(err.Error())
	}

	c <- &body
}