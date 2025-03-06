package main

import (
	"github.com/chris-birch/docker-dns-sync/proto/technitium/v1/service"
	"github.com/chris-birch/docker-dns-sync/src/technitium"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// Configure time/timestamp format
	zerolog.TimeFieldFormat = time.RFC1123Z

	// Default level is info
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Add timestamp and pretty output
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log.Logger = zerolog.New(output).With().Timestamp().Logger()
}

func main() {
	// create a TCP listener on the specified port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal().Msgf("tcp connection failed: %v", err)
	}
	log.Printf("listening at %v", lis.Addr())

	// create a gRPC server instance
	s := grpc.NewServer()
	reflection.Register(s)

	// Initialise app config
	conf := new(technitium.Config)
	conf.Init()

	// Assign the GRPc services
	technitiumService := technitium.Service{Cfg: conf}

	// Register the services
	service.RegisterTechnitiumServiceServer(s, &technitiumService)

	// start listening to requests
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("grpc server failed: %v", err)
	}
}
