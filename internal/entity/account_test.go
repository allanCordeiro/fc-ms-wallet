package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John", "john@email.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithEmptyClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John", "john@email.com")
	account := NewAccount(client)
	assert.NotNil(t, account)

	account.Credit(100.0)
	assert.Equal(t, 100.0, account.Balance)

}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John", "john@email.com")
	account := NewAccount(client)
	assert.NotNil(t, account)

	account.Credit(100.0)
	account.Debit(50.0)
	assert.Equal(t, 50.0, account.Balance)

}
