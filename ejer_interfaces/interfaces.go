package ejerinterfaces

import (
	"fmt"

	"github.com/santiago2687/godesde0/interfaces"
)

func HumanosRespirando(humano interfaces.Humano) {
	humano.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", humano.Sexo())
}
