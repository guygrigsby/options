package main

// START OMIT
type Struct struct {
	str string
	i   int
}

func create() {
	s := &Struct{"test", 1984} // HL
	println(s)
}

// END OMIT
