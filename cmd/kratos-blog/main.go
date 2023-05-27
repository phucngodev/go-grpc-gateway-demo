package main

import (
	"flag"
	"os"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	configFile string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&configFile, "f", "../../configs", "config path, eg: -f config.yaml")
}

func main() {
	flag.Parse()
	app, cleanup, err := createApp(configFile)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
