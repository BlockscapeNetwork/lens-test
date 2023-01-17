package cosmos

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// ProposalExcerpt is an excerpt of a proposal.
type ProposalExcerpt struct {
	ProposalId uint64                  `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Content    *codectypes.Any         `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Status     govtypes.ProposalStatus `protobuf:"varint,3,opt,name=status,proto3,enum=cosmos.gov.v1beta1.ProposalStatus" json:"status,omitempty"`
	// final_tally_result is the final tally result of the proposal. When
	// querying a proposal via gRPC, this field is not populated until the
	// proposal's voting period has ended.
	FinalTallyResult govtypes.TallyResult `protobuf:"bytes,4,opt,name=final_tally_result,json=finalTallyResult,proto3" json:"final_tally_result"`
	SubmitTime       time.Time            `protobuf:"bytes,5,opt,name=submit_time,json=submitTime,proto3,stdtime" json:"submit_time"`
	DepositEndTime   time.Time            `protobuf:"bytes,6,opt,name=deposit_end_time,json=depositEndTime,proto3,stdtime" json:"deposit_end_time"`
	TotalDeposit     cosmostypes.Coins    `json:"-"` /* 141-byte string literal not displayed */
	VotingStartTime  time.Time            `protobuf:"bytes,8,opt,name=voting_start_time,json=votingStartTime,proto3,stdtime" json:"voting_start_time"`
	VotingEndTime    time.Time            `protobuf:"bytes,9,opt,name=voting_end_time,json=votingEndTime,proto3,stdtime" json:"voting_end_time"`
}
