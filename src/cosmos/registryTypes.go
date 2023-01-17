package cosmos

import "encoding/json"

func UnmarshalRegistryResponse(data []byte) (RegistryResponse, error) {
	var r RegistryResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RegistryResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RegistryResponse struct {
	Schema       string     `json:"$schema"`
	ChainName    string     `json:"chain_name"`
	Status       string     `json:"status"`
	NetworkType  string     `json:"network_type"`
	Website      string     `json:"website"`
	PrettyName   string     `json:"pretty_name"`
	ChainID      string     `json:"chain_id"`
	Bech32Prefix string     `json:"bech32_prefix"`
	DaemonName   string     `json:"daemon_name"`
	NodeHome     string     `json:"node_home"`
	KeyAlgos     []string   `json:"key_algos"`
	Slip44       int64      `json:"slip44"`
	Fees         Fees       `json:"fees"`
	Staking      Staking    `json:"staking"`
	Codebase     Codebase   `json:"codebase"`
	Peers        Peers      `json:"peers"`
	Apis         Apis       `json:"apis"`
	Explorers    []Explorer `json:"explorers"`
}

type Apis struct {
	RPC  []Url `json:"rpc"`
	REST []Url `json:"rest"`
	Grpc []Url `json:"grpc"`
}

type Url struct {
	Address  string `json:"address"`
	Provider string `json:"provider"`
}

type Codebase struct {
	GitRepo            string   `json:"git_repo"`
	RecommendedVersion string   `json:"recommended_version"`
	CompatibleVersions []string `json:"compatible_versions"`
	CosmosSDKVersion   string   `json:"cosmos_sdk_version"`
	TendermintVersion  string   `json:"tendermint_version"`
	CosmwasmVersion    string   `json:"cosmwasm_version"`
	CosmwasmEnabled    bool     `json:"cosmwasm_enabled"`
	Genesis            Genesis  `json:"genesis"`
}

type Genesis struct {
	GenesisURL string `json:"genesis_url"`
}

type Explorer struct {
	Kind        string  `json:"kind"`
	URL         string  `json:"url"`
	TxPage      string  `json:"tx_page"`
	AccountPage *string `json:"account_page,omitempty"`
}

type Fees struct {
	FeeTokens []FeeToken `json:"fee_tokens"`
}

type FeeToken struct {
	Denom            string  `json:"denom"`
	FixedMinGasPrice float64 `json:"fixed_min_gas_price"`
	LowGasPrice      float64 `json:"low_gas_price"`
	AverageGasPrice  float64 `json:"average_gas_price"`
	HighGasPrice     float64 `json:"high_gas_price"`
}

type Peers struct {
	Seeds           []Seed           `json:"seeds"`
	PersistentPeers []PersistentPeer `json:"persistent_peers"`
}

type PersistentPeer struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

type Seed struct {
	ID       string  `json:"id"`
	Address  string  `json:"address"`
	Provider *string `json:"provider,omitempty"`
}

type Staking struct {
	StakingTokens []StakingToken `json:"staking_tokens"`
}

type StakingToken struct {
	Denom string `json:"denom"`
}
