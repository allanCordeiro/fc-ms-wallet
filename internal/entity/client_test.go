package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	expectedName := "John Doe"
	expectedEmail := "john_doe@client.com"
	client, err := NewClient(expectedName, expectedEmail)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, expectedName, client.Name)
	assert.Equal(t, expectedEmail, client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	expectedName := "Jane Doe"
	expectedEmail := "jane@client.com"
	client, _ := NewClient("John Doe", "j.doe@client.com")
	err := client.Update(expectedName, expectedEmail)

	assert.Nil(t, err)
	assert.Equal(t, expectedName, client.Name)
	assert.Equal(t, expectedEmail, client.Email)
	assert.Greater(t, client.UpdatedAt, client.CreatedAt)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	expectedName := ""
	expectedEmail := ""
	client, _ := NewClient("John Doe", "j.doe@client.com")
	err := client.Update(expectedName, expectedEmail)

	assert.NotNil(t, err)
	assert.Equal(t, expectedName, client.Name)
	assert.Equal(t, expectedEmail, client.Email)
}

func TestAddAcountToClient(t *testing.T) {
	expectedName := "John Doe"
	expectedEmail := "john_doe@client.com"
	client, _ := NewClient(expectedName, expectedEmail)
	account := NewAccount(client)
	err := client.AddAcount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))

}
