package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommitConfig(t *testing.T) {
	exampleConfig := CommitPluginConfig{
		SourceStartBlock:       222,
		DestStartBlock:         333,
		OffRamp:                "0x123",
		TokenPricesUSDPipeline: `merge [type=merge left="{}" right="{\"0xC79b96044906550A5652BCf20a6EA02f139B9Ae5\":\"1000000000000000000\"}"];`,
	}

	bts, err := json.Marshal(exampleConfig)
	require.NoError(t, err)

	parsedConfig := CommitPluginConfig{}
	require.NoError(t, json.Unmarshal(bts, &parsedConfig))

	require.Equal(t, exampleConfig, parsedConfig)
}

func TestExecutionConfig(t *testing.T) {
	exampleConfig := ExecutionPluginConfig{
		SourceStartBlock: 222,
		DestStartBlock:   333,
	}

	bts, err := json.Marshal(exampleConfig)
	require.NoError(t, err)

	parsedConfig := ExecutionPluginConfig{}
	require.NoError(t, json.Unmarshal(bts, &parsedConfig))

	require.Equal(t, exampleConfig, parsedConfig)
}
