package main

import (
	"flag"
	"log"
	"os"

	"code.uber.internal/infra/mesos-go/cmd/example-scheduler/app"
)

func main() {
	cfg := app.NewConfig()
	fs := flag.NewFlagSet("scheduler", flag.ExitOnError)
	cfg.AddFlags(fs)
	fs.Parse(os.Args[1:])

	if err := app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
