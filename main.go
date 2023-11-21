package main

import (
	"github.com/ghozilaaa/iso8583-simulator/config"
	"github.com/ghozilaaa/iso8583-simulator/internal"
)

func main() {
	cfg := config.ParseConfigFile("config.yaml")
	internal.NewApp(cfg).Start()
}
