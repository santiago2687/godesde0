package main

import (
	"fmt"

	"github.com/santiago2687/godesde0/goroutines"
)

func main() {
	/*variables.MuestroEnteros()
	variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux.":
		fmt.Println("Esto es Linux")
	case "darwin.":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	numeroEntero, mensaje := ejercicios.DevolverIntAndString("50")
	fmt.Println(numeroEntero)
	fmt.Println(mensaje)

	teclado.IngresoNumeros()*/

	//fmt.Println(ejercicios.TablaMultiplicar())

	//files.GrabaTabla()

	//files.SumaTabla()

	//files.LeerArchivo()

	//funciones.Calculos()

	//funciones.LlamarClosure()

	//funciones.Exponencia(2)

	//arreglosslices.MuestroArreglos()

	//arreglosslices.MuestroSlice()

	//arreglosslices.Capacidad()

	//mapas.MostrarMapas()

	//users.AltaUsuario()

	/*Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)*/

	//deferpanic.VemosDefer()

	canal1 := make(chan bool)

	go goroutines.MiNombreLentoooo("Santiago Bermudez", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aqui")
}
