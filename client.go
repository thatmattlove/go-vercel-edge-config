package edgeconfig

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
)

// AllItems retrieves all items from an edge config store and returns a map.
func (ec *VercelEdgeConfigClient) AllItems(args ...string) (items EdgeConfigItems, err error) {
	configID, err := ec.getConfigID(args)
	if err != nil {
		return
	}
	var vercelError *VercelAPIError
	req := ec.client.R()
	req.SetError(vercelError)
	res, err := req.Get(fmt.Sprintf("%s/items", configID))
	if err != nil {
		return
	}
	if res.IsError() {
		vercelError = res.Error().(*VercelAPIError)
		err = fmt.Errorf("%d %s Error: %s", res.StatusCode(), res.Status(), vercelError.Error.Message)
		return
	}
	err = json.Unmarshal(res.Body(), &items)
	return
}

// Item retrieves a single item from an edge config store by key.
func (ec *VercelEdgeConfigClient) Item(args ...string) (value string, err error) {
	var configID string
	var key string
	argsLen := len(args)
	if argsLen == 1 && ec.EdgeConfigID == "" {
		err = fmt.Errorf("edge config ID must either be specified when initializing the client, or on each request")
		return
	} else if argsLen == 1 && ec.EdgeConfigID != "" {
		configID = ec.EdgeConfigID
		key = args[0]
	} else if argsLen == 2 {
		configID, err = ec.getConfigID(args)
		if err != nil {
			return
		}
		key = args[1]
	} else {
		err = fmt.Errorf("method Item accepts 1-2 arguments. If the edge config ID is defined when initializing the client, only the key is required. Otherwise, both the edge config ID and the key are required")
		return
	}
	if key == "" {
		err = fmt.Errorf("a key must be provided")
		return
	}
	var vercelError *VercelAPIError
	req := ec.client.R()
	req.SetError(vercelError)
	res, err := req.Get(fmt.Sprintf("%s/item/%s", configID, key))
	if err != nil {
		return
	}
	if res.IsError() {
		vercelError = res.Error().(*VercelAPIError)
		err = fmt.Errorf("%d %s Error: %s", res.StatusCode(), res.Status(), vercelError.Error.Message)
		return
	}
	err = json.Unmarshal(res.Body(), &value)
	return
}

// Digest retrieves the digest for an edge config store.
func (ec *VercelEdgeConfigClient) Digest(args ...string) (digest string, err error) {
	configID, err := ec.getConfigID(args)
	if err != nil {
		return
	}
	var vercelError *VercelAPIError
	req := ec.client.R()
	req.SetError(vercelError)
	res, err := req.Get(fmt.Sprintf("%s/digest", configID))
	if err != nil {
		return
	}
	if res.IsError() {
		vercelError = res.Error().(*VercelAPIError)
		err = fmt.Errorf("%d %s Error: %s", res.StatusCode(), res.Status(), vercelError.Error.Message)
		return
	}
	err = json.Unmarshal(res.Body(), &digest)
	return
}

func (ec *VercelEdgeConfigClient) getConfigID(args []string) (configID string, err error) {
	if len(args) == 0 {
		if ec.EdgeConfigID == "" {
			err = fmt.Errorf("edge config ID must either be specified when initializing the client, or on each request")
			return
		} else {
			configID = ec.EdgeConfigID
			return
		}
	}
	edgeConfigIDArg := args[0]
	if ec.EdgeConfigID != "" {
		configID = ec.EdgeConfigID
	} else if edgeConfigIDArg == "" {
		err = fmt.Errorf("edge config ID must either be specified when initializing the client, or on each request")
		return
	} else {
		configID = edgeConfigIDArg
	}
	return
}

// New creates a new Vercel Edge Config client using options.
func New(options *ClientOptions) (client *VercelEdgeConfigClient, err error) {
	if options.EdgeConfigToken == "" {
		err = fmt.Errorf("missing required option: 'EdgeConfigToken'")
		return
	}

	ecURL, err := url.Parse(VERCEL_EDGE_CONFIG_URL)
	if err != nil {
		return
	}

	ec := resty.New()
	ec.SetBaseURL(ecURL.String())
	ec.SetAuthScheme("Bearer")
	ec.SetAuthToken(options.EdgeConfigToken)

	if options.TeamID != "" {
		ec.SetQueryParam("teamId", options.TeamID)
	}

	api := &VercelAPI{
		client:           resty.New(),
		TeamID:           options.TeamID,
		hasAuthenticated: false,
	}

	client = &VercelEdgeConfigClient{
		edgeConfigURL: ecURL,
		EdgeConfigID:  options.EdgeConfigID,
		TeamID:        options.TeamID,
		client:        ec,
		API:           api,
	}
	return
}

// NewFromConnectionString creates a new Vercel Edge Config client from a connection string.
func NewFromConnectionString(connectionString string) (client *VercelEdgeConfigClient, err error) {
	u, err := url.Parse(connectionString)
	if err != nil {
		return
	}
	token := u.Query().Get("token")
	if token == "" {
		err = fmt.Errorf("token must be specified as a query parameter in the connection string")
		return
	}
	if len(u.Path) == 0 {
		err = fmt.Errorf("edge config ID must be specified as the path")
		return
	}
	paths := strings.Split(u.Path, "/")
	if len(u.Path) != 1 {
		err = fmt.Errorf("invalid path. Edge config ID must be specified as the path")
		return
	}
	configID := paths[0]
	opts := &ClientOptions{
		TeamID:          "",
		EdgeConfigToken: token,
		EdgeConfigID:    configID,
	}
	return New(opts)
}
