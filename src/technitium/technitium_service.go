package technitium

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/chris-birch/docker-dns-sync/proto/technitium/v1/message"
	"net/http"
	"time"

	"github.com/chris-birch/docker-dns-sync/proto/technitium/v1/service"

	"io"

	"github.com/rs/zerolog/log"
)

const (
	ADD = "add"
	DEL = "delete"
)

type Service struct {
	service.UnimplementedTechnitiumServiceServer
	Cfg *Config
}

type Config struct {
	Token  string `env:"TOKEN,required"`
	Port   string `env:"PORT,required"`
	Server string `env:"SERVER,required"`
	Zone   string `env:"ZONE,required"`
	Client *http.Client
}

type RespStatus struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func (c *Config) Init() {
	if err := env.Parse(c); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}
	c.Client = &http.Client{Timeout: 3 * time.Second}
}

func (s *Service) ProcessRecord(stream service.TechnitiumService_ProcessRecordServer) error {
	for {
		rec, err := stream.Recv()
		if err == io.EOF {
			log.Debug().Msg("Connection closed")
		}
		if err != nil {
			return err
		}
		switch rec.Action {
		case 2: // START
			err = apiRequest(rec, s.Cfg, ADD)
			if err != nil {
				log.Error().Err(err).Msg("Failed to process record")
			} else {
				log.Info().
					Str("Zone", s.Cfg.Zone).
					Str("DomainName", rec.Name+"."+s.Cfg.Zone).
					Str("Data", rec.Data).
					Msg("Added record")
			}
		case 3: // DIE
			err = apiRequest(rec, s.Cfg, DEL)
			if err != nil {
				log.Error().Err(err).Msg("Failed to process record")
			} else {
				log.Info().
					Str("Zone", s.Cfg.Zone).
					Str("DomainName", rec.Name+"."+s.Cfg.Zone).
					Str("Data", rec.Data).
					Msg("Deleted record")
			}
		default:
			log.Info().
				Str("Action", rec.Action.String()).
				Str("Hostname", rec.Name).
				Msg("Discarding record")
		}
	}
}

func apiRequest(rec *message.DnsRecord, cfg *Config, method string) error {
	BASE := "http://" + cfg.Server + ":" + cfg.Port + "/api/zones/records"

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", BASE+"/"+method, nil)
	if err != nil {
		return err
	}

	// Set custom headers
	req.Header.Set("User-Agent", "Go-HTTP-Client")

	// Set query params
	q := req.URL.Query()
	q.Add("token", cfg.Token)
	q.Add("zone", cfg.Zone)
	q.Add("domain", rec.Name+"."+cfg.Zone) // Container hostname
	q.Add("type", "CNAME")
	q.Add("cname", rec.Data) // Docker host
	q.Add("comments", rec.ContainerId)
	req.URL.RawQuery = q.Encode()

	// Log the request
	log.Debug().Str("Query", req.URL.String()).Msg("Technitium API")

	// Perform the request
	client := cfg.Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check request status
	if resp.StatusCode != 200 {
		err := fmt.Errorf("technitium api response status code %d", resp.StatusCode)
		return err
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("failed to read response body")
	}

	// Check JSON response
	var s RespStatus
	err = json.Unmarshal(body, &s)
	if err != nil {
		return errors.New("failed to unmarshal response body")
	}

	// Log the API response
	log.Debug().RawJSON("response", body).Msg("Technitium API")

	// Check for any errors
	if s.ErrorMessage != "" {
		err := fmt.Errorf("technitium API error: %s", s.ErrorMessage)
		return err
	}

	return nil
}
