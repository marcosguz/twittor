package main

import (
	"log"

	"github.com/marcosguz/twittor/bd"
	"github.com/marcosguz/twittor/handlers"
)

func main(){
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}