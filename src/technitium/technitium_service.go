package technitium

import (
	"fmt"

	"github.com/chris-birch/docker-dns-sync/proto/technitium/v1/service"

	"io"
)

type TechnitiumService struct {
	service.TechnitiumServiceServer
}

func (*TechnitiumService) ProcessRecord(stream service.TechnitiumService_ProcessRecordServer) error {
	for {
		record, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("end of stream")
		}
		if err != nil {
			return err
		}
		fmt.Println("recv record:", record)
	}
}
