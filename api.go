package edgeconfig

import (
	"encoding/json"
	"fmt"
)

func (v *VercelAPI) checkAuth() (err error) {
	if !v.hasAuthenticated {
		err = fmt.Errorf("call Authenticate() before any other method")
	}
	return
}

func (v *VercelAPI) Authenticate(token string) (err error) {
	v.client.SetAuthScheme("Bearer")
	v.client.SetAuthToken(token)
	v.client.SetBaseURL(VERCEL_API_URL)
	if v.TeamID != "" {
		v.client.SetQueryParam("teamId", v.TeamID)
	}
	req := v.client.R()
	res, err := req.Get("/v6/deployments")
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		err = fmt.Errorf("authentication failed")
		return
	}
	v.hasAuthenticated = true
	return
}

func (v *VercelAPI) ListAllEdgeConfigs() (edgeConfigs []*EdgeConfig, err error) {
	err = v.checkAuth()
	if err != nil {
		return
	}
	var vercelError *VercelAPIError
	req := v.client.R()
	req.SetError(vercelError)
	res, err := req.Get("/v1/edge-config")
	if err != nil {
		return
	}
	if res.IsError() {
		vercelError = res.Error().(*VercelAPIError)
		err = fmt.Errorf("%d %s Error: %s", res.StatusCode(), res.Status(), vercelError.Error.Message)
		return
	}
	err = json.Unmarshal(res.Body(), &edgeConfigs)
	return
}
