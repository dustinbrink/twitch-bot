package main

import (
	"fmt"
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

