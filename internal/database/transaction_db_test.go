package database

import (
	"database/sql"
	"testing"

	"github.com/AllanCordeiro/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDb *TransactionDB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients(id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts(id varchar(255), client_id varchar(255), balance float, created_at date)")
	db.Exec("CREATE TABLE transactions(id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date)")
	s.transactionDb = NewTransactionDB(db)
	s.client, _ = entity.NewClient("John Doe", "john.doe@email.com")
	s.client2, _ = entity.NewClient("Jane Doe", "jane.doe@email.com")
	//creating accounts
	s.accountFrom = entity.NewAccount(s.client)
	s.accountFrom.Credit(1000.0)
	s.accountTo = entity.NewAccount(s.client2)
	s.accountTo.Credit(1000.0)
}

func (s *TransactionDBTestSuite) TearDownSuit() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")

}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDb.Create(transaction)
	s.Nil(err)
}
