package main

import (
	"fmt"
	"net"
)

const DEFAULT_BATCH_SIZE = 100

func main() {
	// --verbose: show prove results
	// --host: host to prove
	// --batch_size: number of requests to send in one batch
	var success int
	var fail int
	for i := 0; i < DEFAULT_BATCH_SIZE; i++ {
		_, err := net.LookupHost("golang.org")
		if err != nil {
			fmt.Println(err)
			fail++
			continue
		}
		success++
		fmt.Println("ok")	
	}

	fmt.Printf("success: %d, fail: %d\n", success, fail)
}