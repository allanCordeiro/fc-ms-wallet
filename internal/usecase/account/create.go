package account

import (
	"github.com/AllanCordeiro/fc-ms-wallet/internal/entity"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/gateway"
	"github.com/AllanCordeiro/fc-ms-wallet/pkg/events"
)

type CreateAccountInputDTO struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway  gateway.AccountGateway
	ClientGateway   gateway.ClientGateway
	EventDispatcher events.EventDispatcherInterface
	AccountCreated  events.EventInterface
}

func NewCreateAccountUseCase(a gateway.AccountGateway,
	c gateway.ClientGateway,
	e events.EventDispatcherInterface,
	ac events.EventInterface) *CreateAccountUseCase {

	return &CreateAccountUseCase{
		AccountGateway:  a,
		ClientGateway:   c,
		EventDispatcher: e,
		AccountCreated:  ac,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	//each account will receive 100.00 moneys to make easy to rollup the chapter challenge
	account.Balance = 100.00
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	var output CreateAccountOutputDTO
	output.ID = account.ID

	uc.AccountCreated.SetPayload(account.ID)
	uc.EventDispatcher.Dispatch(uc.AccountCreated)

	return &output, nil
}
