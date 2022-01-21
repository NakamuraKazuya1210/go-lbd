package lbd

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveUserInformation(t *testing.T) {
	ret, err := l.RetrieveUserInformation(userId)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveUserWalletTransactionHistory(t *testing.T) {
	ret, err := l.RetrieveUserWalletTransactionHistory(userId)
	assert.Nil(t, err)
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestIssueSessionTokenForBaseCoinTransfer(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueSessionTokenForBaseCoinTransfer(userId, owner.Address, big.NewInt(1), RequestTypeAOA)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestFuncTransferNonFungibleUserWallet(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.TransferNonFungibleUserWallet(itemTokenContractId, userId, "tlink10ps670a0x6ma5knthjjswgw89d44vmz6xm3umr", tokenType, "00000009")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestIssueSessionTokenForProxySetting(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueSessionTokenForProxySetting(userId, itemTokenContractId, RequestTypeAOA)
	sessionToken = ret.RequestSessionToken
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveSessionTokenStatus(t *testing.T) {
	if sessionToken == "" {
		t.Skip("Skip because no session token")
	}
	ret, err := l.RetrieveSessionTokenStatus(sessionToken)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestCommitTransaction(t *testing.T) {
	if sessionToken == "" {
		t.Skip("Skip because no session token")
	}
	ret, err := l.CommitTransaction(sessionToken)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveBaseCoinBalance(t *testing.T) {
	ret, err := l.RetrieveBaseCoinBalance(userId)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfAllServiceTokens(t *testing.T) {
	ret, err := l.RetrieveBalanceOfAllServiceTokens(userId)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfSpecificServiceToken(t *testing.T) {
	ret, err := l.RetrieveBalanceOfSpecificServiceToken(userId, serviceTokenContractId)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfAllFungibles(t *testing.T) {
	ret, err := l.RetrieveBalanceOfAllFungibles(userId, itemTokenContractId)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfSpecificFungible(t *testing.T) {
	ret, err := l.RetrieveBalanceOfSpecificFungible(userId, itemTokenContractId, "00000001")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfNonFungiblesWithTokenType(t *testing.T) {
	ret, err := l.RetrieveBalanceOfNonFungiblesWithTokenType(userId, itemTokenContractId, "asc", "", 10)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfSpecificNonFungible(t *testing.T) {
	ret, err := l.RetrieveBalanceOfSpecificNonFungible(userId, itemTokenContractId, tokenType, "00000001")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestRetrieveWhetherTheServiceTokenProxySet(t *testing.T) {
	ret, err := l.RetrieveWhetherTheServiceTokenProxySet(userId, serviceTokenContractId)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestRetrieveWhetherTheItemTokenProxySet(t *testing.T) {
	ret, err := l.RetrieveWhetherTheItemTokenProxySet(userId, itemTokenContractId)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestIssueSessionTokenForServiceTokenTransfer(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueSessionTokenForServiceTokenTransfer(userId, serviceTokenContractId, owner.Address, big.NewInt(1), RequestTypeAOA)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestIssueSessionTokenForServiceTokenProxySetting(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueSessionTokenForServiceTokenProxySetting(userId, serviceTokenContractId, RequestTypeAOA)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestTransferDelegatedServiceToken(t *testing.T) {
	ret, err := l.TransferDelegatedServiceToken(userId, serviceTokenContractId, owner.Address, big.NewInt(1000))
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestTransferDelegatedFungible(t *testing.T) {
	ret, err := l.TransferDelegatedFungible(userId, itemTokenContractId, tokenType, owner.Address, big.NewInt(1000))
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestTransferDelegatedNonFungible(t *testing.T) {
	ret, err := l.TransferDelegatedNonFungible(userId, itemTokenContractId, tokenType, "00000001", owner.Address)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestBatchTransferDelegatedNonFungibles(t *testing.T) {
	ret, err := l.BatchTransferDelegatedNonFungibles(userId, itemTokenContractId, owner.Address, []TransferList{{TokenId: tokenType + "00000001"}, {TokenId: tokenType + "00000002"}})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
