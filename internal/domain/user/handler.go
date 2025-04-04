package user

import (
	"net/http"

	"github.com/thiagocprado/golang-api-structure/pkg/handles"
	"github.com/thiagocprado/golang-api-structure/pkg/presenters"
)

type Handler struct {
	usecase IUseCase
}

func NewHandler(usecase IUseCase) *Handler {
	return &Handler{usecase}
}

func (t *Handler) Create(w http.ResponseWriter, r *http.Request) {
	payload, err := handles.Payload[UserPayload](r.Body)
	if err != nil {
		handles.Error(w, err)
		return
	}

	err = t.usecase.Create(payload)
	if err != nil {
		handles.Error(w, err)
		return
	}

	handles.Success(w, http.StatusNoContent, nil)
}

func (t *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	usersFilter, err := GetFindAllFilters(urlValues)
	if err != nil {
		handles.Error(w, err)
		return
	}

	users, count, err := t.usecase.FindAll(usersFilter)
	if err != nil {
		handles.Error(w, err)
		return
	}

	dto := PresenterUsers(users)
	resp := presenters.BuildResponsePagination(usersFilter.Page, usersFilter.PageSize, count, dto)

	handles.Success(w, http.StatusOK, resp)
}
