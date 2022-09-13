package keeper_test

import (
	"github.com/elesto-dao/elesto/v2/x/credential"
	"github.com/elesto-dao/elesto/v2/x/credential/keeper"
)

func (s *KeeperTestSuite) TestGenesis() {
	//default genesis on export
	genState := keeper.ExportGenesis(s.ctx, s.keeper)
	s.Require().Equal(genState, &credential.GenesisState{})
}
