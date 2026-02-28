
package main

import (
	"log"

	"github.com/akhilbabu26/multi-brand_backend_2/internal/app"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/routes"
)

func main() {

	// initialize application
	app.Init()

	// setup router
	r := routes.Setup()

	// run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("server failed:", err)
	}
}