package handler

import (
	"fmt"
	"sync"

	"github.com/AllanCordeiro/fc-ms-wallet/pkg/events"
	"github.com/AllanCordeiro/fc-ms-wallet/pkg/kafka"
)

type CreateAccountKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewCreateAccountKafkaHandler(kafka *kafka.Producer) *CreateAccountKafkaHandler {
	return &CreateAccountKafkaHandler{
		Kafka: kafka,
	}
}

func (h *CreateAccountKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	h.Kafka.Publish(message, nil, "accounts")
	fmt.Println("Create accounts kafka handler called")
}
