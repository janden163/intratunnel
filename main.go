package main

import (
	"flag"
	"intratunnel/client"
	"intratunnel/server"
	"log"
)

func main() {
	mode := flag.String("mode", "", "start mode: client or server")
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *mode == "" {
		log.Fatalf("Mode not specified. Use -mode=client or -mode=server")
	}

	if *configPath == "" {
		log.Fatalf("Config file path not specified. Use -config=path/to/config.json")
	}

	switch *mode {
	case "client":
		clientConfig, err := client.LoadConfig(*configPath)
		if err != nil {
			log.Fatalf("Failed to load client config: %v", err)
		}
		client.Run(clientConfig)

	case "server":
		serverConfig, err := server.LoadConfig(*configPath)
		if err != nil {
			log.Fatalf("Failed to load server config: %v", err)
		}
		server.Run(serverConfig)

	default:
		log.Fatalf("Unknown mode: %s. Use -mode=client or -mode=server", *mode)
	}
}
