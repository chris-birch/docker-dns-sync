package technitium

import (
	"fmt"
	"github.com/chris-birch/docker-dns-sync/proto/technitium/v1/service"

	"io"
)

type Service struct {
	service.UnimplementedTechnitiumServiceServer
}

func (*Service) ProcessRecord(stream service.TechnitiumService_ProcessRecordServer) error {
	for {
		rec, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("end of stream")
		}
		if err != nil {
			return err
		}
		switch rec.Action {
		case 0: // CREATE
			fmt.Printf("Action %s, ID: %s\n", rec.Action, rec.ContainerId)
		case 3: // DIE
			fmt.Printf("Action %s, ID: %s\n", rec.Action, rec.ContainerId)
		default:
			fmt.Printf("DISGARDING - Action %s, ID: %s\n", rec.Action, rec.ContainerId)

		}
	}
}
