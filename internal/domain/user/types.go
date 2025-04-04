package user

import (
	"fmt"
	"log/slog"
	"net/url"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/thiagocprado/golang-api-structure/pkg/errs"
)

type IUseCase interface {
	Create(payload UserPayload) *errs.Err
	FindAll(filters FindAllFilters) ([]UserModel, int, *errs.Err)
}

type IRepository interface {
	Create(payload UserPayload) error
	FindAll(filters FindAllFilters) ([]UserModel, int, error)
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

type UserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type UserModel struct {
	ID       int
	Email    string
	Name     string
	Password string
}

type FindAllFilters struct {
	Page     int    `json:"page" validate:"omitempty,numeric"`
	PageSize int    `json:"pageSize" validate:"omitempty,numeric"`
	OrderBy  string `json:"orderBy" validate:"omitempty"`
	Sorting  string `json:"sorting" validate:"omitempty"`
}

func GetFindAllFilters(urlValues url.Values) (FindAllFilters, *errs.Err) {
	page, err := strconv.Atoi(urlValues.Get("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(urlValues.Get("page_size"))
	if err != nil {
		pageSize = 10
	}

	orderBy := urlValues.Get("order_by")
	if orderBy == "" {
		orderBy = "ID"
	}

	sorting := urlValues.Get("sorting")
	if sorting == "" {
		sorting = "ASC"
	}

	clientsFilter := FindAllFilters{
		Page:     page,
		PageSize: pageSize,
		OrderBy:  orderBy,
		Sorting:  sorting,
	}

	err = validator.New().Struct(clientsFilter)
	if err != nil {
		slog.Error("Parâmetro inválido", slog.String("err", err.Error()))
		return FindAllFilters{}, errs.BadRequest(fmt.Sprintf("Falha ao validar filtros: %v", err.Error()), err)
	}

	return clientsFilter, nil
}
