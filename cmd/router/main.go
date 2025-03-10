package main

import (
	// internal "ras/gen/protobuf"
	"context"
	"fmt"
	"net"
	"ras/config"
	"ras/management"

	"github.com/charmbracelet/log"
	"github.com/mackerelio/go-osstat/uptime"
	"github.com/spf13/viper"
	api "gitlab.com/ras995910/router-api-protos/gen"
	"google.golang.org/grpc"
)

var port int

func init() {
	config.LoadConfigFile()
	port = viper.GetInt("grpc.port")
}

type ApiService struct {
	api.UnimplementedPingServiceServer
}

func (s *ApiService) Ping(ctx context.Context, r *api.ServerStatusRequest) (*api.ServerStatusResponse, error) {
	log.Info("Ping request handling", "request", r)
	uptime, err := uptime.Get()
	if err != nil {
		err = fmt.Errorf("failed get uptime of router: %s", err)
		log.Error(err)
		return nil, err
	}
	resp := api.ServerStatusResponse{
		Uptime: uptime.String(),
	}
	log.Infof("Uptime - %s", uptime)
	return &resp, nil
}

func main() {
	server := grpc.NewServer(
	// grpc.ConnectionTimeout(time.Duration(viper.GetInt64("grpc.timeout"))),
	)
	service := ApiService{}

	api.RegisterPingServiceServer(server, &service)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Failed create listener on port %d: %s", port, err)
	}

	log.Info("Server running", "port", port)
	log.Info(management.GetOSInfo())

	if err = server.Serve(listener); err != nil {
		log.Errorf("Failed serve grpc services: %s", err)
	}
	log.Info("Server running")
}
