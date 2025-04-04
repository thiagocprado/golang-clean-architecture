package parsers

import (
	"encoding/json"
	"net/http"

	"github.com/thiagocprado/golang-api-structure/pkg/errs"
)

func Response[T any](resp *http.Response) (T, *errs.Err) {
	defer resp.Body.Close()

	var result T
	err := json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, errs.InternalServerError("Erro ao decodificar resposta!", err)
	}

	return result, nil
}
