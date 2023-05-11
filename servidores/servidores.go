package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidores(servidor string) {
	_, error := http.Get(servidor)

	if error != nil {
		fmt.Println(servidor, "no esta disponible")
	} else {
		fmt.Println(servidor, "esta funcionando correctamente")
	}
}

func main() {
	inicio := time.Now()
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		go revisarServidores(servidor)
		// revisarServidores(servidor)
	}
	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoPaso)
}
