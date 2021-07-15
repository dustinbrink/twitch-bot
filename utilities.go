package main

import (
	"fmt"
	"time"
)

func log(s string) {
	fmt.Printf("[%s UTC] : %s \n", time.Now().UTC().Format("2006-01-02 15:04:05.000"), s)
}


