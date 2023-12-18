package users

import (
	"fmt"
	"time"

	"github.com/santiago2687/godesde0/modelos"
)

func AltaUsuario() {
	usuario := new(modelos.User)
	usuario.AddUser(10, "Santiago", time.Now(), true)
	fmt.Println(usuario)
}
