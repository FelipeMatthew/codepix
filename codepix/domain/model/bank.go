package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// Criou o base para sempre ter o id, create e updated automaticamente
type Bank struct {
	Base `valid:"required"`
	Code string `json: "code" valid:"notnull"`
	Name string `json: "name" valid:"notnull"`
	// Relations
	Accounts []*Account `valid:"-"`
}

// Method
// Vai validar o new bank, se não for valido retorna error
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

// Constructor
// * - static method
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil
}
