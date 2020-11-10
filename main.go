package main

import (
	"net"

	proxyproto "github.com/pires/go-proxyproto"
	"go.uber.org/zap"
)

const (
	listenAddress = "0.0.0.0:8080"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		logger.Error("failed to create listener", zap.String("address", listenAddress), zap.Error(err))
		panic(err)
	}

	proxyListener := &proxyproto.Listener{Listener: listener}
	defer proxyListener.Close()

	for {
		connection, err := proxyListener.Accept()
		if err != nil {
			logger.Error("failed to establish new connection", zap.Error(err))
			continue
		} else {
			logger.Debug("successfully established new connection")
		}

		remoteAddress := connection.RemoteAddr()
		if remoteAddress == nil {
			logger.Error("failed to retrieve remote address")
		} else {
			logger.Info("successfully retrieved remote address", zap.String("remote", remoteAddress.String()))
			connection.Write([]byte(remoteAddress.String()))
		}

		if err := connection.Close(); err != nil {
			logger.Error("failed to close connection", zap.Error(err))
			continue
		}
		logger.Debug("successfully closed connection")
	}
}
