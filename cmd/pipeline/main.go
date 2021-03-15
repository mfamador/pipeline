// main package for app application
package main

import (
	"github.com/Jeffail/benthos/v3/lib/service"
	_ "github.com/mfamador/pipeline/processor"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting Benthos Data Pipeline")
	service.Run()
}
