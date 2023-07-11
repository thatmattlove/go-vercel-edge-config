package edgeconfig_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	edgeconfig "github.com/thatmattlove/go-vercel-edge-config"
)

func TestVercelEdgeConfig_AllItems(t *testing.T) {
	t.Run("list all items with arg", func(t *testing.T) {
		env, err := LoadEnv()
		assert.NoError(t, err)
		opts := &edgeconfig.ClientOptions{
			EdgeConfigToken: env.EdgeConfigToken,
		}
		client, err := edgeconfig.New(opts)
		assert.NoError(t, err)
		items, err := client.AllItems(env.EdgeConfigID)
		assert.NoError(t, err)
		assert.Len(t, items, 1)
	})
	t.Run("list all items with opt", func(t *testing.T) {
		env, err := LoadEnv()
		assert.NoError(t, err)
		opts := &edgeconfig.ClientOptions{
			EdgeConfigToken: env.EdgeConfigToken,
			EdgeConfigID:    env.EdgeConfigID,
		}
		client, err := edgeconfig.New(opts)
		assert.NoError(t, err)
		items, err := client.AllItems()
		assert.NoError(t, err)
		assert.Len(t, items, 1)
	})
}

func TestVercelEdgeConfig_Item(t *testing.T) {
	t.Run("get single items with arg", func(t *testing.T) {
		env, err := LoadEnv()
		assert.NoError(t, err)
		opts := &edgeconfig.ClientOptions{
			EdgeConfigToken: env.EdgeConfigToken,
		}
		client, err := edgeconfig.New(opts)
		assert.NoError(t, err)
		value, err := client.Item(env.EdgeConfigID, env.TestKey)
		assert.NoError(t, err)
		assert.Equal(t, env.TestValue, value)
	})
	t.Run("get single items with opt", func(t *testing.T) {
		env, err := LoadEnv()
		assert.NoError(t, err)
		opts := &edgeconfig.ClientOptions{
			EdgeConfigToken: env.EdgeConfigToken,
			EdgeConfigID:    env.EdgeConfigID,
		}
		client, err := edgeconfig.New(opts)
		assert.NoError(t, err)
		value, err := client.Item(env.TestKey)
		assert.NoError(t, err)
		assert.Equal(t, env.TestValue, value)
	})
}
func TestVercelEdgeConfig_Digest(t *testing.T) {
	t.Run("get digest with arg", func(t *testing.T) {
		env, err := LoadEnv()
		assert.NoError(t, err)
		opts := &edgeconfig.ClientOptions{
			EdgeConfigToken: env.EdgeConfigToken,
		}
		client, err := edgeconfig.New(opts)
		assert.NoError(t, err)
		digest, err := client.Digest(env.EdgeConfigID)
		assert.NoError(t, err)
		assert.Equal(t, env.Digest, digest)
	})
	t.Run("get digest with opt", func(t *testing.T) {
		env, err := LoadEnv()
		assert.NoError(t, err)
		opts := &edgeconfig.ClientOptions{
			EdgeConfigToken: env.EdgeConfigToken,
			EdgeConfigID:    env.EdgeConfigID,
		}
		client, err := edgeconfig.New(opts)
		assert.NoError(t, err)
		digest, err := client.Digest()
		assert.NoError(t, err)
		assert.Equal(t, env.Digest, digest)
	})
}
