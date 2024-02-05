package client

import (
	"github.com/fibonacci-chain/fbc-social/x/farm/client/cli"
	"github.com/fibonacci-chain/fbc-social/x/farm/client/rest"
	govcli "github.com/fibonacci-chain/fbc-social/x/gov/client"
)

var (
	// ManageWhiteListProposalHandler alias gov NewProposalHandler
	ManageWhiteListProposalHandler = govcli.NewProposalHandler(cli.GetCmdManageWhiteListProposal, rest.ManageWhiteListProposalRESTHandler)
)
