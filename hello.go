package main

import (
	"fmt"

	"github.com/joemooney/string"
	"github.com/joemooney/util"
)

func main() {
	fmt.Println("hello: "+string.Reverse("hello"))
	fmt.Printf("%v is prime? %v\n", 23, util.IsPrime(23))
}
