package technitium

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/chris-birch/docker-dns-sync/proto/technitium/v1/service"

	"io"
	"log"
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
			apiRequest()
		case 3: // DIE
			fmt.Printf("Action %s, ID: %s\n", rec.Action, rec.ContainerId)
		default:
			fmt.Printf("DISCARDING - Action %s, ID: %s\n", rec.Action, rec.ContainerId)
		}
	}
}

func httpClient() *http.Client {
	// Reuse the http.Client throughout your code base.
	client := &http.Client{Timeout: 3 * time.Second}
	return client
}

func apiRequest() {
	PORT := "5380"
	SERVER := "192.168.2.4"
	BASE := "http://" + SERVER + ":" + PORT + "/api/zones/records"

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", BASE+"/add", nil)
	if err != nil {
		log.Fatal(err)
	}

	// You can set custom headers here, if necessary
	req.Header.Set("User-Agent", "Go-HTTP-Client")
	ZONE := "dummy.net" //TODO get this as an envar
	q := req.URL.Query()
	q.Add("token", "0663aea32e1520aeefd175f8f9b9656394ac8012568259fd1dce0b0ebbe4bf18")
	q.Add("zone", ZONE)
	q.Add("domain", "CONTAINER_HOSTNAME"+"."+ZONE) // Container hostname
	q.Add("type", "CNAME")
	q.Add("cname", "yyyy.dummy.net")      // Docker host
	q.Add("comment", "This is a comment") //TODO Make this a JSON obj containing useful data
	req.URL.RawQuery = q.Encode()

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response status and body
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
