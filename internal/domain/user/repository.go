package user

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &Repository{db}
}

func (t *Repository) Create(payload UserPayload) error {
	return nil
}

func (t *Repository) FindAll(budgetFilters FindAllFilters) ([]UserModel, int, error) {
	var users []UserModel
	var count int

	return users, count, nil
}
