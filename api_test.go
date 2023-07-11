package edgeconfig_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	edgeconfig "github.com/thatmattlove/go-vercel-edge-config"
)

func TestVercelEdgeConfigClient_ListAll(t *testing.T) {
	t.Run("list all configs", func(t *testing.T) {
		env, err := LoadEnv()
		assert.NoError(t, err)
		opts := &edgeconfig.ClientOptions{
			EdgeConfigToken: env.EdgeConfigToken,
		}
		client, err := edgeconfig.New(opts)
		assert.NoError(t, err)
		err = client.API.Authenticate(env.APIToken)
		assert.NoError(t, err)
		configs, err := client.API.ListAllEdgeConfigs()
		assert.NoError(t, err)
		assert.Len(t, configs, 1)
	})
}
