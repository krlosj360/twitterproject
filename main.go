package main

import (
	"github.com/krlosj360/twitterproject/bd"
	"github.com/krlosj360/twitterproject/handlers"
	"log"

)

func main(){
	if bd.CheckConnection() == 0 {
		log.Fatal("sin coneccion base de datos")
		return
	}
	handlers.Manejadores()
}