package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numeroMultiplicador int
var err error
var texto string

func TablaMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese numero que represente la tabla de multiplicar que quiere formar : ")
		if scanner.Scan() {
			numeroMultiplicador, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("No ha ingreso ningun valor numerico, por favor ingreselo nuevamente : ")
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numeroMultiplicador, i, i*numeroMultiplicador)
	}

	return texto
}
