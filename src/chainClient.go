package src

import (
	"context"
	"lens-test/src/cosmos"
	"log"
	"os"

	lens "github.com/strangelove-ventures/lens/client"
	registry "github.com/strangelove-ventures/lens/client/chain_registry"
	"go.uber.org/zap"
)

func GetModule(logger *zap.Logger, chainRegistryName, walletMnemonic string) *cosmos.Module {

	chainClient, chainConfig, chainRegistry, srcWalletAddress := getChainClient(logger, chainRegistryName, walletMnemonic)

	module := cosmos.Module{
		Logger:        logger,
		ChainConfig:   chainConfig,
		ChainClient:   chainClient,
		ChainRegistry: chainRegistry,

		SrcWalletAddress: srcWalletAddress,
	}

	return &module
}

func getChainClient(logger *zap.Logger, chainRegistryName, walletMnemonic string) (*lens.ChainClient, *lens.ChainClientConfig, *cosmos.ChainRegistry, string) {
	//	Fetches chain info from chain registry
	chainInfo, err := registry.DefaultChainRegistry(logger).GetChain(context.Background(), chainRegistryName)
	if err != nil {
		log.Fatalf("Failed to get chain info. Err: %v \n", err)
	}

	//	Use Chain info to select random endpoint
	rpc, err := chainInfo.GetRandomRPCEndpoint(context.Background())
	if err != nil {
		log.Fatalf("Failed to get random RPC endpoint on chain %s. Err: %v \n", chainInfo.ChainID, err)
	}

	// For this example, lets place the key directory in your PWD.
	pwd, _ := os.Getwd()
	key_dir := pwd + "/keys"

	// Build chain config
	chainConfig := lens.ChainClientConfig{
		Key:     "default",
		ChainID: chainInfo.ChainID,
		RPCAddr: rpc,
		// GRPCAddr       string,
		AccountPrefix:  chainInfo.Bech32Prefix,
		KeyringBackend: "test",
		GasAdjustment:  1.3,
		// GasPrices:      "0.01uosmo",
		KeyDirectory: key_dir,
		Debug:        true,
		Timeout:      "20s",
		OutputFormat: "json",
		SignModeStr:  "direct",
		Modules:      lens.ModuleBasics,
	}

	// Creates client object to pull chain info
	chainClient, err := lens.NewChainClient(&zap.Logger{}, &chainConfig, key_dir, os.Stdin, os.Stdout)
	if err != nil {
		log.Fatalf("Failed to build new chain client for %s. Err: %v \n", chainInfo.ChainID, err)
	}

	keyName := "source_key"

	var srcWalletAddress string
	if chainClient.KeyExists(keyName) {
		srcWalletAddress, err = chainClient.ShowAddress(keyName)
		if err != nil {
			log.Fatalf("Failed to show address. Err: %v \n", err)
		}
	} else {
		// Lets restore a key with funds and name it keyName, this is the wallet we'll use to send tx.
		srcWalletAddress, err = chainClient.RestoreKey(keyName, walletMnemonic, 0)
		if err != nil {
			log.Fatalf("Failed to restore key. Err: %v \n", err)
		}
	}

	// Now that we know our key name, we can set it in our chain config
	chainConfig.Key = "source_key"

	chainRegistry, err := cosmos.GetChainRegistry(chainRegistryName)
	if err != nil {
		logger.Sugar().Error("Failed to fetch chain registry for %v", chainRegistryName)
	}

	return chainClient, &chainConfig, chainRegistry, srcWalletAddress
}
