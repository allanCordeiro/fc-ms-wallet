package main

import (
	"database/sql"
	"fmt"

	"github.com/AllanCordeiro/fc-ms-wallet/internal/database"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/event"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/usecase/account"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/usecase/client"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/usecase/transaction"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/web"
	"github.com/AllanCordeiro/fc-ms-wallet/internal/web/webserver"
	"github.com/AllanCordeiro/fc-ms-wallet/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	transactionCreatedEvent := event.NewTransactionCreated()
	eventDispatcher := events.NewEventDispatcher()
	//eventDispatcher.Register("TransactionCreated", handler)
	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)
	webserver.Start()
}
