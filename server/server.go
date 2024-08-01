package server

import (
	"crypto/tls"
	"intratunnel/pkg/logger"
	"intratunnel/pkg/proxy"
	"log"
	"net"
)

func Run(config *ServerConfig) {
	cert, err := tls.LoadX509KeyPair(config.TLSCert, config.TLSKey)
	if err != nil {
		log.Fatalf("Failed to load certificates: %v", err)
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", config.ListenAddr, tlsConfig)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", config.ListenAddr, err)
	}
	defer listener.Close()
	logger.Init("server.log")
	log.Printf("Server is listening on port %s", config.ListenAddr)

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			logger.Error("Failed to accept client connection: %v", err)
			continue
		}
		go handleClient(clientConn)
	}
}

func handleClient(clientConn net.Conn) {
	defer clientConn.Close()

	userConn, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		logger.Error("Failed to connect to user: %v", err)
		return
	}
	defer userConn.Close()

	errCh := make(chan error, 2)
	go proxy.ProxyData(clientConn, userConn, errCh)
	go proxy.ProxyData(userConn, clientConn, errCh)

	err = <-errCh
	if err != nil {
		logger.Error("Connection closed with error: %v", err)
	}
}
