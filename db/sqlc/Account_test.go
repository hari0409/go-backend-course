package db

import (
	"context"
	"testing"

	"github.com/hari0409/backend-go/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, args.Owner, account.Owner)
	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc := CreateRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.NoError(t, err)
	require.Equal(t, acc.ID, account.ID)
	require.NoError(t, err)
	require.Equal(t, acc.Balance, account.Balance)
}

func TestDeleteAccount(t *testing.T) {
	acc := CreateRandomAccount(t)
	testQueries.DeleteAccount(context.Background(), acc.ID)
	accDel, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.Error(t, err)
	require.Empty(t, accDel)
}

func TestUpdateAccount(t *testing.T) {
	acc := CreateRandomAccount(t)
	arg := UpdateAccountParams{Balance: util.RandomMoney(), ID: acc.ID}
	accUpdated, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accUpdated)
	require.Equal(t, arg.Balance, accUpdated.Balance)
}

func TestListAccounts(t *testing.T) {
	accounts := []Account{}
	for i := 0; i < 5; i++ {
		acc := CreateRandomAccount(t)
		accounts = append(accounts, acc)
	}
	listAccs, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{Limit: 5, Offset: 0})
	require.NoError(t, err)
	require.Len(t, listAccs, len(accounts))
}
