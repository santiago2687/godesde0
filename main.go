package main

import (
	"fmt"

	"github.com/santiago2687/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
