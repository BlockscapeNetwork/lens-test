package cosmos

import (
	"fmt"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func sendCoins(walletAddress, destinationWallet, amountToSend string) *banktypes.MsgSend {
	// Sanitize coin amount and make it readable by SDK
	coin, err := sdk.ParseCoinNormalized(amountToSend)
	if err != nil {
		log.Fatalf("Error parsing coin string. Error: %s", err)
	}

	//	Build transaction message
	req := &banktypes.MsgSend{
		FromAddress: walletAddress,
		ToAddress:   destinationWallet,
		Amount:      sdk.Coins{coin},
	}

	fmt.Println(req)

	// Send message and get response
	// res, err := chainClient.SendMsg(context.Background(), req)
	// if err != nil {
	// 	if res != nil {
	// 		log.Fatalf("failed to send coins: code(%d) msg(%s)", res.Code, res.Logs)
	// 	}
	// 	log.Fatalf("Failed to send coins.Err: %v", err)
	// }
	// fmt.Println(chainClient.PrintTxResponse(res))

	// return res, req, err
	return req
}
