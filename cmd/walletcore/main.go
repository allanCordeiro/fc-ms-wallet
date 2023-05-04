package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/AllanCordeiro/fc-ms-wallet/internal/database"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/event"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/event/handler"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/usecase/account"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/usecase/client"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/usecase/transaction"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/web"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/web/webserver"
	"github.com/AllanCordeiro/fc-ms-wallet/pkg/events"
	"github.com/AllanCordeiro/fc-ms-wallet/pkg/kafka"
	"github.com/AllanCordeiro/fc-ms-wallet/pkg/uow"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://sql/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	m.Up()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))
	eventDispatcher.Register("AccountCreated", handler.NewCreateAccountKafkaHandler(kafkaProducer))

	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()
	accountCreatedEvent := event.NewAccountCreated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)
	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := account.NewCreateAccountUseCase(accountDb, clientDb, eventDispatcher, accountCreatedEvent)
	createTransactionUseCase := transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer(":8080")
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)
	log.Println("server is running")
	webserver.Start()
}
