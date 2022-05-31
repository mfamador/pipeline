// main package for app application
package main

import (
	"context"
	_ "github.com/Jeffail/benthos/v3/public/components/all"
	"github.com/Jeffail/benthos/v3/public/service"
	_ "github.com/mfamador/pipeline/processor"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting Benthos v3 Data Pipeline")
	service.RunCLI(context.Background())
}
