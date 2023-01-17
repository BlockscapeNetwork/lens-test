package cosmos

import (
	"fmt"
	"io"
	"net/http"
)

type ChainRegistry struct {
	RegistryResponse
}

func GetChainRegistry(chainName string) (*ChainRegistry, error) {

	url := "https://proxy.atomscan.com/directory"

	fullUrl := fmt.Sprintf("%v/%v/chain.json", url, chainName)

	resp, err := http.Get(fullUrl)
	if err != nil {
		fmt.Printf("An error occurred sending the request %v", err)
		return &ChainRegistry{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("An error occurred reading the body %v", err)
		return &ChainRegistry{}, err
	}

	parsedBody, err := UnmarshalRegistryResponse(body)
	if err != nil {
		fmt.Printf("An error occurred parsing the chain data %v", err)
	}

	chainRegistry := ChainRegistry{RegistryResponse: parsedBody}

	return &chainRegistry, nil
}
