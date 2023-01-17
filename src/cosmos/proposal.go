package cosmos

import (
	"context"
	"fmt"
	"log"
	"math"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	// "github.com/cosmos/cosmos-sdk/proto/cosmos/gov/v1beta1/gov.proto"

	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/metadata"
)

// type TextProposal struct {
// 	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
// 	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
// 	// contains filtered or unexported fields
// }

func (m *Module) GetProposals() error {
	m.Logger.Info("Getting proposals...")
	// m.Logger().Info("Updating pending proposals...")
	voterAddr, err := m.ChainClient.GetKeyAddress()
	if err != nil {
		return fmt.Errorf("couldn't get account address: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Proposals in the voting period which the voter has already voted on.
	// See https://buf.build/cosmos/cosmos-sdk/docs/main:cosmos.gov.v1beta1#cosmos.gov.v1beta1.Query.Proposals

	done := m.ChainClient.SetSDKContext()
	// voter, err := m.ChainClient.EncodeBech32AccAddr(voterAddr) // Calls SDK config
	if err != nil {
		m.Logger.Sugar().Errorf("Can't stringify address %v", voterAddr)
		done()
		return err
	}

	resp, _, err := m.ChainClient.RunGRPCQuery(
		ctx,
		"/cosmos.gov.v1beta1.Query/Proposals",
		&govtypes.QueryProposalsRequest{
			ProposalStatus: govtypes.StatusVotingPeriod,
			Voter:          "",
			Pagination: &query.PageRequest{
				Limit: math.MaxUint64 - 1,
			},
		},
		metadata.MD{},
	)
	done()

	var proposals govtypes.QueryProposalsResponse

	if err != nil {
		return fmt.Errorf("failed to query proposals: %v",
			err.Error())
	}
	if err := func() error {
		done := m.ChainClient.SetSDKContext()
		defer done()
		if err := proposals.Unmarshal(resp.Value); err != nil { // Calls SDK config
			return err
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("failed to unmarshal response of proposals: %v", err)
	}

	if len(proposals.Proposals) == 0 {
		return fmt.Errorf("no proposals")
	}

	p := proposals.Proposals[0]

	/*
		PROPOSAL UNMARSHALING
	*/

	// resources:
	// https://medium.com/@denniswon/cosmos-sdk-encoding-dd7a75b80b85
	// https://docs.cosmos.network/main/core/encoding#interface-encoding-and-usage-of-any

	// https://github.com/cosmos/cosmos-sdk/blob/v0.45.9/x/gov/spec/02_state.md#proposals
	// https://github.com/cosmos/cosmos-sdk/blob/v0.45.9/proto/cosmos/gov/v1beta1/gov.proto#L68
	// https://github.com/cosmos/cosmos-sdk/blob/main/codec/types/any.go#L108
	// https://github.com/cosmos/cosmos-sdk/blob/main/codec/proto_codec.go#L259
	// https://github.com/cosmos/cosmos-sdk/blob/main/codec/proto_codec.go#L291
	// https://github.com/cosmos/cosmos-sdk/blob/v0.45.9/x/gov/types/proposal.go#L20
	// https://github.com/cosmos/cosmos-sdk/blob/v0.45.9/x/gov/types/query.pb.go#L80

	// https://pkg.go.dev/github.com/cosmos/cosmos-sdk@v0.46.4/codec/types#AnyUnpacker

	fmt.Printf("Proposal ID: %v\n", p.ProposalId)
	fmt.Printf("Status %v\n", p.Status.String())
	fmt.Printf("TallyResult Abstain %v\n", p.FinalTallyResult.Abstain.String())
	fmt.Printf("TotalDeposit %v\n", p.TotalDeposit.String())

	fmt.Printf("Content %T\n", p.Content)

	// all doesn't work, how do I get the content of p.Content? 🙄

	fmt.Printf("getcachedvalue %T\n", p.Content.GetCachedValue())
	fmt.Printf("getcachedvalue %v\n", p.Content.GetCachedValue().(govtypes.Content).GetDescription())

	// fmt.Printf("amino unpackany %v\n", codec.Amino.UnpackAny(p.Content, govtypes.Content))
	// fmt.Printf("amino unpackany %v\n", codec.Amino.UnpackAny(p.Content, govtypes.Content{}))

	// fmt.Printf("unmashal %v\n", p.Content.Unmarshal(p.Content.Value))
	// fmt.Printf("anypb unmarshalto %v\n", anypb.UnmarshalTo(p.Content, p.Content.ProtoMessage()))
	// fmt.Printf("anypb unmashalto %v\n", anypb.UnmarshalTo(p.Content, p.Content.Value))

	/*
		UNMARSHALING END
	*/

	return nil
}

func VoteOnProposal(m *Module) {
	const memo = ""
	resp, err := m.ChainClient.SendMsg(context.Background(), &govtypes.MsgVote{
		ProposalId: 395,
		Voter:      "osmo1cqmk75gg0l8pcevyj0hsn5qpgyl0zpj07rwkv3",
		Option:     govtypes.OptionYes,
	}, memo)
	if err != nil {
		if resp != nil {
			log.Fatalf("failed to send coins: code(%d) msg(%s)", resp.Code, resp.Logs)
		}
		log.Fatalf("Failed to send coins.Err: %v", err)
	}
	fmt.Println(m.ChainClient.PrintTxResponse(resp))
}
