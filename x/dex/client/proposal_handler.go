package client

import (
	"github.com/fibonacci-chain/fbc-social/x/dex/client/cli"
	"github.com/fibonacci-chain/fbc-social/x/dex/client/rest"
	govclient "github.com/fibonacci-chain/fbc-social/x/gov/client"
)

// param change proposal handler
var (
	// DelistProposalHandler alias gov NewProposalHandler
	DelistProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitDelistProposal, rest.DelistProposalRESTHandler)
)
