package main

import "fmt"

func main() {
	u := options.NewUserContext( // HLcter
		options.WithUsername("guy"), // HLcter
	) // HLcter

	fmt.Printf("%#v", u)

}
