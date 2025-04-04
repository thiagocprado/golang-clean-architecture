package user

import "github.com/thiagocprado/golang-api-structure/pkg/errs"

type UseCase struct {
	repository IRepository
}

func NewUseCase(repository IRepository) IUseCase {
	return &UseCase{repository}
}

func (t *UseCase) Create(payload UserPayload) *errs.Err {
	// TODO - lógica para buscar usuário por email e verificar se já existe

	err := t.repository.Create(payload)
	if err != nil {
		return errs.InternalServerError("Erro ao cadastrar usuário!", err)
	}

	return nil
}

func (t *UseCase) FindAll(filters FindAllFilters) ([]UserModel, int, *errs.Err) {
	clients, count, err := t.repository.FindAll(filters)
	if err != nil {
		return nil, count, errs.InternalServerError("Erro ao buscar usuários!", err)
	}

	return clients, count, nil
}
