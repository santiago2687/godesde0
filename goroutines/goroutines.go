package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentoooo(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) // formula para demorar 1 segundo
		fmt.Println(letra)
	}
	canal1 <- true
}
