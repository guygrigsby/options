package main

import (
	"fmt"

	"github.com/guygrigsby/options"
)

func main() {
	u := options.NewContext( // HLcter
		options.WithUsername("guy"), // HLcter
	) // HLcter

	fmt.Printf("%#v", u)

}
