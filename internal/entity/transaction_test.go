package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	expectedClient1Amount := 900.0
	expectedClient2Amount := 1100.0
	client1, _ := NewClient("Client 1", "client1@email.com")
	client2, _ := NewClient("Client 2", "client2@email.com")
	account1 := NewAccount(client1)
	account2 := NewAccount(client2)
	account1.Credit(1000.0)
	account2.Credit(1000.0)

	transaction, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, expectedClient2Amount, account2.Balance)
	assert.Equal(t, expectedClient1Amount, account1.Balance)
}

func TestCreateTransactionWithInsuficientFunds(t *testing.T) {
	client1, _ := NewClient("Client 1", "client1@email.com")
	client2, _ := NewClient("Client 2", "client2@email.com")
	account1 := NewAccount(client1)
	account2 := NewAccount(client2)
	account1.Credit(1000.0)
	account2.Credit(1000.0)

	transaction, err := NewTransaction(account1, account2, 2000.0)
	assert.NotNil(t, err)
	assert.Error(t, err, "insuficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)

}
