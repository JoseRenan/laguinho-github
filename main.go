package main

import (
	"log"
	"github.com/JoseRenan/laguinho-github/app"
)

func main() {
	log.Println("Rodando...")
	a := &app.App{}
	a.NewApp(":8080")
	a.Run()
}
