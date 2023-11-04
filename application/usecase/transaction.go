package usecase

import (
	"errors"
	"log"

	"github.com/felipematthew/codepix/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository         model.PixKeyRepositoryInterface
}

// Registrar uma nova transação
func (t *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string, id string) (*model.Transaction, error) {
	// Verificar se conta existe
	account, err := t.PixRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	// Verificar se existe chave
	pixKey, err := t.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	// verifica possibilidade de transação
	transaction, err := model.NewTransaction(account, amount, pixKey, description, id)
	if err != nil {
		return nil, err
	}

	t.TransactionRepository.Save(transaction)
	// verificar se está retornando algo
	if transaction.ID == "" {
		return transaction, nil
	}

	// Caso não encontre nenhum registro
	return nil, errors.New("Unable to process this transction")
}

// Confirmar Transação 
func (t *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error)  {
	// Vai procurar o id da transição
	transaction, err := t.TransactionRepository.Find(transactionId) 
	if err != nil {
		log.Printf("Transaction not found", transactionId) 
		return nil, err
	}

	// Se existir muda status e coloca como confirmada
	transaction.Status = model.TransactionConfirmed
	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}
	
	return transaction, nil
}

// Completar Tranasação
func (t *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	// Vai procurar o id da transição
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Printf("Transaction not found", transactionId)
		return nil, err
	}
	// Se existir muda status e coloca como completa
	transaction.Status = model.TransactionCompleted
	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Erro na transação
func (t *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	// Vai procurar o id da transição
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	// Altera o status
	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}
	
	return transaction, nil
}		