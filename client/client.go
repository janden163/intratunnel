package client

import (
	"intratunnel/pkg/logger"
	"intratunnel/pkg/proxy"
	"net"
	"time"
)

func Run(config *ClientConfig) {
	logger.Init("client.log")

	for {
		conn, err := net.Dial("tcp", config.ServerAddr)
		if err != nil {
			logger.Error("Failed to connect to server: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		logger.Info("Connected to server")

		go handleConnection(conn, config.LocalAddr)
	}
}

func handleConnection(serverConn net.Conn, localAddr string) {
	defer serverConn.Close()

	localConn, err := net.Dial("tcp", localAddr)
	if err != nil {
		logger.Error("Failed to connect to local service: %v", err)
		return
	}
	defer localConn.Close()

	errCh := make(chan error, 2)
	go proxy.ProxyData(localConn, serverConn, errCh)
	go proxy.ProxyData(serverConn, localConn, errCh)

	err = <-errCh
	if err != nil {
		logger.Error("Connection closed with error: %v", err)
	}
}
