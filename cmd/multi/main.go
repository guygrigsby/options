package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/guygrigsby/options"
)

func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	u := options.NewContext( // HLcter
		options.WithUsername("guy"),        // HLcter
		options.WithPassword("secret"),     // HLcter
		options.WithClient(client),         // HLcter
		options.WithTimeout(time.Second*2), // HLcter
	) // HLcter

	fmt.Printf("%v", u)
}
