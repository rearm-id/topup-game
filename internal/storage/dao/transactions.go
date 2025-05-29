package dao

import (
	"github.com/pocketbase/pocketbase/core"
)

// ensures that the Transaction struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*Transaction)(nil)

func NewTransaction(record *core.Record) *Transaction {
	return &Transaction{
		Record: record,
	}
}

// FindTransactionById finds a transaction by its ID
func FindTransactionById(app core.App, id string) (*Transaction, error) {
	record, err := app.FindRecordById("transactions", id)
	if err != nil {
		return nil, err
	}

	return NewTransaction(record), nil
}

// NewTransactionForGame creates a new Transaction for a given Game
func NewTransactionForGame(app core.App, game *Game, userID string, amount float64) (*Transaction, error) {
	collection, err := app.FindCollectionByNameOrId("transactions")
	if err != nil {
		return nil, err
	}

	record := core.NewRecord(collection)
	transaction := &Transaction{
		Record: record,
	}

	transaction.SetGame(game.Id)
	transaction.SetUser(userID)
	transaction.SetAmount(amount)
	transaction.SetStatus("pending")

	return transaction, nil
}

// FindTransactionsByUser finds all transactions for a user
func FindTransactionsByUser(app core.App, userID string) ([]*Transaction, error) {
	records, err := app.FindRecordsByFilter(
		"transactions",
		"user = {:user}",
		"-created",
		100,
		0,
		map[string]interface{}{"user": userID},
	)
	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, len(records))
	for i, record := range records {
		transactions[i] = NewTransaction(record)
	}

	return transactions, nil
}

// FindPendingTransactionsByUser finds all pending transactions for a user
func FindPendingTransactionsByUser(app core.App, userID string) ([]*Transaction, error) {
	records, err := app.FindRecordsByFilter(
		"transactions",
		"user = {:user} && status = {:status}",
		"-created",
		100,
		0,
		map[string]interface{}{
			"user":   userID,
			"status": "pending",
		},
	)
	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, len(records))
	for i, record := range records {
		transactions[i] = NewTransaction(record)
	}

	return transactions, nil
}

// CompleteTransaction marks a transaction as completed with payment metadata
func CompleteTransaction(app core.App, transaction *Transaction, paymentMetadata map[string]interface{}) error {
	transaction.SetStatus("completed")
	transaction.SetPaymentMetadata(paymentMetadata)

	return app.Save(transaction)
}

// FailTransaction marks a transaction as failed with optional payment metadata
func FailTransaction(app core.App, transaction *Transaction, paymentMetadata map[string]interface{}) error {
	transaction.SetStatus("failed")
	if paymentMetadata != nil {
		transaction.SetPaymentMetadata(paymentMetadata)
	}

	return app.Save(transaction)
}

type Transaction struct {
	*core.Record
}

// ProxyRecord implements the core.RecordProxy interface
func (t *Transaction) ProxyRecord() *core.Record {
	return t.Record
}

// SetProxyRecord implements the core.RecordProxy interface
func (t *Transaction) SetProxyRecord(r *core.Record) {
	t.Record = r
}

// GetGame returns the related game ID
func (t *Transaction) GetGame() string {
	return t.GetString("game")
}

// SetGame sets the related game ID
func (t *Transaction) SetGame(gameID string) {
	t.Set("game", gameID)
}

// GetUser returns the user associated with the transaction
func (t *Transaction) GetUser() string {
	return t.GetString("user")
}

// SetUser sets the user for the transaction
func (t *Transaction) SetUser(user string) {
	t.Set("user", user)
}

// GetAmount returns the transaction amount
func (t *Transaction) GetAmount() float64 {
	return t.GetFloat("amount")
}

// SetAmount sets the transaction amount
func (t *Transaction) SetAmount(amount float64) {
	t.Set("amount", amount)
}

// GetGameMetadata returns the game metadata
func (t *Transaction) GetGameMetadata() map[string]interface{} {
	var result map[string]interface{}
	t.UnmarshalJSONField("game_metadata", &result)
	return result
}

// SetGameMetadata sets the game metadata
func (t *Transaction) SetGameMetadata(metadata map[string]interface{}) {
	t.Set("game_metadata", metadata)
}

// GetPaymentMetadata returns the payment metadata
func (t *Transaction) GetPaymentMetadata() map[string]interface{} {
	var result map[string]interface{}
	t.UnmarshalJSONField("payment_metadata", &result)
	return result
}

// SetPaymentMetadata sets the payment metadata
func (t *Transaction) SetPaymentMetadata(metadata map[string]interface{}) {
	t.Set("payment_metadata", metadata)
}

// GetStatus returns the transaction status
func (t *Transaction) GetStatus() string {
	return t.GetString("status")
}

// SetStatus sets the transaction status
func (t *Transaction) SetStatus(status string) {
	t.Set("status", status)
}

// IsCompleted returns true if the transaction status is "completed"
func (t *Transaction) IsCompleted() bool {
	return t.GetStatus() == "completed"
}

// IsPending returns true if the transaction status is "pending"
func (t *Transaction) IsPending() bool {
	return t.GetStatus() == "pending"
}

// IsFailed returns true if the transaction status is "failed"
func (t *Transaction) IsFailed() bool {
	return t.GetStatus() == "failed"
}
