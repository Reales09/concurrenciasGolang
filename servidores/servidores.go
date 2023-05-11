package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidores(servidor string, canal chan string) {
	_, error := http.Get(servidor)

	if error != nil {
		// fmt.Println(servidor, "no esta disponible")
		canal <- servidor + " no esta disponible"
	} else {
		// fmt.Println(servidor, "esta funcionando correctamente")
		canal <- servidor + " esta funcionando"
	}
}

func main() {
	inicio := time.Now()
	canal := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
		"https://gpsmobile.net/",
	}

	for _, servidor := range servidores {
		go revisarServidores(servidor, canal)
		// revisarServidores(servidor)
	}
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}
	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoPaso)
}
