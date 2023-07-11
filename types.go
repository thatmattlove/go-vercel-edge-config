package edgeconfig

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

const VERCEL_API_URL string = "https://api.vercel.com"
const VERCEL_EDGE_CONFIG_URL string = "https://edge-config.vercel.com"

type ClientOptions struct {
	// Vercel Team ID, if applicable.
	TeamID string
	// Edge-Config Access Token.
	EdgeConfigToken string
	// Edge-Config ID. This can also be passed as an argument to client methods if the ID changes frequently.
	EdgeConfigID string
}

type VercelEdgeConfigClient struct {
	edgeConfigURL *url.URL
	TeamID        string
	client        *resty.Client
	API           *VercelAPI
	EdgeConfigID  string
}

type VercelAPI struct {
	TeamID           string
	client           *resty.Client
	hasAuthenticated bool
}

type EdgeConfig struct {
	Slug        string    `json:"slug"`
	ItemCount   int       `json:"itemCount"`
	CreatedAt   Timestamp `json:"createdAt"`
	UpdatedAt   Timestamp `json:"updatedAt"`
	ID          string    `json:"id"`
	Digest      string    `json:"digest"`
	SizeInBytes int64     `json:"sizeInBytes"`
	OwnerID     string    `json:"ownerId"`
}

type VercelAPIErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type VercelAPIError struct {
	Error VercelAPIErrorDetail `json:"error"`
}

type Timestamp struct {
	time.Time
}

type EdgeConfigItems map[string]string

func (p *Timestamp) UnmarshalJSON(bytes []byte) error {
	var raw int64
	err := json.Unmarshal(bytes, &raw)

	if err != nil {
		fmt.Printf("error decoding timestamp: %s\n", err)
		return err
	}

	p.Time = time.Unix(raw, 0)
	return nil
}
