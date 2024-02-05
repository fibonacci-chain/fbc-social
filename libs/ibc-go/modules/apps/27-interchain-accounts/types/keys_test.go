package types_test

import "github.com/fibonacci-chain/fbc-social/libs/ibc-go/modules/apps/27-interchain-accounts/types"

func (suite *TypesTestSuite) TestKeyActiveChannel() {
	key := types.KeyActiveChannel("port-id", "connection-id")
	suite.Require().Equal("activeChannel/port-id/connection-id", string(key))
}

func (suite *TypesTestSuite) TestKeyOwnerAccount() {
	key := types.KeyOwnerAccount("port-id", "connection-id")
	suite.Require().Equal("owner/port-id/connection-id", string(key))
}
