package main

import (
	"log"
	"twitterproject/bd"
	"twitterproject/handlers"
)

func main(){
	if bd.CheckConnection() == 0 {
		log.Fatal("sin coneccion base de datos")
		return
	}
	handlers.Manejadores()
}