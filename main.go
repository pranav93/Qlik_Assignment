package main

import (
	"log"

	"github.com/pranav93/Qlik_Assignment/scripts"
	"github.com/pranav93/Qlik_Assignment/setup"
)

func init() {
	// Ideally there should be separate migration dockerfile
	log.SetPrefix("[AppLog] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	scripts.Migrate()
}

func main() {
	setup.Server().Run()
}
