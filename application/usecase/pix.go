package usecase

import (
	"github.com/felipematthew/codepix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key, kind, accountId string) (*model.PixKey, error) {
	// Verificar se a conta existe
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(account, kind, key)
	// Verificar se a chave está validada
	if err != nil {
		return nil, err
	}

	// Register key
	p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, err
	}

	// Retorna a chave de erro sem erro e limpa
	return pixKey, nil
}

// Vai buscar a chave e vai retornar a chave e o tipo 
func (p *PixUseCase) FindKey(key, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}