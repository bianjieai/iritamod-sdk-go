package integration

import (
	"fmt"

	"github.com/bianjieai/iritamod-sdk-go/wasm"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func (s IntegrationTestSuite) TestWasm() {
	baseTx := types.BaseTx{
		From:     s.Account().Name,
		Gas:      4000000,
		Fee:      types.NewDecCoins(types.NewInt64DecCoin("upoint", 24)),
		Memo:     "test",
		Mode:     types.Commit,
		Password: s.Account().Password,
	}

	request := wasm.StoreRequest{
		WASMFile: "./scripts/election.wasm",
	}

	codeID, err := s.Wasm.Store(request, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), codeID)

	args := wasm.NewArgs().
		Put("start", 1).
		Put("end", 1000).
		Put("candidates", []string{"iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa", "iaa1zk2tse0pkk87p2v8tcsfs0ytfw3t88kejecye5"})

	initReq := wasm.InstantiateRequest{
		CodeID:  codeID,
		Label:   "test wasm",
		InitMsg: args,
	}

	contractAddress, err := s.Wasm.Instantiate(initReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), contractAddress)

	info, err := s.Wasm.QueryContractInfo(contractAddress)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), info)
	require.Equal(s.T(), fmt.Sprintf("%d", info.CodeID), codeID)
	require.Equal(s.T(), info.Label, initReq.Label)

	execAbi := wasm.NewContractABI().
		WithMethod("vote").
		WithArgs("candidate", "iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa")
	_, err = s.Wasm.Execute(contractAddress, execAbi, nil, baseTx)
	require.NoError(s.T(), err)

	queryAbi := wasm.NewContractABI().
		WithMethod("get_vote_info")
	bz, err := s.Wasm.QueryContract(contractAddress, queryAbi)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), bz)
}
