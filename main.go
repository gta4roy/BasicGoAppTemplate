package main

import (
	"flag"
	"gta4roy/app/log"
	"gta4roy/app/util"
)

var (
	listenAddr string
)

func init() {
	// Parse log level from command line
	logLevel := util.GetProperty(util.LogLevel)
	// Calling the SetLogLevel with the command-line argument
	log.SetLogLevel(logLevel, "logs.txt")
	log.Trace.Println("Loging initialised")
	flag.StringVar(&listenAddr, "listen-addr", ":"+util.GetProperty(util.Port), "server listen address")
	flag.Parse()

}

func main() {
	log.Trace.Println("Application started ")
}
