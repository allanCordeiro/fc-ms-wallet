package account

import (
	"testing"

	"github.com/AllanCordeiro/fc-ms-wallet/internal/entity"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/event"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/usecase/mocks"
	"github.com/AllanCordeiro/fc-ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCaseExecute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "john@client.com")
	clientMock := &mocks.ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)
	dispatcher := events.NewEventDispatcher()
	eventAccount := event.NewAccountCreated()

	uc := NewCreateAccountUseCase(accountMock, clientMock, dispatcher, eventAccount)
	input := CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)

}
