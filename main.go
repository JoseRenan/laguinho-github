package main

import "github.com/JoseRenan/laguinho-github/app"

func main() {
	a := &app.App{}
	a.NewApp(":8080")
	a.Run()
}
