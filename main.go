package main

import (
	"fmt"

	"github.com/Alejo1595/paquetes-GO/greet"
	"github.com/Alejo1595/paquetes-GO/test"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println(greet.English())
	fmt.Println(test.Prueba())
	fmt.Println(quoteV3.HelloV3())
	fmt.Println(quoteV3.Concurrency())
}
