package main

import (
	"os"

	"github.com/surt-io/surt/cmd/surt-controller/app"
)

func main() {

	err := app.RunApp()

	if err != nil {
		os.Exit(1)
	}

}
