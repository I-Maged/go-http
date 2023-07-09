package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal("Error: ", err)
		// os.Exit(1)
	}

	fmt.Println(resp)
}
