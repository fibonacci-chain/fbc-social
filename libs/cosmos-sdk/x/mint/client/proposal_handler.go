package client

import (
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/mint/client/cli"
	"github.com/fibonacci-chain/fbc-social/libs/cosmos-sdk/x/mint/client/rest"
	govcli "github.com/fibonacci-chain/fbc-social/x/gov/client"
)

var (
	ManageTreasuresProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdManageTreasuresProposal,
		rest.ManageTreasuresProposalRESTHandler,
	)
	ModifyNextBlockUpdateProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdModifyNextBlockUpdateProposal,
		rest.ModifyNextBlockUpdateProposalRESTHandler,
	)
)
