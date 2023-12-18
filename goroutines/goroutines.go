package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentoooo(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Microsecond) // formula para demorar 1 segundo
		fmt.Println(letra)
	}
}
