package transaction

import (
	"context"
	"testing"

	"github.com/AllanCordeiro/fc-ms-wallet/internal/entity"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/event"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/usecase/mocks"
	"github.com/AllanCordeiro/fc-ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCaseExecute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "j@client.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000.0)

	client2, _ := entity.NewClient("Jane Doe", "jane@client.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000.0)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100.0,
	}

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalance := event.NewBalanceUpdated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransaction, eventBalance)
	output, err := uc.Execute(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
