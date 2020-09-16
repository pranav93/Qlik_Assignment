package main

import (
	"log"

	"github.com/pranav93/Qlik_Assignment/setup"
)

func init() {
	log.SetPrefix("[AppLog] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	setup.Server().Run()
}
