package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

// Caso queira criar uma lista de transição
type Transactions struct {
	Transaction []*Transaction
}

// Criar transação de transferencias entre as contas
type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"notnull"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	// Regra de negocio
	// se o valor for zerado
	if t.Amount <= 0 {
		return errors.New("the amount mus be grater than 0")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("invalid transaction value")
	}

	// nao pode mandar dinheiro para a mesma conta
	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("you can not send money to the same account")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {

	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPending,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// Operações que a entidade realiza, get e set não é tão recomendável
// Criar metodos para utilizar variaveis dentro das entidades
// Completar transação / Cancelar transação / Confirmar Transição
func (t *Transaction) Completed() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Confirmed() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.Description = description
	err := t.isValid()
	return err
}
