package main

import (
	"log"

	"github.com/patosinsuerte/ducky-clip/cmd"
	"github.com/patosinsuerte/ducky-clip/internal/database"
)

func main() {

	if err := database.InitDB(); err != nil {
		log.Fatalf("No se pudo iniciar la DB: %v", err)
	}

	cmd.RootCMD.Execute()
}


