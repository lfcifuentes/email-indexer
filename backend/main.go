package main

import (
	"log"
	"os"

	"github.com/lfcifuentes/email-indexer/backend/src/core/helpers"
	"github.com/lfcifuentes/email-indexer/backend/src/http/server"
)

// @title			Email Indexer API documentation
// @version		1.0
// @description	This is a email indexer API documentation.
// @termsOfService	http://swagger.io/terms/
// @contact.name	Vatia
// @contact.email	lfcifuentes28@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// @Accept			json
// @Produce		json
//
// @Security		Bearer
// @BasePath		/
func main() {
	// Verificar si el archivo .env existe
	if _, err := os.Stat(".env"); err == nil {
		if err := helpers.LoadEnvFile(".env"); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	// Start the api server
	server.StartServer()
}
