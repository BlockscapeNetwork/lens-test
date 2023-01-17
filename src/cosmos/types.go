package cosmos

import (
	lens "github.com/strangelove-ventures/lens/client"
	"go.uber.org/zap"
)

// ModuleConfig is the governance module's configuration.
type ModuleConfig struct {
	LogLevel string `toml:"log_level" json:"log_level"`
	VoterKey string `toml:"voter_key" json:"voter_key"`
}

// Module holds all data for the governance logic.
type Module struct {
	Logger *zap.Logger

	ChainConfig   *lens.ChainClientConfig
	ChainClient   *lens.ChainClient
	ChainRegistry *ChainRegistry

	SrcWalletAddress string
}
