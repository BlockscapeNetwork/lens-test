package main

import (
	"os"

	"lens-test/src"
)

// Needed vars for this example:
var (
	// We will be fetching chain info from chain registry.
	// This string must match the relevant directory name in the chain registry here:
	// https://github.com/cosmos/chain-registry
	chainRegistryName = "cosmoshub"

	walletMnemonic     = os.Getenv("testKeyMn")
	destination_wallet = "osmo18pjx07s8aq42qtxeskw2d7h8edfank7r6clwrj"
	amountToSend       = "1uosmo"
)

func main() {
	logger := src.GetLogger()

	m := src.GetModule(logger, chainRegistryName, walletMnemonic)
	// fmt.Printf("SrcWalletAddress %v", m.SrcWalletAddress)

	m.GetProposals()
}
